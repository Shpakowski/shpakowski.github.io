package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mcp-chain/logging"
	"mcp-chain/types"
	"os"
)

// State — структура для хранения текущего состояния сети MCP Chain
// (см. описание в state_layout.md)
type State struct {
	ChainMeta        types.ChainMeta                // Метаданные цепочки
	Supply           types.Supply                   // Денежная масса MCP
	Accounts         map[string]types.Account       // ключ — string (hex address)
	Validators       map[string]types.ValidatorInfo // ключ — string (hex address)
	CoopsRegistry    map[string]types.CoopData      // Кооперативы
	Governance       map[string]types.Proposal      // Голосования сети
	SoulBound        map[string]types.STBSet        // ключ — string (hex address)
	ContractsStorage types.ContractsStorage         // Хранилище контрактов
	FeeTreasury      types.FeeTreasury              // Казна комиссий
	Blocks           []types.Block                  // Цепочка блоков

	// Не сериализуется, не влияет на StateRoot.
	Cache types.RuntimeCache // Временный кэш
}

// GlobalState — глобальный RAM-state для доступа из любого core-модуля
var GlobalState *State

// NewState создаёт новое пустое состояние
func NewState() *State {
	return &State{
		ChainMeta:        types.ChainMeta{},
		Supply:           types.Supply{},
		Accounts:         make(map[string]types.Account),
		Validators:       make(map[string]types.ValidatorInfo),
		CoopsRegistry:    make(map[string]types.CoopData),
		Governance:       make(map[string]types.Proposal),
		SoulBound:        make(map[string]types.STBSet),
		ContractsStorage: make(types.ContractsStorage),
		FeeTreasury:      types.FeeTreasury{},
		Blocks:           []types.Block{},
		Cache:            types.RuntimeCache{},
	}
}

// TODO: Реализовать методы для работы с состоянием:
// - ApplyTx(tx *types.Transaction) error
// - ApplyBlock(block *types.Block) error
// - DumpJSON(path string) error
// - LoadJSON(path string) error
// - RevertTx(tx *types.Transaction) error
// - PruneOldState() error
// - GetAccount(addr types.Address) *types.Account
// - GetValidator(addr types.Address) *types.ValidatorInfo
// - и др.

// DumpStateToFile сохраняет всё состояние в файл (например, data/state.json)
func DumpStateToFile(st *State, path string) error {
	logging.Logger.Info("DumpStateToFile", "path", path)
	os.MkdirAll("data", 0700)
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0600)
}

// LoadStateFromFile загружает всё состояние из файла
func LoadStateFromFile(path string) (*State, error) {
	logging.Logger.Info("LoadStateFromFile", "path", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	st := NewState()
	err = json.Unmarshal(data, st)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// GetAccount возвращает аккаунт по адресу (или nil, если нет)
func (s *State) GetAccount(addr types.Address) *types.Account {
	acc, ok := s.Accounts[string(addr)]
	if !ok {
		return nil
	}
	logging.Logger.Debug("GetAccount", "addr", addr, "account", acc)
	return &acc
}

// GetBalance возвращает баланс по адресу (0, если нет аккаунта)
func (s *State) GetBalance(addr types.Address) types.Amount {
	acc, ok := s.Accounts[string(addr)]
	if !ok {
		return 0
	}
	logging.Logger.Debug("GetBalance", "addr", addr, "balance", acc.Balance)
	return acc.Balance
}

// ApplyTx применяет транзакцию к state: обновляет балансы, nonce, учитывает комиссию
func (s *State) ApplyTx(tx *types.Transaction) error {
	logging.Logger.Info("ApplyTx: применение транзакции", "from", tx.From, "to", tx.To, "amount", tx.Amount, "fee", tx.Fee, "nonce", tx.Nonce)
	// Получаем аккаунты отправителя и получателя
	sender, ok := s.Accounts[string(tx.From)]
	if !ok {
		return fmt.Errorf("sender %s not found", tx.From)
	}
	recipient, ok := s.Accounts[string(tx.To)]
	if !ok {
		// Если получатель новый — создаём аккаунт
		logging.Logger.Info("ApplyTx: создаём новый аккаунт для получателя", "to", tx.To)
		recipient = types.Account{Balance: 0, Nonce: 0}
	}
	amount := types.Amount(tx.Amount)
	fee := types.Amount(tx.Fee)
	if sender.Balance < amount+fee {
		return fmt.Errorf("insufficient balance: have %d, need %d", sender.Balance, amount+fee)
	}
	oldSender := sender
	oldRecipient := recipient
	// Списываем сумму и комиссию с отправителя
	sender.Balance -= amount + fee
	sender.Nonce++
	// Зачисляем сумму получателю
	recipient.Balance += amount
	// Обновляем аккаунты в state
	s.Accounts[string(tx.From)] = sender
	s.Accounts[string(tx.To)] = recipient
	logging.Logger.Info("ApplyTx: после применения", "from", tx.From, "old_sender", oldSender, "new_sender", sender, "to", tx.To, "old_recipient", oldRecipient, "new_recipient", recipient)
	// Комиссия уходит в FeeTreasury
	oldTreasury := s.FeeTreasury.Balance
	s.FeeTreasury.Balance += fee
	logging.Logger.Info("ApplyTx: FeeTreasury", "old", oldTreasury, "new", s.FeeTreasury.Balance)
	return nil
}

func (s *State) SetAccount(addr string, acc types.Account) {
	old, ok := s.Accounts[addr]
	if ok {
		logging.Logger.Info("SetAccount: обновление аккаунта", "addr", addr, "old", old, "new", acc)
	} else {
		logging.Logger.Info("SetAccount: новый аккаунт", "addr", addr, "new", acc)
	}
	s.Accounts[addr] = acc
}

func (s *State) SetBalance(addr string, balance types.Amount) {
	acc, ok := s.Accounts[addr]
	if !ok {
		logging.Logger.Warn("SetBalance: аккаунт не найден", "addr", addr)
		return
	}
	old := acc.Balance
	acc.Balance = balance
	logging.Logger.Info("SetBalance", "addr", addr, "old_balance", old, "new_balance", balance)
	s.Accounts[addr] = acc
}

// ApplyStateChange — универсальный способ изменить state с автосохранением
func ApplyStateChange(change func(s *State)) error {
	logging.Logger.Info("ApplyStateChange: ДО", "accounts_count", len(GlobalState.Accounts), "accounts", keys(GlobalState.Accounts))
	change(GlobalState)
	logging.Logger.Info("ApplyStateChange: ПОСЛЕ", "accounts_count", len(GlobalState.Accounts), "accounts", keys(GlobalState.Accounts))
	return DumpStateToFile(GlobalState, "data/state.json")
}

// keys возвращает список ключей map[string]V
func keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

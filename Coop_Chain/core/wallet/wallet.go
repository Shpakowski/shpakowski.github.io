package wallet

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mcp-chain/core/state"
	"mcp-chain/crypto"
	"os"
	"path/filepath"

	"mcp-chain/types"

	"mcp-chain/internal/config"

	"mcp-chain/logging"

	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ed25519"
)

const WalletsFile = "data/wallets.json"

type WalletEntry struct {
	Seed   string `json:"seed"`
	PubKey string `json:"pub_key"`
}

type Wallets struct {
	List []WalletEntry `json:"wallets"`
}

type Manager struct {
	file string
}

// WalletBalance — структура для баланса кошелька
type WalletBalance struct {
	Address string
	Balance float64
}

func NewManager(file string) *Manager {
	return &Manager{file: file}
}

func (m *Manager) Load() Wallets {
	var ws Wallets
	data, err := ioutil.ReadFile(m.file)
	if err == nil {
		_ = json.Unmarshal(data, &ws)
	}
	return ws
}

func (m *Manager) Save(ws Wallets) {
	os.MkdirAll(filepath.Dir(m.file), 0700)
	data, _ := json.MarshalIndent(ws, "", "  ")
	_ = ioutil.WriteFile(m.file, data, 0600)
}

func (m *Manager) AddWallet(seed string) (WalletEntry, error) {
	logging.Logger.Info("AddWallet", "seed", seed)
	seedBytes := bip39.NewSeed(seed, "")
	pub, _ := pubKeyFromSeed(seedBytes)
	entry := WalletEntry{Seed: seed, PubKey: crypto.PublicKeyToHex(pub)}
	ws := m.Load()
	for _, w := range ws.List {
		if w.Seed == seed {
			return w, nil // already exists
		}
	}
	ws.List = append(ws.List, entry)
	m.Save(ws)
	return entry, nil
}

func (m *Manager) ListWallets() []WalletEntry {
	logging.Logger.Info("ListWallets: получение всех кошельков")
	return m.Load().List
}

func (m *Manager) ImportWallet(seed string, cfg *config.NetworkConfig) (WalletEntry, error) {
	logging.Logger.Info("ImportWallet", "pubkey", seed)
	if !bip39.IsMnemonicValid(seed) {
		return WalletEntry{}, ErrInvalidMnemonic
	}
	entry, err := m.AddWallet(seed)
	if err != nil {
		return entry, err
	}
	addr := types.Address(entry.PubKey)
	err = state.ApplyStateChange(func(st *state.State) {
		if acc, ok := st.Accounts[string(addr)]; ok {
			acc.Balance += types.Amount(cfg.WalletCreationReward)
			st.Accounts[string(addr)] = acc
		} else {
			st.Accounts[string(addr)] = types.Account{Balance: types.Amount(cfg.WalletCreationReward), Nonce: 0}
		}
		logging.Logger.Info("ImportWallet: Accounts после добавления", "accounts", keys(st.Accounts))
	})
	if err != nil {
		logging.Logger.Error("Ошибка сохранения state после импорта кошелька", "err", err)
		return entry, err
	}
	logging.Logger.Info("ImportWallet", "pubkey", entry.PubKey, "reward", cfg.WalletCreationReward)
	return entry, nil
}

func pubKeyFromSeed(seedBytes []byte) (ed25519.PublicKey, error) {
	if len(seedBytes) < 32 {
		return nil, ErrSeedTooShort
	}
	pub, _, _ := ed25519.GenerateKey(bytesReader(seedBytes))
	return pub, nil
}

func bytesReader(b []byte) *os.File {
	tmp := "data/.tmp_seed"
	ioutil.WriteFile(tmp, b, 0600)
	f, _ := os.Open(tmp)
	return f
}

var (
	ErrInvalidMnemonic = errors.New("invalid mnemonic (seed phrase)")
	ErrSeedTooShort    = errors.New("seed too short")
)

// IsMnemonic возвращает true, если строка — валидная seed-фраза (mnemonic)
func IsMnemonic(s string) bool {
	return bip39.IsMnemonicValid(s)
}

// GetWalletBalances возвращает список балансов для всех кошельков
func GetWalletBalances(mgr *Manager) []WalletBalance {
	logging.Logger.Info("GetWalletBalances: получение балансов всех кошельков")
	var res []WalletBalance
	for addr, acc := range state.GlobalState.Accounts {
		res = append(res, WalletBalance{Address: addr, Balance: float64(acc.Balance)})
	}
	return res
}

// GetWalletBalance возвращает баланс для одного адреса
func GetWalletBalance(addr string) WalletBalance {
	logging.Logger.Info("GetWalletBalance: получение баланса кошелька", "addr", addr)
	balance := state.GlobalState.GetBalance(types.Address(addr))
	return WalletBalance{Address: addr, Balance: float64(balance)}
}

// ImportWalletAndSave импортирует кошелёк, начисляет награду, сохраняет state, логирует
func ImportWalletAndSave(cfg *config.NetworkConfig, seed string) (WalletEntry, error) {
	logging.Logger.Info("ImportWalletAndSave: импорт и сохранение кошелька", "seed", seed)
	manager := NewManager("") // не используем файлы для manager
	entry, err := manager.ImportWallet(seed, cfg)
	if err != nil {
		logging.Logger.Error("Ошибка импорта кошелька", "err", err)
		return entry, err
	}
	logging.Logger.Info("ImportWalletAndSave: кошелек импортирован", "pubkey", entry.PubKey)
	return entry, nil
}

// ListWalletsFromState возвращает список всех кошельков (адресов) из state и их балансы
func ListWalletsFromState() []WalletBalance {
	logging.Logger.Info("ListWalletsFromState: получение всех кошельков из state")
	var res []WalletBalance
	for addr, acc := range state.GlobalState.Accounts {
		res = append(res, WalletBalance{Address: addr, Balance: float64(acc.Balance)})
	}
	return res
}

// keys возвращает список ключей map[string]T
func keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

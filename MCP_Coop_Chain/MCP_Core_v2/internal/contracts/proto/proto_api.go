package proto

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"mcp-coop-chain/internal/types"

	"github.com/btcsuite/btcutil/base58"
)

// ProtoAPI содержит встроенные функции Transfer, Balance и др.
// TODO: реализовать Transfer(from, to, amount), CheckBalance(wallet)

// ProtoContract определяет интерфейс встроенного контракта
// Call(storage types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error)
type ProtoContract interface {
	Call(wallets types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error)
}

// protoRegistry содержит карту всех встроенных контрактов
var protoRegistry = map[string]ProtoContract{}

// RegisterProtoContract регистрирует новый proto-контракт
func RegisterProtoContract(name string, contract ProtoContract) {
	protoRegistry[name] = contract
}

// CallProtoContract вызывает встроенный контракт по имени
// Теперь требует передачи storage-абстракций
func CallProtoContract(name string, wallets types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error) {
	c, ok := protoRegistry[name]
	if !ok {
		return nil, errors.New("proto contract not found: " + name)
	}
	return c.Call(wallets, contracts, args)
}

// --- Заглушки для встроенных контрактов ---

// TransferContract реализует перевод MCP между кошельками
func (t *TransferContract) Call(wallets types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error) {
	var req types.TransferArgs
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	if req.From == "" || req.To == "" || req.Amount == 0 {
		return json.Marshal(types.TransferResult{Success: false, Error: "invalid arguments"})
	}
	// Проверка существования кошельков
	_, err := wallets.GetWallet(req.From)
	if err != nil {
		return json.Marshal(types.TransferResult{Success: false, Error: "sender not found"})
	}
	_, err = wallets.GetWallet(req.To)
	if err != nil {
		return json.Marshal(types.TransferResult{Success: false, Error: "recipient not found"})
	}
	// Проверка баланса и комиссии (минимум 1000)
	const MinFee = 1000 // 0.001 MCP
	if req.Amount < MinFee {
		return json.Marshal(types.TransferResult{Success: false, Error: "amount below minimum fee"})
	}
	// Снимаем сумму + комиссию с отправителя
	total := int64(req.Amount)
	err = wallets.UpdateBalance(req.From, -total)
	if err != nil {
		return json.Marshal(types.TransferResult{Success: false, Error: "insufficient funds"})
	}
	// Зачисляем сумму получателю (без комиссии)
	err = wallets.UpdateBalance(req.To, int64(req.Amount-MinFee))
	if err != nil {
		// Откат отправителю
		_ = wallets.UpdateBalance(req.From, total)
		return json.Marshal(types.TransferResult{Success: false, Error: "recipient update failed"})
	}
	// TODO: комиссия может идти в системный кошелёк или сгорать
	return json.Marshal(types.TransferResult{Success: true})
}

// GetBalanceContract реализует получение баланса кошелька
func (g *GetBalanceContract) Call(wallets types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error) {
	var req types.GetBalanceArgs
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	bal, err := wallets.GetBalance(req.Wallet)
	if err != nil {
		return json.Marshal(types.GetBalanceResult{Balance: 0})
	}
	return json.Marshal(types.GetBalanceResult{Balance: bal})
}

// CreateWalletContract реализует создание кошелька по публичному ключу
func (c *CreateWalletContract) Call(wallets types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error) {
	var req types.CreateWalletArgs
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	if req.PublicKey == "" {
		return json.Marshal(types.CreateWalletResult{Address: ""})
	}
	address := getWalletAddress(req.PublicKey)
	w := types.Wallet{
		PublicKey: req.PublicKey,
		Address:   address,
	}
	err := wallets.AddWallet(w)
	if err != nil {
		return json.Marshal(types.CreateWalletResult{Address: ""})
	}
	return json.Marshal(types.CreateWalletResult{Address: address})
}

// AddSmartContractContract реализует регистрацию пользовательского контракта
func (a *AddSmartContractContract) Call(wallets types.WalletStorage, contracts types.ContractStorage, args []byte) ([]byte, error) {
	var req types.AddSmartContractArgs
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	if req.Name == "" || req.Code == "" || req.Author == "" {
		return json.Marshal(types.AddSmartContractResult{Success: false, Error: "invalid arguments"})
	}
	err := contracts.AddContract(req.Name, req.Code, req.Author)
	if err != nil {
		return json.Marshal(types.AddSmartContractResult{Success: false, Error: err.Error()})
	}
	res, _ := json.Marshal(types.AddSmartContractResult{Success: true})
	return res, nil
}

// --- Определения типов контрактов ---
type TransferContract struct{}
type GetBalanceContract struct{}
type CreateWalletContract struct{}
type AddSmartContractContract struct{}

// --- Локальная функция для генерации адреса кошелька (адаптер) ---
func getWalletAddress(pubKey string) string {
	// SHA256 + base58 (дублирует internal/wallet/wallet.go)
	// Для чистоты архитектуры лучше вынести в types, но сейчас дублируем
	hash := sha256.Sum256([]byte(pubKey))
	return base58.Encode(hash[:])
}

// Register all proto contracts (инициализация)
// func init() {
// 	RegisterProtoContract("Transfer", &TransferContract{})
// 	RegisterProtoContract("GetBalance", &GetBalanceContract{})
// 	RegisterProtoContract("CreateWallet", &CreateWalletContract{})
// 	RegisterProtoContract("AddSmartContract", &AddSmartContractContract{})
// }

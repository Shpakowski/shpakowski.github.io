package contracts

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
)

// Contract — интерфейс для встроенных и пользовательских контрактов
// Позволяет расширять систему без изменения диспетчера
// Все контракты должны реализовать метод Call
// params — карта аргументов (имя -> значение)
type Contract interface {
	Call(params map[string]interface{}) (string, error)
}

// CreateCoopContract — встроенный контракт CreateCoop
// (пример реализации)
type CreateCoopContract struct{}

func (c *CreateCoopContract) Call(params map[string]interface{}) (string, error) {
	// TODO: валидация параметров, логика создания кооператива
	return "CoopCreated", nil
}

// VoteContract — встроенный контракт Vote
type VoteContract struct{}

func (c *VoteContract) Call(params map[string]interface{}) (string, error) {
	// TODO: валидация параметров, логика голосования
	return "Voted", nil
}

// IssueTokenContract — встроенный контракт IssueToken
type IssueTokenContract struct{}

func (c *IssueTokenContract) Call(params map[string]interface{}) (string, error) {
	// TODO: валидация параметров, логика выпуска токена
	return "TokenIssued", nil
}

// Контейнер для всех встроенных контрактов
var protoContracts = map[string]Contract{
	"CreateCoop": &CreateCoopContract{},
	"Vote":       &VoteContract{},
	"IssueToken": &IssueTokenContract{},
}

// CallProtoContract маршрутизирует вызов на нужный контракт по имени
func CallProtoContract(name string, params map[string]interface{}) (string, error) {
	logging.Logger.Info("CallProtoContract", "name", name, "params", params)
	contract, ok := protoContracts[name]
	if !ok {
		logging.Logger.Error("CallProtoContract: неизвестный контракт", "name", name)
		return "", fmt.Errorf("неизвестный контракт: %s", name)
	}
	return contract.Call(params)
}

// UserContract — MVP-структура для пользовательского контракта
type UserContract struct {
	Code  string
	State map[string]interface{}
}

func (uc *UserContract) Call(method string, args ...string) (string, error) {
	// MVP: просто возвращаем info о вызове
	return "UserContract called: method=" + method + ", args=" + fmt.Sprint(args), nil
}

// userContracts — пул пользовательских контрактов (address -> contract)
var userContracts = map[string]*UserContract{}

// DeployContract деплоит новый контракт, возвращает адрес и txID
func DeployContract(code string, args ...string) (string, string, error) {
	addr := randomHex(16)
	txID := randomHex(16)
	userContracts[addr] = &UserContract{Code: code, State: map[string]interface{}{}}
	logging.Logger.Info("DeployContract", "address", addr, "txID", txID)
	return addr, txID, nil
}

// CallUserContract вызывает метод пользовательского контракта
func CallUserContract(address string, method string, args ...string) (string, error) {
	c, ok := userContracts[address]
	if !ok {
		return "", fmt.Errorf("контракт не найден: %s", address)
	}
	return c.Call(method, args...)
}

// randomHex генерирует случайную hex-строку длины n
func randomHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// DeployContractAndSave деплоит контракт, сохраняет state, логирует
func DeployContractAndSave(cfg *config.NetworkConfig, code string, args ...string) (string, string, error) {
	addr, txID, err := DeployContract(code, args...)
	if err != nil {
		logging.Logger.Error("Ошибка деплоя контракта", "err", err)
		return addr, txID, err
	}
	logging.Logger.Info("DeployContractAndSave: контракт задеплоен", "address", addr, "txID", txID)
	return addr, txID, nil
}

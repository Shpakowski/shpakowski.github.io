package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mcp-coop-chain/internal/types"
	"os"
	"time"
)

// FileStorage — файловое хранилище (JSON) для блокчейна MCP Coop Chain
// Сохраняет и загружает состояние сети в файл

type FileStorage struct {
	FilePath string // Путь к файлу для хранения состояния
}

// NewFileStorage создает новый экземпляр файлового хранилища
func NewFileStorage(path string) *FileStorage {
	return &FileStorage{FilePath: path}
}

// SaveState сохраняет состояние сети в файл
func (f *FileStorage) SaveState(data []byte) error {
	return ioutil.WriteFile(f.FilePath, data, 0644)
}

// LoadState загружает состояние сети из файла
func (f *FileStorage) LoadState() ([]byte, error) {
	if _, err := os.Stat(f.FilePath); os.IsNotExist(err) {
		return nil, nil // Файл не существует — возвращаем nil
	}
	return ioutil.ReadFile(f.FilePath)
}

// SaveFullSnapshot сериализует полный снимок состояния сети в JSON-файл (перезапись).
func SaveFullSnapshot(snapshot *types.FullChainSnapshot, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(snapshot)
}

// LoadFullSnapshot загружает полный снимок состояния сети из JSON-файла.
func LoadFullSnapshot(path string) (*types.FullChainSnapshot, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var snapshot types.FullChainSnapshot
	dec := json.NewDecoder(file)
	if err := dec.Decode(&snapshot); err != nil {
		return nil, errors.New("ошибка чтения snapshot: " + err.Error())
	}
	return &snapshot, nil
}

// CreateGenesisSnapshot создаёт genesis-блок и пустые структуры, сохраняет всё в JSON.
func CreateGenesisSnapshot(path string) (*types.FullChainSnapshot, error) {
	genesisBlock := types.Block{
		Header: types.BlockHeader{
			BlockID:       "genesis",
			Height:        0,
			Timestamp:     time.Now().UTC(),
			PreviousHash:  "",
			PostStateHash: "",
			Tag:           "genesis",
			TxRoot:        "",
		},
		Proposer:      "system",
		ValidatorSigs: nil,
		Transactions:  nil,
		ContractCalls: nil,
	}
	snapshot := &types.FullChainSnapshot{
		Blocks:        []types.Block{genesisBlock},
		Mempool:       []types.Transaction{},
		Organizations: []types.Organization{},
		Wallets:       []types.Wallet{},
		State:         types.ChainState{LastHash: "", PostStateHash: ""},
		Contracts:     []types.ContractCall{},
		Timestamp:     time.Now().UTC(),
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(snapshot); err != nil {
		return nil, err
	}
	return snapshot, nil
}

// StartChainFromSnapshot определяет, что делать: загрузить или создать снимок.
func StartChainFromSnapshot(path string) (*types.FullChainSnapshot, error) {
	if _, err := os.Stat(path); err == nil {
		return LoadFullSnapshot(path)
	}
	return CreateGenesisSnapshot(path)
}

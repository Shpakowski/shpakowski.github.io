package storage

import (
	"errors"
	"mcp-coop-chain/internal/types"
	"sync"
	"time"
)

// MemoryStorage реализует Storage, хранит блоки и сериализованное состояние в памяти.
type MemoryStorage struct {
	Mu         sync.RWMutex
	Blocks     []types.Block
	BlockIndex map[string]types.Block
	State      []byte // сериализованное состояние цепочки
	Wallets    map[string]*memoryWallet
	Contracts  map[string]*memoryContract
}

// NewMemoryStorage возвращает новый экземпляр in-memory хранилища.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Mu:         sync.RWMutex{},
		Blocks:     make([]types.Block, 0),
		BlockIndex: make(map[string]types.Block),
		State:      nil,
		Wallets:    make(map[string]*memoryWallet),
		Contracts:  make(map[string]*memoryContract),
	}
}

// AddBlock добавляет новый блок, если его ещё нет.
func (ms *MemoryStorage) AddBlock(block types.Block) error {
	ms.Mu.Lock()
	defer ms.Mu.Unlock()
	blockID := block.Header.BlockID
	if _, exists := ms.BlockIndex[blockID]; exists {
		return errors.New("block already exists")
	}
	ms.Blocks = append(ms.Blocks, block)
	ms.BlockIndex[blockID] = block
	return nil
}

// GetBlockByHash возвращает блок по его хешу (BlockID).
func (ms *MemoryStorage) GetBlockByHash(hash string) (types.Block, error) {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	block, exists := ms.BlockIndex[hash]
	if !exists {
		return types.Block{}, errors.New("block not found")
	}
	return block, nil
}

// GetLastBlock возвращает последний добавленный блок.
func (ms *MemoryStorage) GetLastBlock() (types.Block, error) {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	if len(ms.Blocks) == 0 {
		return types.Block{}, errors.New("no blocks in storage")
	}
	return ms.Blocks[len(ms.Blocks)-1], nil
}

// HasBlock проверяет наличие блока по хешу.
func (ms *MemoryStorage) HasBlock(hash string) bool {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	_, exists := ms.BlockIndex[hash]
	return exists
}

// GetAllBlocks возвращает копию слайса всех блоков.
func (ms *MemoryStorage) GetAllBlocks() []types.Block {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	blocksCopy := make([]types.Block, len(ms.Blocks))
	copy(blocksCopy, ms.Blocks)
	return blocksCopy
}

// SaveState сохраняет сериализованное состояние в памяти.
func (ms *MemoryStorage) SaveState(data []byte) error {
	ms.Mu.Lock()
	defer ms.Mu.Unlock()
	ms.State = make([]byte, len(data))
	copy(ms.State, data)
	return nil
}

// LoadState возвращает сериализованное состояние из памяти.
func (ms *MemoryStorage) LoadState() ([]byte, error) {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	if ms.State == nil {
		return nil, nil
	}
	res := make([]byte, len(ms.State))
	copy(res, ms.State)
	return res, nil
}

// --- In-memory реализация WalletStorage и ContractStorage ---

// Добавляю поля для in-memory хранения кошельков, балансов и контрактов
// (добавить в MemoryStorage)
type memoryWallet struct {
	Wallet  types.Wallet
	Balance uint64 // в микро-MCP
}

type memoryContract struct {
	Name   string
	Code   string
	Author string
}

// WalletStorage impl
func (ms *MemoryStorage) GetWallet(address string) (types.Wallet, error) {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	w, ok := ms.Wallets[address]
	if !ok {
		return types.Wallet{}, errors.New("wallet not found")
	}
	return w.Wallet, nil
}

// buildSnapshot собирает актуальное состояние сети для сериализации
func (ms *MemoryStorage) buildSnapshot() *types.FullChainSnapshot {
	wallets := make([]types.Wallet, 0, len(ms.Wallets))
	for _, mw := range ms.Wallets {
		wallets = append(wallets, mw.Wallet)
	}
	contracts := make([]types.ContractCall, 0, len(ms.Contracts)) // TODO: сериализация контрактов, если нужно
	return &types.FullChainSnapshot{
		Blocks:        ms.GetAllBlocks(),
		Mempool:       []types.Transaction{}, // TODO: если есть mempool в storage
		Organizations: []types.Organization{},
		Wallets:       wallets,
		State:         types.ChainState{}, // TODO: актуальное состояние, если нужно
		Contracts:     contracts,
		Timestamp:     time.Now().UTC(),
	}
}

func (ms *MemoryStorage) AddWallet(wallet types.Wallet) error {
	ms.Mu.Lock()
	defer ms.Mu.Unlock()
	if _, exists := ms.Wallets[wallet.Address]; exists {
		return errors.New("wallet already exists")
	}
	ms.Wallets[wallet.Address] = &memoryWallet{Wallet: wallet, Balance: 0}
	return nil
}

func (ms *MemoryStorage) UpdateBalance(address string, delta int64) error {
	ms.Mu.Lock()
	defer ms.Mu.Unlock()
	w, ok := ms.Wallets[address]
	if !ok {
		return errors.New("wallet not found")
	}
	if delta < 0 && w.Balance < uint64(-delta) {
		return errors.New("insufficient funds")
	}
	w.Balance = uint64(int64(w.Balance) + delta)
	return nil
}

func (ms *MemoryStorage) GetBalance(address string) (uint64, error) {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	w, ok := ms.Wallets[address]
	if !ok {
		return 0, errors.New("wallet not found")
	}
	return w.Balance, nil
}

// ContractStorage impl
func (ms *MemoryStorage) AddContract(name, code, author string) error {
	ms.Mu.Lock()
	defer ms.Mu.Unlock()
	if _, exists := ms.Contracts[name]; exists {
		return errors.New("contract already exists")
	}
	ms.Contracts[name] = &memoryContract{Name: name, Code: code, Author: author}
	return nil
}

func (ms *MemoryStorage) GetContract(name string) (string, error) {
	ms.Mu.RLock()
	defer ms.Mu.RUnlock()
	c, ok := ms.Contracts[name]
	if !ok {
		return "", errors.New("contract not found")
	}
	return c.Code, nil
}

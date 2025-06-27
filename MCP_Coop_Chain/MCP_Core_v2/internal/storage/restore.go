package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"mcp-coop-chain/internal/types"
)

// LoadStateFromDisk загружает JSON-файл снапшота и парсит в структуру FullChainSnapshot
func LoadStateFromDisk(path string) (*types.FullChainSnapshot, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("snapshot file not found: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open snapshot file: %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read snapshot file: %w", err)
	}

	var state types.FullChainSnapshot
	if err := json.Unmarshal(bytes, &state); err != nil {
		return nil, fmt.Errorf("failed to parse snapshot JSON: %w", err)
	}

	// Проверка на дубликаты блоков, корректность данных и т.д. (см. ValidateRestoredState)
	if err := ValidateRestoredState(&state); err != nil {
		return nil, fmt.Errorf("restored state validation failed: %w", err)
	}

	return &state, nil
}

// ValidateRestoredState проверяет согласованность всех компонентов FullChainSnapshot
func ValidateRestoredState(state *types.FullChainSnapshot) error {
	// Проверка связности блоков
	for i := 1; i < len(state.Blocks); i++ {
		if state.Blocks[i].Header.PreviousHash != state.Blocks[i-1].Header.BlockID {
			return fmt.Errorf("block #%d: previous hash mismatch", i)
		}
	}
	// Проверка на дубликаты транзакций в mempool
	txIDs := make(map[string]struct{})
	for _, tx := range state.Mempool {
		if _, exists := txIDs[tx.TxID]; exists {
			return fmt.Errorf("duplicate transaction in mempool: %s", tx.TxID)
		}
		txIDs[tx.TxID] = struct{}{}
	}
	// Проверка балансов (можно расширить по бизнес-правилам)
	// TODO: Проверить соответствие балансов последнему StateRoot, если реализовано
	return nil
}

// ApplyRestoredState применяет восстановленное состояние к системам (цепочка, mempool, кошельки, контракты)
func ApplyRestoredState(state *types.FullChainSnapshot, storageIface interface{}, chainPtr *[]*types.Block, mempoolPtr *[]*types.Transaction) error {
	memStorage, ok := storageIface.(*MemoryStorage)
	if !ok {
		return fmt.Errorf("ApplyRestoredState: only MemoryStorage is supported in MVP")
	}

	// Восстанавливаем цепочку блоков
	*chainPtr = make([]*types.Block, 0, len(state.Blocks))
	for i := range state.Blocks {
		block := state.Blocks[i]
		*chainPtr = append(*chainPtr, &block)
	}

	// Восстанавливаем mempool
	*mempoolPtr = make([]*types.Transaction, 0, len(state.Mempool))
	for i := range state.Mempool {
		tx := state.Mempool[i]
		*mempoolPtr = append(*mempoolPtr, &tx)
	}

	// Восстанавливаем кошельки
	memStorage.Wallets = make(map[string]*memoryWallet)
	for _, w := range state.Wallets {
		memStorage.Wallets[w.Address] = &memoryWallet{Wallet: w, Balance: 0} // Баланс = 0, если не сериализован отдельно
	}
	// TODO: восстановить балансы, если они сериализуются отдельно

	// TODO: добавить восстановление contracts/orgs, если реализовано в MemoryStorage

	return nil
}

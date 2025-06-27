package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"mcp-coop-chain/internal"
	"mcp-coop-chain/internal/storage"
	"mcp-coop-chain/internal/types"
	"mcp-coop-chain/internal/wallet"
	"sort"
	"time"
)

// Block содержит данные блока в цепочке
// TODO: добавить поля и методы согласно архитектуре

// NewBlock создает новый блок MCP Coop Chain по заданным параметрам.
// Хеши и подписи пока временно-заглушечные.
func NewBlock(
	height uint64,
	prevHash string,
	tag string,
	proposer string,
	transactions []types.Transaction,
	contractCalls []types.ContractCall,
) types.Block {
	timestamp := time.Now().UTC()
	blockID := fmt.Sprintf("block-%d-%d", height, timestamp.UnixNano()) // временный уникальный ID
	header := types.BlockHeader{
		BlockID:       blockID,
		Height:        height,
		Timestamp:     timestamp,
		PreviousHash:  prevHash,
		PostStateHash: "", // TODO: вычислять после обработки блока
		Tag:           tag,
		TxRoot:        "", // TODO: вычислять меркл-корень транзакций
	}
	return types.Block{
		Header:        header,
		Proposer:      proposer,
		ValidatorSigs: nil, // Пока без подписей
		Transactions:  transactions,
		ContractCalls: contractCalls,
	}
}

// CreateGenesisBlock генерирует первый блок с фиксированным timestamp и predefined wallet ("founder").
func CreateGenesisBlock() (types.Block, error) {
	genesis := types.Block{
		Header: types.BlockHeader{
			BlockID:       "",
			Height:        0,
			Timestamp:     time.Now().UTC(),
			PreviousHash:  "",
			PostStateHash: "",
			Tag:           "genesis",
			TxRoot:        "",
		},
		Proposer:      "founder",
		ValidatorSigs: nil,
		Transactions:  nil,
		ContractCalls: nil,
	}
	// Баланс founder = startEmission из конфига (в микро-MCP)
	startEmission := uint64(internal.CurrentConfig.StartEmission * 1_000_000)
	genesis.Header.PostStateHash = hashState(nil, map[string]uint64{"founder": startEmission})
	genesis.Header.BlockID = hashBlockHeader(&genesis)
	return genesis, nil
}

// CreateBlock собирает транзакции из mempool, валидирует, формирует новый блок
func CreateBlock(mempool []types.Transaction, prevBlock types.Block, state map[string]uint64) (types.Block, error) {
	// Приоритизация по комиссии (по убыванию)
	sort.Slice(mempool, func(i, j int) bool {
		return mempool[i].Fee > mempool[j].Fee
	})
	// Валидация транзакций
	validTxs := make([]types.Transaction, 0, len(mempool))
	for _, tx := range mempool {
		if wallet.VerifyTransaction(tx, nil) { // TODO: передавать pubKey
			validTxs = append(validTxs, tx)
		}
	}
	// Меркл-корень
	merkle := calcMerkleRoot(validTxs)
	// Применяем транзакции к state
	newState := applyTransactions(state, validTxs)
	stateRoot := hashState(newState, nil)
	block := types.Block{
		Header: types.BlockHeader{
			BlockID:       "",
			Height:        prevBlock.Header.Height + 1,
			Timestamp:     time.Now().UTC(),
			PreviousHash:  prevBlock.Header.BlockID,
			PostStateHash: stateRoot,
			Tag:           "",
			TxRoot:        merkle,
		},
		Proposer:      "validator", // MVP
		ValidatorSigs: nil,
		Transactions:  validTxs,
		ContractCalls: nil,
	}
	block.Header.BlockID = hashBlockHeader(&block)
	return block, nil
}

// ValidateBlock проверяет подпись, хеши, MerkleRoot, StateRoot, уникальность BlockID
func ValidateBlock(block types.Block, prevBlock types.Block, prevState map[string]uint64) error {
	if block.Header.PreviousHash != prevBlock.Header.BlockID {
		return errors.New("invalid previous block hash")
	}
	if block.Header.TxRoot != calcMerkleRoot(block.Transactions) {
		return errors.New("invalid MerkleRoot")
	}
	// Проверка подписи всех транзакций
	for _, tx := range block.Transactions {
		if !wallet.VerifyTransaction(tx, nil) {
			return errors.New("block contains transaction with invalid signature")
		}
	}
	// Проверка уникальности TxID (replay protection)
	txidSet := make(map[string]struct{})
	for _, tx := range prevBlock.Transactions {
		txidSet[tx.TxID] = struct{}{}
	}
	for _, tx := range block.Transactions {
		if _, exists := txidSet[tx.TxID]; exists {
			return errors.New("duplicate TxID in block (replay attack)")
		}
	}
	newState := applyTransactions(prevState, block.Transactions)
	if block.Header.PostStateHash != hashState(newState, nil) {
		return errors.New("invalid StateRoot")
	}
	if block.Header.BlockID != hashBlockHeader(&block) {
		return errors.New("invalid BlockID")
	}
	// TODO: проверить подпись валидатора
	return nil
}

// AddBlockToChain добавляет блок в хранилище, обновляет состояние, mempool, сохраняет снапшот
func AddBlockToChain(block types.Block, chain *[]types.Block, state *map[string]uint64, mempool *[]types.Transaction, snapshotPath string) error {
	*chain = append(*chain, block)
	*state = applyTransactions(*state, block.Transactions)
	// Удаляем обработанные транзакции из mempool
	newMempool := make([]types.Transaction, 0)
	for _, tx := range *mempool {
		found := false
		for _, btx := range block.Transactions {
			if tx.TxID == btx.TxID {
				found = true
				break
			}
		}
		if !found {
			newMempool = append(newMempool, tx)
		}
	}
	*mempool = newMempool
	// Сохраняем снапшот
	snapshot := types.FullChainSnapshot{
		Blocks:    ToTypesBlocks(*chain),
		Mempool:   *mempool,
		State:     types.ChainState{PostStateHash: block.Header.PostStateHash},
		Timestamp: time.Now(),
	}
	return storage.FlushSnapshot(&snapshot)
}

// GetLatestBlock возвращает последний валидный блок
func GetLatestBlock(chain []types.Block) types.Block {
	if len(chain) == 0 {
		return types.Block{}
	}
	return chain[len(chain)-1]
}

// --- Вспомогательные функции ---

// calcMerkleRoot вычисляет MerkleRoot для списка транзакций
func calcMerkleRoot(txs []types.Transaction) string {
	hashes := make([][]byte, len(txs))
	for i, tx := range txs {
		h := sha256.Sum256([]byte(tx.TxID))
		hashes[i] = h[:]
	}
	for len(hashes) > 1 {
		next := [][]byte{}
		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				pair := append(hashes[i], hashes[i+1]...)
				h := sha256.Sum256(pair)
				next = append(next, h[:])
			} else {
				next = append(next, hashes[i])
			}
		}
		hashes = next
	}
	if len(hashes) == 0 {
		return ""
	}
	return hex.EncodeToString(hashes[0])
}

// hashState сериализует и хеширует state
func hashState(state map[string]uint64, extra map[string]uint64) string {
	m := map[string]uint64{}
	for k, v := range state {
		m[k] = v
	}
	for k, v := range extra {
		m[k] = v
	}
	b, _ := json.Marshal(m)
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

// hashBlockHeader сериализует и хеширует заголовок блока
func hashBlockHeader(b *types.Block) string {
	head := struct {
		BlockID       string
		Height        uint64
		Timestamp     time.Time
		PreviousHash  string
		PostStateHash string
		Tag           string
		TxRoot        string
		Proposer      string
		ValidatorSigs []types.ValidatorSig
	}{
		BlockID:       b.Header.BlockID,
		Height:        b.Header.Height,
		Timestamp:     b.Header.Timestamp,
		PreviousHash:  b.Header.PreviousHash,
		PostStateHash: b.Header.PostStateHash,
		Tag:           b.Header.Tag,
		TxRoot:        b.Header.TxRoot,
		Proposer:      b.Proposer,
		ValidatorSigs: b.ValidatorSigs,
	}
	data, _ := json.Marshal(head)
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// burnFee списывает комиссию с баланса отправителя (MVP: комиссия сгорает)
func burnFee(state map[string]uint64, from string, fee float64) bool {
	feeMicro := uint64(fee * 1_000_000)
	if state[from] < feeMicro {
		return false // недостаточно средств для комиссии
	}
	state[from] -= feeMicro
	// MVP: комиссия просто сгорает, можно логировать здесь
	return true
}

// applyTransactions применяет транзакции к state (MVP: копирует state, TODO: реализовать логику списания/зачисления)
func applyTransactions(state map[string]uint64, txs []types.Transaction) map[string]uint64 {
	newState := map[string]uint64{}
	for k, v := range state {
		newState[k] = v
	}
	for _, tx := range txs {
		// Проверка подписи транзакции
		if !wallet.VerifyTransaction(tx, nil) {
			continue // невалидная подпись — пропускаем
		}
		var transfer types.TransferArgs
		if err := json.Unmarshal(tx.Payload, &transfer); err == nil && transfer.From != "" && transfer.To != "" && transfer.Amount > 0 {
			fee := tx.Fee
			total := transfer.Amount + uint64(fee*1_000_000)
			if newState[transfer.From] < total {
				continue // недостаточно средств
			}
			newState[transfer.From] -= total
			newState[transfer.To] += transfer.Amount
			// fee сгорает (можно доработать для системного кошелька)
			// Для аудита: логировать списание комиссии здесь
			continue
		}
		// Для всех остальных транзакций списываем только комиссию
		from := tx.Origin
		if !burnFee(newState, from, tx.Fee) {
			continue // недостаточно средств для комиссии
		}
		// TODO: обработка payload для других типов транзакций/контрактов
	}
	return newState
}

// ToTypesBlocks преобразует []types.Block (ядро) в []types.Block (для снапшота)
func ToTypesBlocks(chain []types.Block) []types.Block {
	return chain
}

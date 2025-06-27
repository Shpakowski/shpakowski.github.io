package mempool

import (
	"mcp-chain/core/transaction"
	"mcp-chain/internal/config"
	"mcp-chain/types"
	"sort"
	"sync"
	"time"
)

// Mempool — интерфейс для очередей транзакций
// Позволяет расширять типы mempool без изменения кода ядра
// (например, обычные TX, смарт-контракты, future TX)
type Mempool interface {
	AddTx(tx *types.Transaction)
	RemoveExpired()
	GetTxs() []*types.Transaction
}

const TxTimeout = 10 * time.Minute // Таймаут до авто-майна блока

// TxQueue — in-memory очередь транзакций
// Сортировка по Fee (desc), Timestamp (asc)
type TxQueue struct {
	mu      sync.Mutex
	txQueue []txWithTime
	limit   int
}

type txWithTime struct {
	tx *transaction.Tx
	ts int64
}

// NewTxQueue создаёт очередь с лимитом
func NewTxQueue(cfg *config.NetworkConfig) *TxQueue {
	return &TxQueue{
		txQueue: make([]txWithTime, 0),
		limit:   cfg.MemPoolLimit,
	}
}

// Add добавляет транзакцию после валидации
func (q *TxQueue) Add(tx *transaction.Tx, cfg *config.NetworkConfig, pubKey []byte) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.txQueue) >= q.limit {
		return transaction.ErrTxInvalid("mempool full")
	}
	if err := transaction.ValidateTx(tx, cfg, pubKey); err != nil {
		return err
	}
	txTime := time.Now().Unix()
	q.txQueue = append(q.txQueue, txWithTime{tx, txTime})
	q.sortQueue()
	return nil
}

// sortQueue сортирует по Fee (desc), Timestamp (asc)
func (q *TxQueue) sortQueue() {
	sort.Slice(q.txQueue, func(i, j int) bool {
		if q.txQueue[i].tx.Fee == q.txQueue[j].tx.Fee {
			return q.txQueue[i].ts < q.txQueue[j].ts
		}
		return q.txQueue[i].tx.Fee > q.txQueue[j].tx.Fee
	})
}

// FlushBatch возвращает срез ≤ cfg.MaxTxPerBlock
func (q *TxQueue) FlushBatch(cfg *config.NetworkConfig) []*transaction.Tx {
	q.mu.Lock()
	defer q.mu.Unlock()
	batchSize := cfg.MaxTxPerBlock
	if len(q.txQueue) < batchSize {
		batchSize = len(q.txQueue)
	}
	batch := make([]*transaction.Tx, batchSize)
	for i := 0; i < batchSize; i++ {
		batch[i] = q.txQueue[i].tx
	}
	q.txQueue = q.txQueue[batchSize:]
	return batch
}

// TxMempool — in-memory очередь транзакций
// Реализация интерфейса Mempool
// Сортирует по комиссии (Fee), удаляет старые
type TxMempool struct {
	mu    sync.Mutex
	txs   []*types.Transaction
	times map[string]time.Time // ID -> время поступления
}

// NewTxMempool создаёт новую очередь
func NewTxMempool() *TxMempool {
	return &TxMempool{
		txs:   make([]*types.Transaction, 0),
		times: make(map[string]time.Time),
	}
}

// AddTx добавляет транзакцию в mempool
func (m *TxMempool) AddTx(tx *types.Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.txs = append(m.txs, tx)
	m.times[string(tx.ID)] = time.Now()
	m.sortByFee()
}

// sortByFee сортирует транзакции по комиссии (по убыванию)
func (m *TxMempool) sortByFee() {
	sort.Slice(m.txs, func(i, j int) bool {
		return m.txs[i].Fee > m.txs[j].Fee
	})
}

// RemoveExpired удаляет транзакции, которые висят дольше TxTimeout
func (m *TxMempool) RemoveExpired() {
	m.mu.Lock()
	defer m.mu.Unlock()
	now := time.Now()
	newTxs := make([]*types.Transaction, 0, len(m.txs))
	for _, tx := range m.txs {
		if now.Sub(m.times[string(tx.ID)]) < TxTimeout {
			newTxs = append(newTxs, tx)
		} else {
			delete(m.times, string(tx.ID))
		}
	}
	m.txs = newTxs
}

// GetTxs возвращает все актуальные транзакции
func (m *TxMempool) GetTxs() []*types.Transaction {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]*types.Transaction(nil), m.txs...)
}

// Clear полностью очищает mempool
func (m *TxMempool) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.txs = []*types.Transaction{}
	m.times = map[string]time.Time{}
}

var GlobalTxMempool *TxMempool

func init() {
	GlobalTxMempool = NewTxMempool()
}

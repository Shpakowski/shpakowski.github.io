package mempool

import (
	"mcp-chain/core/transaction"
	"mcp-chain/internal/config"
	"sync"
)

// SCTxQueue — очередь для контрактных транзакций (Payload != nil)
type SCTxQueue struct {
	mu      sync.Mutex
	txQueue []*transaction.Tx
	limit   int
}

// NewSCTxQueue создаёт очередь с лимитом
func NewSCTxQueue(cfg *config.NetworkConfig) *SCTxQueue {
	return &SCTxQueue{
		txQueue: make([]*transaction.Tx, 0),
		limit:   cfg.MemPoolLimit,
	}
}

// Add добавляет контрактную транзакцию
func (q *SCTxQueue) Add(tx *transaction.Tx) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.txQueue) >= q.limit {
		return transaction.ErrTxInvalid("sc mempool full")
	}
	if len(tx.Payload) == 0 {
		return transaction.ErrTxInvalid("not a contract tx")
	}
	q.txQueue = append(q.txQueue, tx)
	return nil
}

// FlushBatch возвращает срез ≤ cfg.MaxGasPerBlock (по сумме gas контрактов)
func (q *SCTxQueue) FlushBatch(cfg *config.NetworkConfig, gasOf func(*transaction.Tx) uint64) []*transaction.Tx {
	q.mu.Lock()
	defer q.mu.Unlock()
	var (
		batch  []*transaction.Tx
		gasSum uint64
	)
	for _, tx := range q.txQueue {
		gas := gasOf(tx)
		if gasSum+gas > uint64(cfg.MaxGasPerBlock) {
			break
		}
		batch = append(batch, tx)
		gasSum += gas
	}
	q.txQueue = q.txQueue[len(batch):]
	return batch
}

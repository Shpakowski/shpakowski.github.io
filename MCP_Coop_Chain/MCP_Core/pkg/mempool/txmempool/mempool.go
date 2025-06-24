// mempool.go - Transaction mempool public interface (see logic/ for business logic)
package txmempool

import (
	"github.com/mcpcoop/chain/pkg/types"
)

// txMempool holds pending standard transactions for block inclusion.
// See whitepaper section 4.1 (Transaction Pool).
type txMempool struct {
	// private fields (do not access directly)
	pending []types.Transaction
}

// AddTx adds a transaction to the mempool.
func (m *txMempool) AddTx(tx types.Transaction) error {
	m.pending = append(m.pending, tx)
	return nil
}

// RemoveTx removes a transaction from the mempool by ID.
func (m *txMempool) RemoveTx(txID string) error {
	// TODO: Implement removal logic by txID
	return nil
}

// GetPendingTxs returns all pending transactions.
func (m *txMempool) GetPendingTxs() []types.Transaction {
	return m.pending
}

// Size returns the number of pending transactions
func (m *txMempool) Size() int {
	return len(m.pending)
} 
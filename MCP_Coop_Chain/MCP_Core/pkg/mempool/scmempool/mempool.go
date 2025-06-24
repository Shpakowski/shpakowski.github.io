// mempool.go - Smart contract mempool public interface (see logic/ for business logic)
package scmempool

import (
	"github.com/mcpcoop/chain/pkg/types"
)

// scMempool holds pending smart contract transactions for block inclusion.
// See whitepaper section 4.2 (Smart Contract Pool).
type scMempool struct {
	pending []types.Transaction // or a more specific type for SC txs
}

// AddTx adds a smart contract transaction to the pool.
func (m *scMempool) AddTx(tx types.Transaction) error {
	m.pending = append(m.pending, tx)
	return nil
}

// GetPendingTxs returns all pending smart contract transactions.
func (m *scMempool) GetPendingTxs() []types.Transaction {
	return m.pending
}

// RemoveTx removes a smart contract transaction from the pool
func (m *scMempool) RemoveTx(txID string) error {
	// TODO: Implement removal logic for smart contract txs
	return nil
}

// Size returns the number of pending smart contract transactions
func (m *scMempool) Size() int {
	return len(m.pending)
} 
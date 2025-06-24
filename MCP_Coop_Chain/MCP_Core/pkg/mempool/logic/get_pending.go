package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
)

// GetPendingTxs returns all transactions ready for block proposal from TxMempool.
// Whitepaper: Section 4.1
// Logs retrieval.
func GetPendingTxs(m *types.TxMempool) []types.Transaction {
	fmt.Printf("[MEMPOOL] Retrieving pending transactions\n")
	return m.Pending
} 
package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
)

// GetPendingScTxs returns all smart contract transactions ready for block proposal from ScMempool.
// Whitepaper: Section 4.2
// Logs retrieval.
func GetPendingScTxs(m *types.ScMempool) []types.Transaction {
	fmt.Printf("[SCMEMPOOL] Retrieving pending smart contract transactions\n")
	return m.Pending
} 
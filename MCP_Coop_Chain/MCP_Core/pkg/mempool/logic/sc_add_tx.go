package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
)

// AddScTx adds a smart contract transaction to the ScMempool after validation.
// Whitepaper: Section 4.2 (Smart Contract Pool)
// Logs all add attempts and errors.
func AddScTx(m *types.ScMempool, tx types.Transaction) error {
	fmt.Printf("[SCMEMPOOL] Adding smart contract transaction: %+v\n", tx)
	// TODO: Validate and add logic
	m.Pending = append(m.Pending, tx)
	return nil
} 
package logic

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
)

// AddTx adds a transaction to the TxMempool after validation.
// Whitepaper: Section 4.1 (Transaction Pool), 4.3 (Validation)
// Logs all add attempts and errors.
func AddTx(m *types.TxMempool, tx types.Transaction) error {
	// TODO: Validate transaction (see validate.go)
	fmt.Printf("[MEMPOOL] Adding transaction: %+v\n", tx)
	m.Pending = append(m.Pending, tx)
	return nil
} 
package logic

import (
	"errors"
	"github.com/mcpcoop/chain/pkg/types"
)

// ProcessTransaction validates and adds a transaction to the mempool.
// Returns an error if the transaction is invalid or the sender has insufficient balance.
func ProcessTransaction(c *types.Chain, tx types.Transaction) error {
	if tx.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	if c.Balances[string(tx.From)] < tx.Amount {
		return errors.New("insufficient balance")
	}
	// Additional validation can be added here
	c.Mempool = append(c.Mempool, tx)
	return nil
} 
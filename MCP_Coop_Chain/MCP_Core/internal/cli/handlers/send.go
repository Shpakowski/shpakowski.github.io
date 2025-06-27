package handlers

import (
	"time"

	"github.com/mcpcoop/chain/pkg/chain"
	"github.com/mcpcoop/chain/pkg/types"
)

// HandleSend creates and processes a transaction using the chain service.
func HandleSend(c *types.Chain, from, to string, amount float64) error {
	tx := types.Transaction{
		From:      types.Address(from),
		To:        types.Address(to),
		Amount:    amount,
		Timestamp: time.Now(),
	}
	return chain.ProcessTransaction(c, tx)
}

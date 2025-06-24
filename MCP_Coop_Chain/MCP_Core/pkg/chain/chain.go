// chain.go - Blockchain core logic and storage
package chain

import (
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/chain/logic"
)

// NewChain creates a new blockchain with a genesis block
func NewChain() *types.Chain {
	return logic.NewChain()
}

// AddBlock adds a new block to the chain
func AddBlock(c *types.Chain, txs []types.Transaction) {
	logic.AddBlock(c, txs)
}

// Save saves the blockchain state to disk
func Save(c *types.Chain, filename string) error {
	return logic.SaveChain(c, filename)
}

// Load loads the blockchain state from disk
func Load(c *types.Chain, filename string) error {
	return logic.LoadChain(c, filename)
}

// UpdateBalances updates balances after a block or transaction
func UpdateBalances(c *types.Chain) {
	logic.UpdateBalances(c)
}

// ProcessTransaction processes a single transaction
func ProcessTransaction(c *types.Chain, tx types.Transaction) error {
	return logic.ProcessTransaction(c, tx)
}

// Height returns the current block height
func Height(c *types.Chain) int {
	return logic.BlockHeight(c)
} 
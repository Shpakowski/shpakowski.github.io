package logic

import "github.com/mcpcoop/chain/pkg/types"

// UpdateBalances recalculates all wallet balances from the blockchain.
func UpdateBalances(c *types.Chain) {
	balances := make(map[string]float64)
	for _, block := range c.Blocks {
		for _, tx := range block.Transactions {
			balances[string(tx.From)] -= tx.Amount
			balances[string(tx.To)] += tx.Amount
		}
	}
	c.Balances = balances
} 
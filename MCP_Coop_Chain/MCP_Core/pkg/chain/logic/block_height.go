package logic

import "github.com/mcpcoop/chain/pkg/types"

// BlockHeight returns the current block height of the chain.
func BlockHeight(c *types.Chain) int {
	return len(c.Blocks) - 1
} 
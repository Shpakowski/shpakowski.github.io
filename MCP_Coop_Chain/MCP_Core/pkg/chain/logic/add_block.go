package logic

import (
	"fmt"
	"time"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/crypto"
)

// AddBlock adds a new block with the given transactions to the chain.
// It clears the mempool and updates the chain's block list.
func AddBlock(c *types.Chain, txs []types.Transaction) {
	prev := c.Blocks[len(c.Blocks)-1]
	hashInput := fmt.Sprintf("%d%s", prev.BlockHeader.Index+1, time.Now().String())
	hash := crypto.HashData([]byte(hashInput))
	block := types.Block{
		BlockHeader: types.BlockHeader{
			Index:     prev.BlockHeader.Index + 1,
			Timestamp: time.Now(),
			PrevHash:  prev.BlockHeader.Hash,
			Hash:      types.Hash(fmt.Sprintf("%x", hash)),
		},
		Transactions: txs,
	}
	c.Blocks = append(c.Blocks, block)
	c.Mempool = []types.Transaction{}
	UpdateBalances(c)
} 
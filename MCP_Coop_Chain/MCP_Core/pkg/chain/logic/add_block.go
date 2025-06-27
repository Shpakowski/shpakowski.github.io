package logic

import (
	"fmt"
	"time"

	"github.com/mcpcoop/chain/pkg/crypto"
	"github.com/mcpcoop/chain/pkg/types"
)

// AddBlock adds a new block with the given transactions to the chain.
// It clears the mempool and updates the chain's block list.
func computeMerkleRoot(txs []types.Transaction) types.Hash {
	if len(txs) == 0 {
		return ""
	}
	hashes := make([][]byte, len(txs))
	for i, tx := range txs {
		hashes[i] = []byte(tx.Hash())
	}
	for len(hashes) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				combined := append(hashes[i], hashes[i+1]...)
				h := types.Hash(fmt.Sprintf("%x", crypto.HashData(combined)))
				nextLevel = append(nextLevel, []byte(h))
			} else {
				nextLevel = append(nextLevel, hashes[i])
			}
		}
		hashes = nextLevel
	}
	return types.Hash(fmt.Sprintf("%x", hashes[0]))
}

func AddBlock(c *types.Chain, txs []types.Transaction) {
	prev := c.Blocks[len(c.Blocks)-1]
	hashInput := fmt.Sprintf("%d%s", prev.BlockHeader.Index+1, time.Now().String())
	hash := crypto.HashData([]byte(hashInput))
	merkleRoot := computeMerkleRoot(txs)
	block := types.Block{
		BlockHeader: types.BlockHeader{
			Index:      prev.BlockHeader.Index + 1,
			Timestamp:  time.Now(),
			PrevHash:   prev.BlockHeader.Hash,
			Hash:       types.Hash(fmt.Sprintf("%x", hash)),
			MerkleRoot: merkleRoot,
		},
		Transactions: txs,
	}
	c.Blocks = append(c.Blocks, block)
	c.Mempool = []types.Transaction{}
	UpdateBalances(c)
}

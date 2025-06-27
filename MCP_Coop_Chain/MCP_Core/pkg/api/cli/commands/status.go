package commands

import (
	"fmt"

	"github.com/mcpcoop/chain/pkg/chain"
	"github.com/mcpcoop/chain/pkg/types"
)

// Status displays the current status of the blockchain node
func Status(c *types.Chain, args []string) {
	height := chain.Height(c)
	mempool := len(c.Mempool)
	var latestHash, merkleRoot string
	var health string = "healthy"
	if height >= 0 {
		latestBlock := c.Blocks[height]
		latestHash = string(latestBlock.Hash)
		merkleRoot = string(latestBlock.MerkleRoot)
	}
	fmt.Printf("Status: running\n")
	fmt.Printf("Chain height: %d\n", height)
	fmt.Printf("Latest block hash: %s\n", latestHash)
	fmt.Printf("Merkle root: %s\n", merkleRoot)
	fmt.Printf("Mempool transactions: %d\n", mempool)
	fmt.Printf("Node health: %s\n", health)
	// Uptime and other info can be added if needed
}

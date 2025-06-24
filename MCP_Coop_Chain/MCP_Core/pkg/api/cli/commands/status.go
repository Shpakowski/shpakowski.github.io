package commands

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/pkg/chain"
)

// Status displays the current status of the blockchain node
func Status(c *types.Chain, args []string) {
	height := chain.Height(c)
	mempool := len(c.Mempool)
	fmt.Printf("Status: running\n")
	fmt.Printf("Chain height: %d\n", height)
	fmt.Printf("Mempool transactions: %d\n", mempool)
	// Uptime and other info can be added if needed
} 
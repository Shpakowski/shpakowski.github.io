package commands

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
)

// Stop stops the blockchain node
func Stop(c *types.Chain, args []string) {
	fmt.Printf("[INFO] Node stopped\n")
	stopBlockTimer()
} 
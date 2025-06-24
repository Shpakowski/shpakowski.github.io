package commands

import (
	"fmt"
	"github.com/mcpcoop/chain/pkg/types"
)

// Restart restarts the blockchain node
func Restart(c *types.Chain, args []string) {
	fmt.Printf("[INFO] Restarting node...\n")
	Stop(c, args)
	Start(c, args)
} 
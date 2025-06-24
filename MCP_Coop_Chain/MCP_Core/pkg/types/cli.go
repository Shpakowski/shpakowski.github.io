// cli.go - CLI state and status types for MCP Coop Chain
package types

import (
	"time"
)

// nodeState - full persistent state for CLI
// Includes: status, chain, mempool, wallets, balances
//
type NodeState struct {
	Status      string                `json:"status"`
	Chain       []Block               `json:"chain"`
	Mempool     []Transaction         `json:"mempool"`
	Wallets     []Wallet              `json:"wallets"`
	Balances    map[string]float64    `json:"balances"`
	StartTime   string                `json:"start_time"`
}

// nodeStatus - running status for CLI
//
type NodeStatus struct {
	Running   bool
	StartTime time.Time
} 
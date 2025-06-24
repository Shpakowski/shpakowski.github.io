// chain.go - Blockchain state type for MCP Coop Chain
package types

// Chain represents the blockchain and its state
//
type Chain struct {
	Blocks    []Block
	Balances  map[string]float64
	Mempool   []Transaction
} 
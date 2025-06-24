// block.go - Block struct and related definitions for MCP Coop Chain
package types

import "time"

// BlockHeader defines the minimal header fields for a block.
// Used for block validation and chain linking.
type BlockHeader struct {
	Index     int       `json:"index"`
	Timestamp time.Time `json:"timestamp"`
	PrevHash  Hash      `json:"prev_hash"`
	Hash      Hash      `json:"hash"`
}

// Block represents a full block in the blockchain.
// Contains header and all included transactions.
type Block struct {
	BlockHeader
	Transactions []Transaction `json:"transactions"`
} 
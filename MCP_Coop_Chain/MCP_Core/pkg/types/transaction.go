// transaction.go - Transaction struct and related definitions for MCP Coop Chain
package types

import (
	"time"
)

// TransactionType enumerates all supported transaction types.
type TransactionType int

const (
	TxTypeTransfer TransactionType = iota // Standard value transfer
	TxTypeContractCall                    // Smart contract call
	TxTypeContractDeploy                  // Smart contract deployment
)

// Transaction represents a blockchain transaction.
type Transaction struct {
	From      Address         `json:"from"`
	To        Address         `json:"to"`
	Amount    float64         `json:"amount"`
	Type      TransactionType `json:"type"`
	Payload   []byte          `json:"payload,omitempty"` // For contract calls/deploys
	Timestamp time.Time       `json:"timestamp"`
	Signature Signature       `json:"signature"`
} 
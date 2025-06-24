// wallet.go - Wallet type for MCP Coop Chain
package types

import "crypto/ed25519"

// Wallet represents a user wallet (address, public key, etc).
type Wallet struct {
	Address  Address            `json:"address"`
	PubKey   ed25519.PublicKey  `json:"pubkey"`
	Balance  float64            `json:"balance"`
}

// Metadata holds tags, aliases, and other user-defined info for a wallet
// Not serialized on-chain, only for local management
//
type Metadata struct {
	Alias string   `json:"alias,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	Notes string   `json:"notes,omitempty"`
} 
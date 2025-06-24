// crypto.go - Cryptographic types for MCP Coop Chain
package types

import "crypto/ed25519"

// KeyPair represents a public/private key pair
//
type KeyPair struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
} 
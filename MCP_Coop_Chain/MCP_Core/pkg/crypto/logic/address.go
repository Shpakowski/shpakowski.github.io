package logic

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"crypto/ed25519"
)

// AddressFromPubKey derives a wallet address from a public key.
// Business requirement: Unique, deterministic address for each wallet (whitepaper section 2.3).
// Logs address derivation for monitoring.
func AddressFromPubKey(pub ed25519.PublicKey) string {
	hash := sha256.Sum256(pub)
	address := hex.EncodeToString(hash[:8]) // Use first 8 bytes for address (mock)
	fmt.Printf("[CRYPTO] Address derived from public key: %s\n", address)
	return address
} 
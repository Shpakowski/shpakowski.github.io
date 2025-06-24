package logic

import (
	"crypto/sha256"
	"fmt"
)

// HashData hashes arbitrary data using SHA-256.
// Business requirement: All transactions and blocks must be hashed for integrity and uniqueness (whitepaper section 3.3).
// Logs hashing for monitoring.
func HashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	fmt.Printf("[CRYPTO] Data hashed\n")
	return hash[:]
} 
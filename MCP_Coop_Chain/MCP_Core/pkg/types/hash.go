// hash.go - Hash type and helpers for MCP Coop Chain
package types

import "encoding/hex"

// Hash represents a block or transaction hash (32 bytes, hex-encoded).
type Hash string

// IsValidHash checks if a string is a valid 32-byte hex hash.
func IsValidHash(h string) bool {
	if len(h) != 64 {
		return false
	}
	_, err := hex.DecodeString(h)
	return err == nil
} 
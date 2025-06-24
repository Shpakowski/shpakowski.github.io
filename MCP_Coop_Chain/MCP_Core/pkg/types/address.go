// address.go - Address type and helpers for MCP Coop Chain
package types

import "encoding/hex"

// Address represents a wallet or contract address (hex-encoded string).
type Address string

// IsValidAddress checks if a string is a valid address (length and hex).
func IsValidAddress(a string) bool {
	if len(a) < 16 || len(a) > 64 {
		return false
	}
	_, err := hex.DecodeString(a)
	return err == nil
} 
// signature.go - Signature type and helpers for MCP Coop Chain
package types

// Signature represents a digital signature (hex-encoded string).
type Signature string

// IsValidSignature checks if a string is a valid signature (length, hex, etc).
func IsValidSignature(s string) bool {
	return len(s) > 0 // Add more checks as needed
} 
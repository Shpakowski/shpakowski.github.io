// errors.go - Custom error types for MCP Coop Chain
package types

import "errors"

var (
	ErrInvalidAddress   = errors.New("invalid address")
	ErrInvalidHash      = errors.New("invalid hash")
	ErrInvalidSignature = errors.New("invalid signature")
	ErrInvalidTx        = errors.New("invalid transaction")
	ErrInsufficientFunds = errors.New("insufficient funds")
) 
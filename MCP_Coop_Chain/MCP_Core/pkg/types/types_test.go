// types_test.go - Tests for MCP Coop Chain types
package types

import "testing"

func TestIsValidHash(t *testing.T) {
	if !IsValidHash("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef") {
		t.Error("Expected valid hash")
	}
	if IsValidHash("notavalidhash") {
		t.Error("Expected invalid hash")
	}
}

func TestIsValidAddress(t *testing.T) {
	if !IsValidAddress("0123456789abcdef01234567") {
		t.Error("Expected valid address")
	}
	if IsValidAddress("badaddress") {
		t.Error("Expected invalid address")
	}
}

func TestInit(t *testing.T) {} 
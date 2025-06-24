// cli_test.go - Tests for CLI commands and handlers
package cli

import (
	"strings"
	"testing"
)

func TestInit(t *testing.T) {}

func TestDeterministicWalletCreation(t *testing.T) {
	seed := "alpha beta gamma delta epsilon zeta eta theta iota kappa lambda mu"
	args := []string{seed}
	NewWalletCmd(args)
	addr1 := wallets[len(wallets)-1]
	NewWalletCmd(args)
	addr2 := wallets[len(wallets)-1]
	if addr1 != addr2 {
		t.Errorf("Expected deterministic address, got %s and %s", addr1, addr2)
	}
}

func TestSendUpdatesBalances(t *testing.T) {
	from := "abcdef01"
	to := "12345678"
	mockBalances[from] = 100.0
	mockBalances[to] = 10.0
	SendCmd([]string{from, to, "25.5"})
	if mockBalances[from] != 74.5 {
		t.Errorf("Expected from balance 74.5, got %f", mockBalances[from])
	}
	if mockBalances[to] != 35.5 {
		t.Errorf("Expected to balance 35.5, got %f", mockBalances[to])
	}
} 
// import.go - Import wallets from file(s) for MCP Coop Chain
// Business purpose: User wallet recovery and migration (see whitepaper section 2.2)
package wallet

import (
	"encoding/json"
	"os"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/internal/logger"
)

// ImportWalletsFromFile loads wallets from a JSON file
func ImportWalletsFromFile(filename string) ([]*types.Wallet, error) {
	logger.Logger.Info("Importing wallets from file", "file", filename)
	f, err := os.Open(filename)
	if err != nil {
		logger.Logger.Error("Failed to open wallet file", "error", err)
		return nil, err
	}
	defer f.Close()
	var wallets []*types.Wallet
	dec := json.NewDecoder(f)
	if err := dec.Decode(&wallets); err != nil {
		logger.Logger.Error("Failed to decode wallet file", "error", err)
		return nil, err
	}
	return wallets, nil
} 
// export.go - Export wallets to file(s) for MCP Coop Chain
// Business purpose: User wallet backup and migration (see whitepaper section 2.2)
package wallet

import (
	"encoding/json"
	"os"
	"github.com/mcpcoop/chain/pkg/types"
	"github.com/mcpcoop/chain/internal/logger"
)

// ExportWalletsToFile saves wallets to a JSON file
func ExportWalletsToFile(wallets []*types.Wallet, filename string) error {
	logger.Logger.Info("Exporting wallets to file", "file", filename)
	f, err := os.Create(filename)
	if err != nil {
		logger.Logger.Error("Failed to create wallet file", "error", err)
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(wallets); err != nil {
		logger.Logger.Error("Failed to encode wallet file", "error", err)
		return err
	}
	return nil
} 
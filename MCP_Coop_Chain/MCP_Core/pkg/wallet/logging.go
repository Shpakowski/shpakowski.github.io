// logging.go - Logging for wallet operations in MCP Coop Chain
// Business purpose: Auditability and traceability of wallet actions (see whitepaper section 5.2)
package wallet

import "github.com/mcpcoop/chain/internal/logger"

// LogWalletEvent logs a wallet operation
func LogWalletEvent(event string, details ...interface{}) {
	logger.Logger.Info("[WALLET] "+event, details...)
}

// LogWalletError logs a wallet error
func LogWalletError(event string, err error) {
	logger.Logger.Error("[WALLET] "+event, "error", err)
} 
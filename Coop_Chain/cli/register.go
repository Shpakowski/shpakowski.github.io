package cli

import (
	"mcp-chain/core/blockchain"
	"mcp-chain/core/wallet"
	"mcp-chain/internal/config"
	"path/filepath"

	"github.com/spf13/cobra"
)

// RegisterCommands регистрирует все подкоманды CLI в rootCmd
func RegisterCommands(rootCmd *cobra.Command, chain **blockchain.Chain, cfg *config.NetworkConfig, startTime *int64) {
	// Регистрируем blockchain-команду с передачей глобального состояния
	walletsPath := filepath.Join("data", "wallets.json")
	manager := wallet.NewManager(walletsPath)
	rootCmd.AddCommand(InitBlockchainCmd(chain, cfg, startTime))

	// Регистрируем wallet-команду с in-memory manager
	rootCmd.AddCommand(InitWalletCmd(manager, cfg))

	// Регистрируем stake-команду с in-memory state
	rootCmd.AddCommand(InitStakeCmd(cfg))

	// Регистрируем transaction-команду с in-memory state
	rootCmd.AddCommand(InitTransactionCmd(chain, cfg))

	// Регистрируем contracts-команду с in-memory state
	rootCmd.AddCommand(InitContractsCmd(chain, cfg))

	// TODO: Передавать chain, cfg, startTime в команды
	rootCmd.AddCommand(nodeCmd)
}

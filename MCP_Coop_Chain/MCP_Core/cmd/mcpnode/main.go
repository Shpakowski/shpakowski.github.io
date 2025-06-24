// main.go - Entry point for MCP Coop Chain node CLI
package main

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/mcpcoop/chain/internal/logger"
	"github.com/mcpcoop/chain/pkg/api/cli"
)

func main() {
	logger.Init()

	var seed string
	var from, to string
	var amount float64

	rootCmd := &cobra.Command{
		Use:   "mcpnode",
		Short: "MCP Coop Chain node CLI",
	}

	newWalletCmd := &cobra.Command{
		Use:   "new-wallet",
		Short: "Create a new wallet",
		Run: func(cmd *cobra.Command, args []string) {
			cli.NewWalletCmd([]string{seed})
		},
	}
	newWalletCmd.Flags().StringVar(&seed, "seed", "", "12-word seed phrase (space-separated)")

	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "Send a transaction",
		Run: func(cmd *cobra.Command, args []string) {
			cli.SendCmd([]string{from, to, cmd.Flag("amount").Value.String()})
		},
	}
	sendCmd.Flags().StringVar(&from, "from", "", "Sender address")
	sendCmd.Flags().StringVar(&to, "to", "", "Receiver address")
	sendCmd.Flags().Float64Var(&amount, "amount", 0, "Amount to send")

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "start",
			Short: "Start the node",
			Run:   func(cmd *cobra.Command, args []string) { cli.StartCmd(args) },
		},
		&cobra.Command{
			Use:   "stop",
			Short: "Stop the node",
			Run:   func(cmd *cobra.Command, args []string) { cli.StopCmd(args) },
		},
		&cobra.Command{
			Use:   "restart",
			Short: "Restart the node",
			Run:   func(cmd *cobra.Command, args []string) { cli.RestartCmd(args) },
		},
		newWalletCmd,
		sendCmd,
		&cobra.Command{
			Use:   "status",
			Short: "Show node status",
			Run:   func(cmd *cobra.Command, args []string) { cli.StatusCmd(args) },
		},
	)

	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Error("CLI error", "error", err)
		os.Exit(1)
	}
} 
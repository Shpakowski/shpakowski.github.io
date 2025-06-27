// main.go - Entry point for MCP Coop Chain node CLI
package main

import (
	"fmt"
	"os"

	"github.com/mcpcoop/chain/internal/logger"
	"github.com/mcpcoop/chain/pkg/api/cli"
	"github.com/spf13/cobra"
)

func main() {
	logger.Init()

	var seed string
	var from, to string
	var amount float64
	var blockIndex int
	var txIndex int
	var proof []string
	var txHash string

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

	merkleProofCmd := &cobra.Command{
		Use:   "merkle-proof",
		Short: "Generate Merkle proof for a transaction in a block",
		Run: func(cmd *cobra.Command, args []string) {
			c := cli.LoadChainState()
			proof, err := cli.GenerateMerkleProof(c, blockIndex, txIndex)
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err)
				return
			}
			fmt.Printf("Merkle proof: %v\n", proof)
		},
	}
	merkleProofCmd.Flags().IntVar(&blockIndex, "block", 0, "Block index")
	merkleProofCmd.Flags().IntVar(&txIndex, "tx", 0, "Transaction index in block")

	verifyProofCmd := &cobra.Command{
		Use:   "verify-proof",
		Short: "Verify Merkle proof for a transaction in a block",
		Run: func(cmd *cobra.Command, args []string) {
			c := cli.LoadChainState()
			ok, err := cli.VerifyMerkleProof(c, blockIndex, txHash, proof)
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err)
				return
			}
			if ok {
				fmt.Println("Proof is valid.")
			} else {
				fmt.Println("Proof is invalid.")
			}
		},
	}
	verifyProofCmd.Flags().IntVar(&blockIndex, "block", 0, "Block index")
	verifyProofCmd.Flags().StringVar(&txHash, "txhash", "", "Transaction hash")
	verifyProofCmd.Flags().StringSliceVar(&proof, "proof", nil, "Merkle proof (comma-separated)")

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
		merkleProofCmd,
		verifyProofCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Error("CLI error", "error", err)
		os.Exit(1)
	}
}

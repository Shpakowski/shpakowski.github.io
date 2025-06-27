package cli

import (
	"fmt"
	"mcp-chain/core/blockchain"
	"mcp-chain/internal/config"
	"mcp-chain/types"
	"strconv"

	"mcp-chain/core/services"

	"github.com/spf13/cobra"
)

// Инициализация transaction-команд с in-memory state и chain
func InitTransactionCmd(chain **blockchain.Chain, cfg *config.NetworkConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [from_seed|pubkey] [to_address] [amount] [fee]",
		Short: "Отправить транзакцию (from, to, amount, [fee])",
		Args:  cobra.RangeArgs(3, 4),
		Run: func(cmd *cobra.Command, args []string) {
			fromInput := args[0]
			to := types.Address(args[1])
			amount, err1 := strconv.ParseUint(args[2], 10, 64)
			fee := uint64(0)
			var err2 error
			if len(args) > 3 {
				fee, err2 = strconv.ParseUint(args[3], 10, 64)
			}
			if err1 != nil || (len(args) > 3 && err2 != nil) {
				fmt.Println("Некорректные аргументы amount или fee")
				return
			}
			txID, err := services.SendTxAndSave(*chain, cfg, fromInput, to, amount, fee)
			if err != nil {
				fmt.Println("Ошибка отправки транзакции:", err)
				return
			}
			fmt.Printf("Транзакция отправлена! ID: %s\n", txID)
		},
	}
	return cmd
}

func init() {
	// Для регистрации: rootCmd.AddCommand(transactionCmd)
}

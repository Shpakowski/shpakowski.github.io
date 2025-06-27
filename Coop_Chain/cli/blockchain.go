package cli

import (
	"fmt"
	"time"

	"mcp-chain/core/blockchain"
	"mcp-chain/core/state"
	"mcp-chain/core/wallet"
	"mcp-chain/internal/config"

	"github.com/spf13/cobra"
)

// blockchainCmd — корневая команда для работы с блокчейном.
// Содержит подкоманды для просмотра статуса, списка блоков и проверки целостности цепи.
var blockchainCmd = &cobra.Command{
	Use:   "chain",
	Short: "Операции с блокчейном (статус, список, проверка)",
}

// Инициализация blockchain-команд с in-memory state
func InitBlockchainCmd(chain **blockchain.Chain, cfg *config.NetworkConfig, startTime *int64) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain",
		Short: "Операции с блокчейном (статус, список, проверка)",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "head",
		Short: "Показать статус сети (количество блоков, высота, баланс и др.)",
		Run: func(cmd *cobra.Command, args []string) {
			if *startTime == 0 {
				*startTime = time.Now().Unix()
			}
			stats := blockchain.GetChainStats(*chain, time.Unix(*startTime, 0))
			fmt.Println("===== СТАТУС СЕТИ =====")
			fmt.Printf("Блоков в цепи: %d\n", stats.BlockCount)
			fmt.Printf("Кошельков: %d\n", stats.WalletCount)
			fmt.Printf("Валидаторов: %d\n", stats.ValidatorCount)
			fmt.Printf("Общий баланс: %.2f MCP\n", stats.TotalBalance)
			fmt.Printf("Высота последнего блока: %d\n", stats.LastBlockHeight)
			fmt.Printf("Хэш последнего блока: %s\n", stats.LastBlockHash)
			fmt.Printf("Время генерации последнего блока: %v\n", stats.LastBlockTime)
			fmt.Printf("Адрес последнего валидатора: %s\n", stats.LastValidator)
			fmt.Printf("Аптайм узла: %s\n", stats.Uptime)
			fmt.Println("========================")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "ls",
		Short: "Показать список блоков",
		Run: func(cmd *cobra.Command, args []string) {
			infos := blockchain.BlocksInfo()
			fmt.Println("===== СПИСОК БЛОКОВ =====")
			for _, b := range infos {
				fmt.Printf("Блок #%d | Хэш: %s | Время: %v | Валидатор: %s\n",
					b.Height,
					b.Hash,
					b.Timestamp,
					b.Proposer)
			}
			fmt.Println("==========================")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "verify",
		Short: "Проверить целостность цепи",
		Run: func(cmd *cobra.Command, args []string) {
			if (*chain).VerifyChain() {
				fmt.Println("Цепь валидна: OK")
			} else {
				fmt.Println("Цепь повреждена: FAIL")
			}
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "state-summary",
		Short: "Показать сводку по state (кошельки, балансы, валидаторы)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("===== СВОДКА STATE =====")
			wallets := wallet.ListWalletsFromState()
			fmt.Printf("Кошельков: %d\n", len(wallets))
			total := 0.0
			for _, w := range wallets {
				fmt.Printf("Адрес: %s | Баланс: %.2f MCP\n", w.Address, w.Balance)
				total += w.Balance
			}
			fmt.Printf("Общий баланс: %.2f MCP\n", total)
			fmt.Printf("Валидаторов: %d\n", len(state.GlobalState.Validators))
			if len(state.GlobalState.Validators) > 0 {
				fmt.Println("Список валидаторов:")
				for addr := range state.GlobalState.Validators {
					fmt.Println("  ", addr)
				}
			}
		},
	})

	return cmd
}

func init() {
	// Для регистрации: rootCmd.AddCommand(blockchainCmd)
}

package cli

import (
	"fmt"
	"mcp-chain/core/wallet"
	"mcp-chain/internal/config"
	"mcp-chain/logging"

	"github.com/spf13/cobra"
)

// Инициализация wallet-команд с in-memory manager и state
func InitWalletCmd(manager *wallet.Manager, cfg *config.NetworkConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wallet",
		Short: "Операции с кошельком (импорт, список, баланс)",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "import <seed>",
		Short: "Импортировать кошелек по seed-фразе (и получить награду, если новый)",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Укажите seed-фразу для импорта")
				return
			}
			entry, err := wallet.ImportWalletAndSave(cfg, args[0])
			if err != nil {
				fmt.Println("Ошибка импорта:", err)
				return
			}
			fmt.Printf("Кошелек импортирован! Публичный ключ: %s\n", entry.PubKey)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Показать список всех кошельков (по state)",
		Run: func(cmd *cobra.Command, args []string) {
			wallets := wallet.ListWalletsFromState()
			fmt.Printf("Всего кошельков: %d\n", len(wallets))
			for _, w := range wallets {
				fmt.Printf("Адрес: %s | Баланс: %.2f MCP\n", w.Address, w.Balance)
			}
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "balance [address|pubkey]",
		Short: "Показать баланс кошелька (по адресу/публичному ключу или всех)",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				balances := wallet.GetWalletBalances(manager)
				logging.Logger.Info("Балансы всех кошельков", "count", len(balances))
				if len(balances) == 0 {
					fmt.Println("Нет кошельков для отображения баланса.")
					return
				}
				fmt.Println("Балансы всех кошельков:")
				for _, b := range balances {
					fmt.Printf("%s: %.2f MCP\n", b.Address, b.Balance)
				}
			} else {
				b := wallet.GetWalletBalance(args[0])
				logging.Logger.Info("Баланс кошелька", "address", b.Address, "balance", b.Balance)
				fmt.Printf("Баланс %s: %.2f MCP\n", b.Address, b.Balance)
			}
		},
	})

	return cmd
}

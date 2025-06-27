package cli

import (
	"fmt"
	"mcp-chain/core/staking"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
	"strconv"

	"github.com/spf13/cobra"
)

// Инициализация stake-команд с in-memory state
func InitStakeCmd(cfg *config.NetworkConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake",
		Short: "Операции со стейкингом (lock, info)",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "lock [address|pubkey|seed] [amount] [duration_sec]",
		Short: "Залочить токены на адресе (аргументы: адрес/публичный ключ/seed, сумма, срок в секундах)",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				fmt.Println("Использование: stake lock [address|pubkey|seed] [amount] [duration_sec]")
				return
			}
			amt, err1 := strconv.ParseUint(args[1], 10, 64)
			dur, err2 := strconv.ParseInt(args[2], 10, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректные аргументы amount или duration_sec")
				return
			}
			addr, until, err := staking.LockStakeWithInputAndSave(cfg, args[0], amt, dur)
			if err != nil {
				fmt.Println("Ошибка залочки:", err)
				return
			}
			fmt.Printf("Токены залочены! Адрес: %s, Сумма: %d, До: %d (unix)\n", addr, amt, until)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "info [address|pubkey|seed]",
		Short: "Показать информацию о стейке для адреса/публичного ключа/seed",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Использование: stake info [address|pubkey|seed]")
				return
			}
			addr, info, err := staking.GetStakeInfoByInput(args[0])
			if err != nil {
				logging.Logger.Error("Ошибка получения информации о стейке", "err", err)
				fmt.Println("Ошибка получения информации о стейке:", err)
				return
			}
			logging.Logger.Info("Информация о стейке", "address", addr, "amount", info.Amount, "until", info.LockedUntil, "slash", info.SlashCount)
			fmt.Printf("Стейк для %s:\n  Сумма: %d\n  До: %d (unix)\n  Штрафы: %d\n", addr, info.Amount, info.LockedUntil, info.SlashCount)
		},
	})

	return cmd
}

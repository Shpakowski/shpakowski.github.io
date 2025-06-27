package cli

import (
	"fmt"
	"mcp-chain/contracts"
	"mcp-chain/core/blockchain"
	"mcp-chain/internal/config"
	"mcp-chain/logging"

	"github.com/spf13/cobra"
)

// Инициализация contracts-команд с in-memory state
func InitContractsCmd(chain **blockchain.Chain, cfg *config.NetworkConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract",
		Short: "Операции с контрактами (deploy, call)",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "deploy <code> [args...]",
		Short: "Деплоить пользовательский контракт (code, args)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			deployArgs := []string{}
			if len(args) > 1 {
				deployArgs = args[1:]
			}
			addr, txID, err := contracts.DeployContractAndSave(cfg, code, deployArgs...)
			if err != nil {
				fmt.Println("Ошибка деплоя:", err)
				return
			}
			fmt.Printf("Контракт задеплоен! Адрес: %s, TxID: %s\n", addr, txID)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "call <address> <method> [args...]",
		Short: "Вызвать метод пользовательского контракта (address, method, args)",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0]
			method := args[1]
			callArgs := []string{}
			if len(args) > 2 {
				callArgs = args[2:]
			}
			result, err := contracts.CallUserContract(address, method, callArgs...)
			if err != nil {
				logging.Logger.Error("Ошибка вызова пользовательского контракта", "err", err)
				fmt.Println("Ошибка вызова:", err)
				return
			}
			logging.Logger.Info("Вызов пользовательского контракта", "address", address, "method", method, "args", callArgs, "result", result)
			fmt.Printf("Результат вызова: %s\n", result)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "proto-call <name> [key=value ...]",
		Short: "Вызвать протокольный (встроенный) контракт (name, key=value ...)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			params := map[string]interface{}{}
			for _, kv := range args[1:] {
				var k, v string
				sep := '='
				for i, c := range kv {
					if c == sep {
						k = kv[:i]
						v = kv[i+1:]
						break
					}
				}
				if k != "" {
					params[k] = v
				}
			}
			result, err := contracts.CallProtoContract(name, params)
			if err != nil {
				logging.Logger.Error("Ошибка вызова протокольного контракта", "err", err)
				fmt.Println("Ошибка вызова протокольного контракта:", err)
				return
			}
			logging.Logger.Info("Вызов протокольного контракта", "name", name, "params", params, "result", result)
			fmt.Printf("Результат вызова: %s\n", result)
		},
	})

	return cmd
}

func init() {
	// Для регистрации: rootCmd.AddCommand(contractsCmd)
}

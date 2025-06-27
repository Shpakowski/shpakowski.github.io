package cli

import (
	"fmt"

	"mcp-chain/core/blockchain"

	"github.com/spf13/cobra"
)

// nodeCmd — команда для запуска полноценного узла блокчейна.
// Запускает цепочку, майнит новые блоки по правилам из конфига (AutoMineInterval, MinStake и др.).
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Запустить полноценный блокчейн-узел (майнинг, ядро)",
	Run: func(cmd *cobra.Command, args []string) {
		blockchain.RunNode(func(format string, a ...any) {
			fmt.Printf(format+"\n", a...)
		})
	},
}

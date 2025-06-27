package main

import (
	"fmt"
	"os"
	"time"

	"mcp-chain/api"
	"mcp-chain/cli"
	"mcp-chain/core/blockchain"
	"mcp-chain/core/state"
	"mcp-chain/internal/config"
	"mcp-chain/logging"

	"github.com/spf13/cobra"
)

// Глобальные переменные для флагов
var (
	dataDir  string // Путь к директории данных
	logLevel string // Уровень логирования

	// Глобальные переменные для in-memory state
	chain     *blockchain.Chain
	cfg       config.NetworkConfig
	startTime int64
)

// Корневая команда CLI
var rootCmd = &cobra.Command{
	Use:   "mcp",
	Short: "MCP Chain CLI — управление узлом и кошельком",
	Long:  `Командная строка для управления блокчейном MCP Chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Если не указана подкоманда — выводим справку
		_ = cmd.Help()
	},
}

func init() {
	// Глобальные флаги
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "./data", "Путь к директории данных")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Уровень логирования (info|debug|error)")

	// Регистрация команд с передачей глобального состояния
	cli.RegisterCommands(rootCmd, &chain, &cfg, &startTime)
}

func main() {
	// Инициализация логгера
	logPath := dataDir + "/logs/mcp.log"
	logging.InitLogger(logLevel, logPath)

	// Инициализация глобального состояния: всегда грузим state из файла ОДИН РАЗ при старте
	st, err := state.LoadStateFromFile(config.StateFile)
	if err != nil {
		if os.IsNotExist(err) {
			logging.Logger.Warn("state.json не найден, создаю новый state")
			state.GlobalState = state.NewState()
			err2 := state.ApplyStateChange(func(st *state.State) {})
			if err2 != nil {
				logging.Logger.Error("Ошибка при сохранении нового state", "err", err2)
			}
		} else {
			logging.Logger.Error("Ошибка загрузки state, state НЕ пересоздаётся!", "err", err)
			panic(fmt.Sprintf("Ошибка загрузки state: %v", err))
		}
	} else {
		state.GlobalState = st
		logging.Logger.Info("state успешно загружен из файла при инициализации CLI/ноды")
	}
	chain = blockchain.NewChain()
	cfg = config.GetDefaultNetworkConfig()
	startTime = time.Now().Unix()

	// Запуск API сервера только если команда node
	if len(os.Args) > 1 && os.Args[1] == "node" {
		go func() {
			// Порт берём из node config
			apiPort := config.GetDefaultNodeConfig().APIPort
			// Передаём chain и cfg
			api.StartAPIServer(apiPort, chain, &cfg)
		}()
	}

	if err := rootCmd.Execute(); err != nil {
		logging.Logger.Error("Ошибка выполнения CLI", "err", err)
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}

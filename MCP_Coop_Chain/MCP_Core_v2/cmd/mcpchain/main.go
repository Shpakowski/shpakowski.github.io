package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"mcp-coop-chain/internal"
	"mcp-coop-chain/internal/blockchain"
	"mcp-coop-chain/internal/logging"
	"mcp-coop-chain/internal/storage"
	"mcp-coop-chain/internal/types"

	"go.uber.org/zap"
)

// main — точка входа для запуска MCP Coop Chain FullNode
func main() {
	fmt.Println("MCP Coop Chain node starting...")

	// Загружаем конфиг сети через internal.LoadConfig
	cfg, err := internal.LoadConfig("configs/blockchain_config.json")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация логгера через DI
	logger, err := logging.NewLogger(cfg.Logger)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	logger.Info("Конфигурация сети загружена")

	// Запуск через boot: либо загружаем, либо создаём полный снимок
	snapshotPath := "data/full_state_snapshot.json"
	var snapshot *types.FullChainSnapshot
	var restoreErr error
	snapshot, restoreErr = storage.LoadStateFromDisk(snapshotPath)
	if restoreErr != nil {
		if os.IsNotExist(restoreErr) {
			fmt.Println("✔ Snapshot не найден, создаём genesis...")
			snapshot, restoreErr = storage.CreateGenesisSnapshot(snapshotPath)
			if restoreErr != nil {
				logger.Error("Ошибка создания genesis snapshot", restoreErr)
				return
			}
		} else {
			logger.Error("Ошибка восстановления состояния", restoreErr)
			return
		}
	}
	logger.Info("Снимок состояния сети загружен", zap.String("path", snapshotPath))
	logger.Info("Структура снимка:", zap.Any("snapshot", snapshot))

	memStorage := storage.NewMemoryStorage()
	var chain []*types.Block
	var mempool []*types.Transaction
	if err := storage.ApplyRestoredState(snapshot, memStorage, &chain, &mempool); err != nil {
		logger.Error("Ошибка применения восстановленного состояния", err)
		return
	}

	node := blockchain.NewFullNode(cfg, memStorage, logger)
	node.Chain = chain
	node.Mempool = mempool

	logger.Info("✔ Состояние блокчейна успешно восстановлено",
		zap.Int("blocks", len(chain)),
		zap.Int("wallets", len(snapshot.Wallets)),
		zap.Int("mempool", len(snapshot.Mempool)),
		zap.Any("lastBlock", func() interface{} {
			if len(chain) > 0 {
				return chain[len(chain)-1].Header
			}
			return nil
		}()),
	)

	node.Start()
	logger.Info("FullNode запущен, блоки будут создаваться даже без транзакций (allowEmptyBlocks=true)")

	// Для теста: логируем структуру цепочки каждые 10 секунд, останавливаем через 35 секунд
	for i := 0; i < 4; i++ {
		time.Sleep(10 * time.Second)
		logger.Info("Текущая цепочка блоков:")
		for _, b := range node.Chain {
			logger.Info("Блок:", zap.Any("block", b))
		}
	}
	node.Stop()
	logger.Info("FullNode остановлен (тест завершён)")
}

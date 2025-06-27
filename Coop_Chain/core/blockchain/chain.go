package blockchain

import (
	"fmt"
	"mcp-chain/core/mempool"
	"mcp-chain/core/state"
	"mcp-chain/core/transaction"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
	"mcp-chain/types"
	"os"
	"time"
)

// Chain — структура для работы с блоками теперь использует state.Blocks
// Использует БД и текущее состояние
// (упрощённая версия для MVP)
type Chain struct {
}

// NewChain создаёт новую цепочку
func NewChain() *Chain {
	return &Chain{}
}

// AppendBlock добавляет новый блок в цепочку
func (c *Chain) AppendBlock(b types.Block) error {
	logging.Logger.Info("AppendBlock: ДО добавления блока",
		"Accounts", len(state.GlobalState.Accounts),
		"AccountsList", keys(state.GlobalState.Accounts),
		"Validators", len(state.GlobalState.Validators),
		"Blocks", len(state.GlobalState.Blocks),
	)
	err := state.ApplyStateChange(func(st *state.State) {
		st.Blocks = append(st.Blocks, b)
		logging.Logger.Info("AppendBlock: ПОСЛЕ добавления блока",
			"Accounts", len(st.Accounts),
			"AccountsList", keys(st.Accounts),
			"Validators", len(st.Validators),
			"Blocks", len(st.Blocks),
		)
	})
	if err != nil {
		logging.Logger.Error("Ошибка сохранения state после добавления блока", "err", err)
	}
	return err
}

// keys возвращает список ключей map[string]T
func keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// VerifyChain проверяет целостность цепочки блоков
func (c *Chain) VerifyChain() bool {
	// TODO: реализовать проверку связности и валидности блоков
	return true
}

// Head возвращает последний блок
func (c *Chain) Head() *types.Block {
	block := &state.GlobalState.Blocks[len(state.GlobalState.Blocks)-1]
	logging.Logger.Debug("GetLastBlock (Head)", "height", block.Header.Height, "block", block)
	return block
}

// Blocks возвращает копию всех блоков цепи
func (c *Chain) Blocks() []types.Block {
	logging.Logger.Debug("Blocks called", "blocks", len(state.GlobalState.Blocks))
	return append([]types.Block(nil), state.GlobalState.Blocks...)
}

// InitChain создает новую цепочку блоков, состояние и Genesis-блок по заданному конфигу сети.
// Возвращает цепочку, состояние и конфиг.
func InitChain() (*Chain, config.NetworkConfig) {
	logging.Logger.Info("InitChain: загрузка state из файла")
	cfg := config.GetDefaultNetworkConfig()
	st, err := state.LoadStateFromFile(config.StateFile)
	if err != nil {
		if os.IsNotExist(err) {
			logging.Logger.Warn("InitChain: state.json не найден, создаю новый state")
			state.GlobalState = state.NewState()
			_ = state.ApplyStateChange(func(st *state.State) {})
		} else {
			logging.Logger.Error("InitChain: ошибка загрузки state", "err", err)
			panic(fmt.Sprintf("Ошибка загрузки state: %v", err))
		}
	} else {
		logging.Logger.Info("InitChain: state успешно загружен из файла")
		state.GlobalState = st
	}
	return NewChain(), cfg
}

// MineBlock создает новый блок с транзакциями из mempool
func (c *Chain) MineBlock(cfg *config.NetworkConfig, mp *mempool.TxMempool) types.Block {
	prev := c.Head()

	var txHashes []types.Hash
	var txObjs []types.Transaction
	for _, tx := range mp.GetTxs() {
		txHashes = append(txHashes, tx.ID)
		txObjs = append(txObjs, *tx)
	}

	newBlock := types.Block{
		Header: types.Header{
			Height:    prev.Header.Height + 1,
			PrevHash:  prev.Hash,
			Timestamp: types.Timestamp(time.Now().Unix()),
			Proposer:  prev.Header.Proposer, // MVP: тот же валидатор
		},
		Body: types.Body{
			Txs:          txHashes,
			Transactions: txObjs,
		},
	}
	// Хэш блока (MVP: без подписи)
	blockForHash := Block{
		Header: Header{
			Height:    newBlock.Header.Height,
			PrevHash:  newBlock.Header.PrevHash,
			Timestamp: int64(newBlock.Header.Timestamp),
			Proposer:  newBlock.Header.Proposer,
		},
		Txs: []transaction.Tx{}, // Не используется для хэша
	}
	newBlock.Hash = blockForHash.CalcHash(cfg)

	// Возвращаем карту транзакций вместе с блоком (MVP: через глобальную переменную или отдельный storage)
	// Для простоты: сохраняем карту в Chain (или возвращаем из MineBlock)
	state.GlobalState.Cache.RecentBlocks = append(state.GlobalState.Cache.RecentBlocks, newBlock.Hash)
	// Очищаем mempool после применения транзакций
	// mp.Clear() вызывается после применения транзакций в RunNode

	return newBlock
}

// ChainStats — структура для статистики цепи
type ChainStats struct {
	BlockCount      int
	WalletCount     int
	ValidatorCount  int
	TotalBalance    float64
	LastBlockHeight uint64
	LastBlockHash   types.Hash
	LastBlockTime   int64
	LastValidator   string
	Uptime          time.Duration
}

// GetChainStats возвращает статистику по цепи и состоянию
func GetChainStats(chain *Chain, startTime time.Time) ChainStats {
	blocks := chain.Blocks()
	blockCount := len(blocks)
	walletCount := len(state.GlobalState.Accounts)
	totalBalance := 0.0
	for _, acc := range state.GlobalState.Accounts {
		totalBalance += float64(acc.Balance)
	}
	validatorCount := len(state.GlobalState.Validators)
	var lastHeight uint64
	var lastHash types.Hash
	var lastTime int64
	var lastValidator string
	if blockCount > 0 {
		lastBlock := blocks[blockCount-1]
		lastHeight = lastBlock.Header.Height
		lastHash = lastBlock.Hash
		lastTime = int64(lastBlock.Header.Timestamp)
		lastValidator = string(lastBlock.Header.Proposer)
	}
	return ChainStats{
		BlockCount:      blockCount,
		WalletCount:     walletCount,
		ValidatorCount:  validatorCount,
		TotalBalance:    totalBalance,
		LastBlockHeight: lastHeight,
		LastBlockHash:   lastHash,
		LastBlockTime:   lastTime,
		LastValidator:   lastValidator,
		Uptime:          time.Since(startTime),
	}
}

// BlockInfo — структура для информации о блоке
type BlockInfo struct {
	Height    uint64
	Hash      types.Hash
	Timestamp int64
	Proposer  types.Address
}

// BlocksInfo возвращает информацию о всех блоках
func BlocksInfo() []BlockInfo {
	blocks := state.GlobalState.Blocks
	var res []BlockInfo
	for _, b := range blocks {
		res = append(res, BlockInfo{
			Height:    b.Header.Height,
			Hash:      b.Hash,
			Timestamp: int64(b.Header.Timestamp),
			Proposer:  b.Header.Proposer,
		})
	}
	return res
}

// RunNode запускает узел и майнит блоки бесконечно (реальная сеть)
func RunNode(print func(string, ...any)) {
	logging.Logger.Info("RunNode start")
	chain, cfg := InitChain()
	mp := mempool.GlobalTxMempool // Используем глобальный mempool!

	if len(state.GlobalState.Blocks) == 0 {
		logging.Logger.Info("Creating Genesis block")
		print("[NODE] Нет ни одного блока, создаю Genesis...")
		genesisBlock := CreateGenesis(&cfg)
		typesBlock := blockchainBlockToTypesBlock(genesisBlock, &cfg)
		state.GlobalState.Blocks = append(state.GlobalState.Blocks, typesBlock)
		logging.Logger.Info("AppendBlock (Genesis)", "height", typesBlock.Header.Height, "hash", typesBlock.Hash)
		logging.Logger.Info("Перед DumpStateToFile (Genesis)",
			"Accounts", len(state.GlobalState.Accounts),
			"Validators", len(state.GlobalState.Validators),
			"CoopsRegistry", len(state.GlobalState.CoopsRegistry),
			"Governance", len(state.GlobalState.Governance),
			"SoulBound", len(state.GlobalState.SoulBound),
			"ContractsStorage", len(state.GlobalState.ContractsStorage),
			"FeeTreasury", state.GlobalState.FeeTreasury.Balance,
			"Blocks", len(state.GlobalState.Blocks),
		)
		_ = state.ApplyStateChange(func(st *state.State) {})
		logging.Logger.Info("Genesis block created", "height", typesBlock.Header.Height, "hash", typesBlock.Hash)
	}
	print("[NODE] Genesis-блок добавлен. Высота: %d", chain.Head().Header.Height)

	for {
		// --- HOT-RELOAD state.json ---
		// УДАЛЕНО: больше не подгружаем state.json в цикле
		// --- END HOT-RELOAD ---

		time.Sleep(time.Duration(cfg.AutoMineInterval) * time.Second)
		newBlock := chain.MineBlock(&cfg, mp)
		chain.AppendBlock(newBlock)
		// Применяем транзакции блока к state
		for _, tx := range newBlock.Body.Transactions {
			state.GlobalState.ApplyTx(&tx)
		}
		mp.Clear()
		print("[NODE] Блок #%d добавлен. Хэш: %s", newBlock.Header.Height, newBlock.Hash)
		logging.Logger.Info("Перед DumpStateToFile (RunNode)",
			"Accounts", len(state.GlobalState.Accounts),
			"Validators", len(state.GlobalState.Validators),
			"CoopsRegistry", len(state.GlobalState.CoopsRegistry),
			"Governance", len(state.GlobalState.Governance),
			"SoulBound", len(state.GlobalState.SoulBound),
			"ContractsStorage", len(state.GlobalState.ContractsStorage),
			"FeeTreasury", state.GlobalState.FeeTreasury.Balance,
			"Blocks", len(state.GlobalState.Blocks),
		)
		_ = state.ApplyStateChange(func(st *state.State) {})
		// DEBUG: Выводим количество блоков и их хэши из памяти
		fmt.Printf("[DEBUG] Всего блоков в памяти: %d\n", len(chain.Blocks()))
		blocks := chain.Blocks()
		for i, b := range blocks {
			fmt.Printf("[DEBUG] Block #%d | Hash: %s\n", b.Header.Height, b.Hash)
			if i > 20 {
				fmt.Printf("[DEBUG] ... (ещё %d блоков)\n", len(blocks)-21)
				break
			}
		}
	}
}

// GetChainStatsFromFile читает state из файла и возвращает статистику цепи
func GetChainStatsFromFile(stateFile string, startTime time.Time) (ChainStats, error) {
	st, err := state.LoadStateFromFile(stateFile)
	if err != nil {
		return ChainStats{}, err
	}
	state.GlobalState = st
	return GetChainStats(NewChain(), startTime), nil
}

// BlocksInfoFromFile читает state из файла и возвращает информацию о блоках
func BlocksInfoFromFile(stateFile string) ([]BlockInfo, error) {
	st, err := state.LoadStateFromFile(stateFile)
	if err != nil {
		return nil, err
	}
	state.GlobalState = st
	return BlocksInfo(), nil
}

// Вспомогательная функция для конвертации core/blockchain.Block в types.Block
func blockchainBlockToTypesBlock(b *Block, cfg *config.NetworkConfig) types.Block {
	return types.Block{
		Header: types.Header{
			Height:    b.Header.Height,
			PrevHash:  b.Header.PrevHash,
			Timestamp: types.Timestamp(b.Header.Timestamp),
			Proposer:  b.Header.Proposer,
		},
		Body: types.Body{
			Txs: nil, // или преобразовать транзакции, если нужно
		},
		Hash: b.CalcHash(cfg),
	}
}

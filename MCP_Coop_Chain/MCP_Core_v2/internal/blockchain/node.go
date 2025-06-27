package blockchain

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"mcp-coop-chain/internal"
	"mcp-coop-chain/internal/contracts/proto"
	"mcp-coop-chain/internal/logging"
	"mcp-coop-chain/internal/storage"
	"mcp-coop-chain/internal/types"
	storagetypes "mcp-coop-chain/internal/types"
	"mcp-coop-chain/internal/wallet"
)

// FullNode — основная структура узла блокчейна MCP Coop Chain
// Хранит состояние цепочки, mempool, storage, логгер и управляет жизненным циклом узла
// Все поля и методы снабжены подробными комментариями

type FullNode struct {
	Chain   []*types.Block          // Цепочка блоков
	Mempool []*types.Transaction    // Массив неподтверждённых транзакций
	Storage types.Storage           // Интерфейс хранилища состояния
	Config  *types.BlockchainConfig // Конфиг сети
	Logger  logging.Logger          // Централизованный логгер
	Running bool                    // Флаг работы узла
	mu      sync.Mutex              // Мьютекс для потокобезопасности
}

// NewFullNode создает новый FullNode с заданным конфигом, storage и логгером
func NewFullNode(cfg *types.BlockchainConfig, st types.Storage, logger logging.Logger) *FullNode {
	n := &FullNode{
		Chain:   make([]*types.Block, 0),
		Mempool: make([]*types.Transaction, 0),
		Storage: st,
		Config:  cfg,
		Logger:  logger,
		Running: false,
	}
	// --- NEW: загрузка состояния из снапшота ---
	if ms, ok := st.(*storage.MemoryStorage); ok {
		snapshotPath := cfg.StoragePath
		if snapshotPath == "" {
			snapshotPath = "data/full_state_snapshot.json"
		}
		snapshot, err := storage.LoadFullSnapshot(snapshotPath)
		if err == nil && snapshot != nil {
			_ = storage.ApplyRestoredState(snapshot, ms, &n.Chain, &n.Mempool)
			logger.Info("[BOOT] Состояние восстановлено из снапшота")
		}
	}
	return n
}

// AddBlock добавляет новый блок в цепочку с проверкой целостности и логированием
func (n *FullNode) AddBlock(block *types.Block, logger logging.Logger) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if len(n.Chain) > 0 {
		last := n.Chain[len(n.Chain)-1]
		if block.Header.PreviousHash != last.Header.BlockID {
			err := fmt.Errorf("prevHash не совпадает с последним блоком цепочки")
			logger.Error("Ошибка добавления блока", err)
			return err
		}
	}
	n.Chain = append(n.Chain, block)
	logger.Info("Блок добавлен в цепочку")
	n.flushSnapshot()
	return nil
}

// InitGenesisBlock инициализирует генезис-блок (первый блок в цепочке)
func (n *FullNode) InitGenesisBlock(logger logging.Logger) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if len(n.Chain) > 0 {
		return fmt.Errorf("Генезис-блок уже существует")
	}
	// Регистрируем все встроенные proto-контракты
	proto.RegisterAllProtoContracts()
	genesis := NewBlock(
		0,
		"",
		"genesis",
		"system",
		[]types.Transaction{},
		[]types.ContractCall{},
	)
	n.Chain = append(n.Chain, &genesis)
	logger.Info("Генезис-блок инициализирован")
	n.flushSnapshot()
	return nil
}

// Start запускает узел: периодически создает новые блоки из mempool
func (n *FullNode) Start() {
	n.mu.Lock()
	n.Running = true
	n.mu.Unlock()
	// Регистрируем все встроенные proto-контракты
	proto.RegisterAllProtoContracts()
	n.Logger.Info("[LIVE] Форсированный старт: создаём первый блок сразу после запуска")
	n.CreateBlockFromMempool() // форсируем создание первого блока
	go n.blockProducer()
}

// Stop останавливает работу узла
func (n *FullNode) Stop() {
	n.mu.Lock()
	n.Running = false
	n.mu.Unlock()
}

// blockProducer — отдельная горутина, которая создает блоки по таймеру
func (n *FullNode) blockProducer() {
	ticker := time.NewTicker(time.Duration(internal.CurrentConfig.Chain.BlockTimeSeconds) * time.Second)
	defer ticker.Stop()
	for {
		n.mu.Lock()
		if !n.Running {
			n.mu.Unlock()
			return
		}
		n.mu.Unlock()
		<-ticker.C
		n.CreateBlockFromMempool()
	}
}

// AddBlockWithSnapshot добавляет блок в цепочку и snapshot-файл
func (n *FullNode) AddBlockWithSnapshot(block *types.Block) error {
	if err := n.AddBlock(block, n.Logger); err != nil {
		return err
	}
	n.Logger.Info("[LIVE] Добавлен блок и обновлён снапшот: " + block.Header.BlockID)
	n.flushSnapshot()
	return nil
}

// getBlocks возвращает []Block из []*Block
func (n *FullNode) getBlocks() []types.Block {
	blocks := make([]types.Block, len(n.Chain))
	for i, b := range n.Chain {
		blocks[i] = *b
	}
	return blocks
}

// getMempool возвращает []Transaction из []*Transaction
func (n *FullNode) getMempool() []types.Transaction {
	txs := make([]types.Transaction, len(n.Mempool))
	for i, t := range n.Mempool {
		txs[i] = *t
	}
	return txs
}

// CreateBlockFromMempool создает новый блок из транзакций mempool и contractCalls
func (n *FullNode) CreateBlockFromMempool() {
	n.mu.Lock()
	allowEmpty := internal.CurrentConfig.Chain.AllowEmptyBlocks
	maxMempool := internal.CurrentConfig.Chain.MaxTxMempool
	txPerBlock := internal.CurrentConfig.Chain.TxPerBlock
	if len(n.Mempool) == 0 && !allowEmpty {
		n.Logger.Info("Нет транзакций в mempool — блок не создаётся")
		n.mu.Unlock()
		return
	}
	if maxMempool > 0 && len(n.Mempool) > maxMempool {
		n.Logger.Info("Mempool переполнен — досрочно создаём блок")
		// Можно реализовать досрочное создание блока здесь
	}
	prevBlock := n.Chain[len(n.Chain)-1]
	var contractCalls []types.ContractCall
	var txs []types.Transaction
	// Получаем state для списания комиссии
	state := map[string]uint64{}
	if st, ok := n.Storage.(interface{ GetAllBalances() map[string]uint64 }); ok {
		state = st.GetAllBalances()
	}
	count := 0
	for _, txPtr := range n.Mempool {
		if txPerBlock > 0 && count >= txPerBlock {
			break // не добавляем больше, чем разрешено
		}
		tx := *txPtr
		// Если это ContractCall (адрес proto:...)
		if len(tx.SmartContractHashes) > 0 && len(tx.SmartContractHashes[0]) > 6 && tx.SmartContractHashes[0][:6] == "proto:" {
			method := tx.SmartContractHashes[0][6:]
			args := []byte("")
			if len(tx.Payload) > 0 {
				args = tx.Payload
			}
			wallets, ok1 := n.Storage.(storagetypes.WalletStorage)
			contracts, ok2 := n.Storage.(storagetypes.ContractStorage)
			if ok1 && ok2 {
				// Списываем комиссию за contract call
				fee := internal.CurrentConfig.Fees.CallProtoAPIMethodFee
				if !burnFee(state, tx.Origin, fee) {
					n.Logger.Warn("Недостаточно средств для комиссии contract call")
					continue // недостаточно средств — contract call не добавлять
				}
				result, err := proto.CallProtoContract(method, wallets, contracts, args)
				cc := types.ContractCall{
					Contract: method,
					Method:   method,
					Args:     []string{string(args)},
					Caller:   tx.Origin,
					Result:   string(result),
				}
				if err != nil {
					cc.Result += ";error=" + err.Error()
				}
				contractCalls = append(contractCalls, cc)
				count++
			} else {
				n.Logger.Error("Storage не реализует WalletStorage/ContractStorage", nil)
			}
		} else {
			txs = append(txs, tx)
			count++
		}
	}
	block := NewBlock(
		prevBlock.Header.Height+1,
		prevBlock.Header.BlockID,
		"",                 // обычный блок
		"proposer-address", // TODO: получить адрес инициатора
		txs,
		contractCalls,
	)
	n.Logger.Info(fmt.Sprintf("[LIVE] Создаём новый блок: height=%d blockId=%s", block.Header.Height, block.Header.BlockID))
	// Очищаем только те транзакции, которые попали в блок
	if txPerBlock > 0 && len(n.Mempool) > count {
		n.Mempool = n.Mempool[count:]
	} else {
		n.Mempool = []*types.Transaction{}
	}
	n.mu.Unlock()

	n.AddBlockWithSnapshot(&block)

	n.Logger.Info("Создан новый блок")
	// После создания блока — обновляем снапшот
	n.flushSnapshot()
}

// SaveState сохраняет текущее состояние цепочки в storage
func (n *FullNode) SaveState() error {
	n.mu.Lock()
	defer n.mu.Unlock()
	data, err := json.Marshal(n.Chain)
	if err != nil {
		return err
	}
	return n.Storage.SaveState(data)
}

// RestoreState восстанавливает цепочку блоков из storage
func (n *FullNode) RestoreState() error {
	n.mu.Lock()
	defer n.mu.Unlock()
	data, err := n.Storage.LoadState()
	if err != nil {
		return err
	}
	var chain []*types.Block
	if err := json.Unmarshal(data, &chain); err != nil {
		return err
	}
	n.Chain = chain
	return nil
}

// AddTransaction добавляет транзакцию в mempool с учётом лимита и полной валидации
func (n *FullNode) AddTransaction(tx *types.Transaction) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	maxMempool := internal.CurrentConfig.Chain.MaxTxMempool
	if maxMempool > 0 && len(n.Mempool) >= maxMempool {
		n.Logger.Warn("Mempool is full, transaction rejected")
		return fmt.Errorf("mempool is full: limit %d reached", maxMempool)
	}
	// Получаем state (балансы) из Storage, если возможно
	state := map[string]uint64{}
	if st, ok := n.Storage.(interface{ GetAllBalances() map[string]uint64 }); ok {
		state = st.GetAllBalances()
	}
	if err := ValidateTransactionForMempool(tx, state, n.Mempool); err != nil {
		n.Logger.Warn(fmt.Sprintf("Transaction rejected: %v", err))
		return err
	}
	n.Mempool = append(n.Mempool, tx)
	n.Logger.Info(fmt.Sprintf("Transaction added to mempool: %s", tx.TxID))
	// Досрочное создание блока при переполнении mempool
	if maxMempool > 0 && len(n.Mempool) >= maxMempool {
		go n.CreateBlockFromMempool()
	}
	return nil
}

// BuildFullSnapshot собирает всё состояние из FullNode и MemoryStorage
func (n *FullNode) BuildFullSnapshot() *types.FullChainSnapshot {
	var wallets []types.Wallet
	var contracts []types.ContractCall
	var orgs []types.Organization
	var state types.ChainState
	if ms, ok := n.Storage.(*storage.MemoryStorage); ok {
		ms.Mu.RLock()
		for _, mw := range ms.Wallets {
			wallets = append(wallets, mw.Wallet)
		}
		// TODO: если контракты сериализуются иначе — доработать
		ms.Mu.RUnlock()
	}
	// TODO: добавить contracts/orgs/state если нужно
	return &types.FullChainSnapshot{
		Blocks:        n.getBlocks(),
		Mempool:       n.getMempool(),
		Organizations: orgs,
		Wallets:       wallets,
		State:         state,
		Contracts:     contracts,
		Timestamp:     time.Now().UTC(),
	}
}

// flushSnapshot сохраняет актуальное состояние FullNode в снапшот-файл
func (n *FullNode) flushSnapshot() {
	snapshotPath := n.Config.StoragePath
	if snapshotPath == "" {
		snapshotPath = "data/chain_snapshot.json"
	}
	snapshot := n.BuildFullSnapshot()
	_ = storage.SaveFullSnapshot(snapshot, snapshotPath)
}

// AddWallet — централизованное создание кошелька через FullNode
func (n *FullNode) AddWallet() (*types.Wallet, error) {
	priv, err := wallet.CreateWallet()
	if err != nil {
		return nil, err
	}
	pub := wallet.ToPublicWallet(priv)
	if ws, ok := n.Storage.(types.WalletStorage); ok {
		err = ws.AddWallet(pub)
		if err != nil {
			return nil, err
		}
		n.flushSnapshot()
	}
	return &pub, nil
}

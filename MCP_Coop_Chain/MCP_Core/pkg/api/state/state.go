package state

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mcpcoop/chain/pkg/types"
)

const stateDir = "State"
const stateFileName = "State.json"

var stateLock sync.Mutex

const (
	txMempoolMax            = 10000
	txMempoolBlockThreshold = 100
	scMempoolMax            = 1000
)

// StateAPI - централизованный API для управления состоянием
// Используется только через NewStateAPI()
type StateAPI struct{}

func NewStateAPI() *StateAPI {
	return &StateAPI{}
}

func getStateFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return filepath.Join(stateDir, stateFileName)
	}
	return filepath.Join(cwd, stateDir, stateFileName)
}

func ensureStateDir() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(cwd, stateDir), 0755)
}

func defaultState() *types.NodeState {
	genesis := types.Block{
		BlockHeader: types.BlockHeader{
			Index:     0,
			Timestamp: time.Now(),
			PrevHash:  "",
			Hash:      "genesis",
		},
		Transactions: nil,
	}
	return &types.NodeState{
		Status:     "stopped",
		Chain:      []types.Block{genesis},
		TxMempool:  types.TxMempool{Pending: []types.Transaction{}},
		ScMempool:  types.ScMempool{Pending: []types.Transaction{}},
		Wallets:    []types.Wallet{},
		Balances:   map[string]float64{},
		StartTime:  "",
		MerkleRoot: "",
	}
}

func saveStateNoLock(s *types.NodeState) {
	s.MerkleRoot = calculateMerkleRoot(s.Chain)
	b, _ := json.MarshalIndent(s, "", "  ")
	_ = ioutil.WriteFile(getStateFilePath(), b, 0644)
}

func saveState(s *types.NodeState) {
	stateLock.Lock()
	defer stateLock.Unlock()
	saveStateNoLock(s)
}

func loadState() *types.NodeState {
	stateLock.Lock()
	defer stateLock.Unlock()
	ensureStateDir()
	b, err := ioutil.ReadFile(getStateFilePath())
	if err != nil {
		s := defaultState()
		saveStateNoLock(s)
		return s
	}
	var s types.NodeState
	if err := json.Unmarshal(b, &s); err != nil {
		s := defaultState()
		saveStateNoLock(s)
		return s
	}
	if s.Balances == nil {
		s.Balances = map[string]float64{}
	}
	if s.Wallets == nil {
		s.Wallets = []types.Wallet{}
	}
	if s.Chain == nil || len(s.Chain) == 0 {
		s.Chain = defaultState().Chain
	}
	if s.TxMempool.Pending == nil {
		s.TxMempool.Pending = []types.Transaction{}
	}
	if s.ScMempool.Pending == nil {
		s.ScMempool.Pending = []types.Transaction{}
	}
	s.MerkleRoot = calculateMerkleRoot(s.Chain)
	return &s
}

func calculateMerkleRoot(chain []types.Block) string {
	hashes := ""
	for _, b := range chain {
		hashes += string(b.Hash)
	}
	if hashes == "" {
		return ""
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashes)))
}

// --- API методы ---

func (api *StateAPI) StartNode() error {
	s := loadState()
	if s.Status == "running" {
		return fmt.Errorf("Node is already running")
	}
	s.Status = "running"
	s.StartTime = time.Now().Format(time.RFC3339)
	saveState(s)
	fmt.Printf("[INFO] Node started | height: %d | wallets: %d\n", len(s.Chain)-1, len(s.Wallets))
	// Можно добавить запуск таймера блоков, если нужно
	return nil
}

func (api *StateAPI) StopNode() error {
	s := loadState()
	if s.Status != "running" {
		return fmt.Errorf("Node is already stopped")
	}
	s.Status = "stopped"
	saveState(s)
	fmt.Printf("[INFO] Node stopped\n")
	return nil
}

func (api *StateAPI) RestartNode() error {
	s := loadState()
	if s.Status != "running" {
		return fmt.Errorf("Node is not running; cannot restart. Use 'start' instead.")
	}
	if err := api.StopNode(); err != nil {
		return err
	}
	return api.StartNode()
}

func (api *StateAPI) NewWallet(args []string) error {
	s := loadState()
	if s.Status != "running" {
		return fmt.Errorf("Node is not running. Start the node first.")
	}
	if len(args) == 0 || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("Seed phrase required. Use --seed 'word1 word2 ... word12'")
	}
	words := strings.Fields(args[0])
	if len(words) != 12 {
		return fmt.Errorf("Seed phrase must be exactly 12 words.")
	}
	seed := sha256.Sum256([]byte(strings.Join(words, " ")))
	pub, _, err := ed25519.GenerateKey(strings.NewReader(string(seed[:])))
	if err != nil {
		return fmt.Errorf("Failed to generate key: %s", err.Error())
	}
	address := hex.EncodeToString(pub[:8])
	for _, w := range s.Wallets {
		if string(w.Address) == address {
			fmt.Printf("Wallet already exists!\nPublic address: %s\n(Your wallet is ready to use. Keep your seed phrase safe!)\n", address)
			saveState(s)
			return nil
		}
	}
	s.Wallets = append(s.Wallets, types.Wallet{
		Address: types.Address(address),
		PubKey:  pub,
		Balance: 100.0,
	})
	if s.Balances == nil {
		s.Balances = map[string]float64{}
	}
	s.Balances[address] = 100.0 // Give new wallet a mock balance
	saveState(s)
	fmt.Printf("Wallet created!\nPublic address: %s\n(Your wallet is ready to use. Keep your seed phrase safe!)\n", address)
	return nil
}

func (api *StateAPI) Send(args []string) error {
	s := loadState()
	if s.Status != "running" {
		return fmt.Errorf("Node is not running. Start the node first.")
	}
	if len(args) < 3 {
		return fmt.Errorf("Missing required flags: --from, --to, --amount")
	}
	from := args[0]
	to := args[1]
	amountStr := args[2]
	if from == "" || to == "" || amountStr == "" {
		return fmt.Errorf("All of --from, --to, --amount are required.")
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		return fmt.Errorf("Amount must be a positive number.")
	}
	fromBal := s.Balances[from]
	if fromBal < amount {
		return fmt.Errorf("Insufficient balance.")
	}
	toBal := s.Balances[to]
	fromBal -= amount
	toBal += amount
	s.Balances[from] = fromBal
	s.Balances[to] = toBal

	tx := types.Transaction{
		From:      types.Address(from),
		To:        types.Address(to),
		Amount:    amount,
		Timestamp: time.Now(),
	}
	s.TxMempool.Pending = append(s.TxMempool.Pending, tx)
	saveState(s)
	fmt.Printf("Sent %.2f tokens from %s to %s\nNew balance:\n  %s: %.2f\n  %s: %.2f\n",
		amount, from, to, from, fromBal, to, toBal)
	return nil
}

func (api *StateAPI) Status() (string, error) {
	s := loadState()
	now := time.Now()
	uptime := ""
	if s.StartTime != "" {
		if t, err := time.Parse(time.RFC3339, s.StartTime); err == nil {
			uptime = formatUptime(now.Sub(t))
		}
	}
	timeToNextBlock := "N/A"
	if len(s.Chain) > 0 {
		lastBlock := s.Chain[len(s.Chain)-1]
		elapsed := int(now.Sub(lastBlock.Timestamp).Seconds())
		if elapsed < 0 {
			elapsed = 0
		}
		rem := 60 - (elapsed % 60)
		if rem == 60 {
			rem = 0
		}
		timeToNextBlock = fmt.Sprintf("%ds", rem)
	}
	out := fmt.Sprintf("Node status: %s\nBlock height: %d\nTxMempool txs: %d\nScMempool txs: %d\nWallets: %d\nTime: %s\nUptime: %s\nNext block in: %s\n",
		s.Status, len(s.Chain)-1, len(s.TxMempool.Pending), len(s.ScMempool.Pending), len(s.Wallets), now.Format(time.RFC3339), uptime, timeToNextBlock)
	return out, nil
}

func formatUptime(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}

// Добавление стандартной транзакции в txMempool
func (api *StateAPI) AddTxToMempool(tx types.Transaction) error {
	s := loadState()
	if len(s.TxMempool.Pending) >= txMempoolMax {
		msg := "Mempool full – try again later."
		fmt.Println("[MEMPOOL] " + msg)
		return fmt.Errorf(msg)
	}
	s.TxMempool.Pending = append(s.TxMempool.Pending, tx)
	fmt.Printf("[MEMPOOL] Transaction added to txMempool. Total: %d\n", len(s.TxMempool.Pending))
	saveState(s)
	// Если достигнут порог — немедленно формируем блок
	if len(s.TxMempool.Pending) >= txMempoolBlockThreshold {
		api.CreateBlockIfNeeded(s)
	}
	return nil
}

// Добавление транзакции смарт-контракта в scMempool
func (api *StateAPI) AddScTransaction(tx types.Transaction) error {
	s := loadState()
	if len(s.ScMempool.Pending) >= scMempoolMax {
		msg := "Smart-contract mempool full – try again later."
		fmt.Println("[MEMPOOL] " + msg)
		return fmt.Errorf(msg)
	}
	if err := validateSmartContractTx(tx, s); err != nil {
		fmt.Printf("[MEMPOOL] SC tx rejected: %s\n", err)
		return err
	}
	s.ScMempool.Pending = append(s.ScMempool.Pending, tx)
	fmt.Printf("[MEMPOOL] SC transaction added. Total: %d\n", len(s.ScMempool.Pending))
	saveState(s)
	return nil
}

// Очистка mempool после создания блока
func (api *StateAPI) ClearMempools(s *types.NodeState, includedTxs []types.Transaction, includedScTxs []types.Transaction) {
	s.TxMempool.Pending = filterOut(s.TxMempool.Pending, includedTxs)
	s.ScMempool.Pending = filterOut(s.ScMempool.Pending, includedScTxs)
	saveState(s)
}

func filterOut(pool []types.Transaction, included []types.Transaction) []types.Transaction {
	m := make(map[string]struct{})
	for _, tx := range included {
		m[tx.Hash()] = struct{}{}
	}
	var out []types.Transaction
	for _, tx := range pool {
		if _, ok := m[tx.Hash()]; !ok {
			out = append(out, tx)
		}
	}
	return out
}

// Проверка смарт-контрактной транзакции (заглушка)
func validateSmartContractTx(tx types.Transaction, s *types.NodeState) error {
	// TODO: Реализовать реальную логику проверки SC-транзакций
	if tx.To == "" {
		return fmt.Errorf("SC tx: empty destination")
	}
	return nil
}

// Создание блока по правилам (txMempool >= 100 или scMempool > 0)
func (api *StateAPI) CreateBlockIfNeeded(s *types.NodeState) {
	if len(s.TxMempool.Pending) >= txMempoolBlockThreshold || len(s.ScMempool.Pending) > 0 {
		allTxs := append([]types.Transaction{}, s.TxMempool.Pending...)
		allTxs = append(allTxs, s.ScMempool.Pending...)
		prev := s.Chain[len(s.Chain)-1]
		block := types.Block{
			BlockHeader: types.BlockHeader{
				Index:     prev.BlockHeader.Index + 1,
				Timestamp: time.Now(),
				PrevHash:  prev.BlockHeader.Hash,
				Hash:      types.Hash(fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d%s", prev.BlockHeader.Index+1, time.Now().String()))))),
			},
			Transactions: allTxs,
		}
		s.Chain = append(s.Chain, block)
		api.ClearMempools(s, s.TxMempool.Pending, s.ScMempool.Pending)
		fmt.Printf("[BLOCK] Block created! Height: %d, txs: %d, scTxs: %d\n", len(s.Chain)-1, len(s.TxMempool.Pending), len(s.ScMempool.Pending))
		saveState(s)
	}
}

// Package cli provides the command-line interface for the blockchain node
package cli

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/mcpcoop/chain/internal/logger"
	"github.com/mcpcoop/chain/pkg/crypto"
	"github.com/mcpcoop/chain/pkg/types"

	"errors"

	"github.com/mcpcoop/chain/pkg/api/cli/commands"
	"github.com/mcpcoop/chain/pkg/chain"
	"github.com/mcpcoop/chain/pkg/wallet"
)

const stateFile = "State/State.json"

// nodeState - full persistent state
// Includes: chain, mempool, wallets, balances
type nodeState struct {
	Chain     []types.Block       `json:"chain"`
	Mempool   []types.Transaction `json:"mempool"`
	Wallets   []*wallet.Wallet    `json:"wallets"`
	Balances  map[string]float64  `json:"balances"`
	StartTime string              `json:"start_time"`
}

var (
	stateLock sync.Mutex
)

func getStateFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return stateFile
	}
	return filepath.Join(cwd, stateFile)
}

func defaultState() *nodeState {
	genesis := types.Block{
		BlockHeader: types.BlockHeader{
			Index:     0,
			Timestamp: time.Now(),
			PrevHash:  "",
			Hash:      "genesis",
		},
		Transactions: nil,
	}
	return &nodeState{
		Chain:     []types.Block{genesis},
		Mempool:   []types.Transaction{},
		Wallets:   []*wallet.Wallet{},
		Balances:  map[string]float64{},
		StartTime: "",
	}
}

func loadState() *nodeState {
	stateLock.Lock()
	defer stateLock.Unlock()
	b, err := ioutil.ReadFile(getStateFilePath())
	if err != nil {
		return defaultState()
	}
	var s nodeState
	if err := json.Unmarshal(b, &s); err != nil {
		return defaultState()
	}
	if s.Balances == nil {
		s.Balances = map[string]float64{}
	}
	if s.Wallets == nil {
		s.Wallets = []*wallet.Wallet{}
	}
	if s.Chain == nil || len(s.Chain) == 0 {
		s.Chain = defaultState().Chain
	}
	if s.Mempool == nil {
		s.Mempool = []types.Transaction{}
	}
	return &s
}

func saveState(s *nodeState) {
	stateLock.Lock()
	defer stateLock.Unlock()
	b, _ := json.MarshalIndent(s, "", "  ")
	_ = ioutil.WriteFile(getStateFilePath(), b, 0644)
}

// --- Block creation timer ---
var (
	ticker     *time.Ticker
	tickerStop chan struct{}
)

func startBlockTimer() {
	if ticker != nil {
		return
	}
	ticker = time.NewTicker(60 * time.Second)
	tickerStop = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				s := loadState()
				if len(s.Mempool) > 0 {
					addBlock(s)
					slog("Block created by timer", s)
					saveState(s)
					fmt.Printf("[TIMER] Block created! New height: %d\n", len(s.Chain)-1)
				} else {
					fmt.Println("[TIMER] No transactions in mempool, block not created.")
				}
			case <-tickerStop:
				ticker.Stop()
				ticker = nil
				return
			}
		}
	}()
}

func stopBlockTimer() {
	if ticker != nil && tickerStop != nil {
		close(tickerStop)
	}
}

func addBlock(s *nodeState) {
	prev := s.Chain[len(s.Chain)-1]
	block := types.Block{
		BlockHeader: types.BlockHeader{
			Index:     prev.BlockHeader.Index + 1,
			Timestamp: time.Now(),
			PrevHash:  prev.BlockHeader.Hash,
			Hash:      types.Hash(fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d%s", prev.BlockHeader.Index+1, time.Now().String()))))),
		},
		Transactions: s.Mempool,
	}
	s.Chain = append(s.Chain, block)
	s.Mempool = []types.Transaction{}
}

func slog(msg string, s *nodeState) {
	logger.Logger.Info(msg,
		"height", len(s.Chain)-1,
		"mempool", len(s.Mempool),
		"wallets", len(s.Wallets),
	)
}

// Execute parses and executes CLI commands
func Execute() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	// Load chain state from disk or create new
	c := chain.NewChain()
	_ = chain.Load(c, stateFile) // Ignore error if file doesn't exist (new chain)

	switch cmd {
	case "start":
		commands.Start(c, args)
	case "stop":
		commands.Stop(c, args)
	case "restart":
		commands.Restart(c, args)
	case "status":
		StatusCmd(args)
	case "new-wallet":
		commands.NewWallet(c, args)
	case "send":
		commands.Send(c, args)
	default:
		fmt.Printf("[ERROR] Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}

	// После выполнения команды — сохранить актуальное состояние:
	ns := &nodeState{
		Chain:     c.Blocks,
		Mempool:   c.Mempool,
		Wallets:   []*wallet.Wallet{},
		Balances:  c.Balances,
		StartTime: "", // если нужно, можно добавить
	}
	for _, w := range c.Wallets {
		ww := &wallet.Wallet{Wallet: w}
		ns.Wallets = append(ns.Wallets, ww)
	}
	saveState(ns)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  start                     Start the blockchain node")
	fmt.Println("  stop                      Stop the blockchain node")
	fmt.Println("  restart                   Restart the blockchain node")
	fmt.Println("  status                    Show node status")
	fmt.Println("  new-wallet <seed>         Create a new wallet using a 12-word seed phrase")
	fmt.Println("  send <from> <to> <amount> Send coins from one wallet to another")
}

func StartCmd(args []string) {
	var s *nodeState
	if _, err := os.Stat(getStateFilePath()); err == nil {
		s = loadState()
	} else {
		s = defaultState()
	}
	s.StartTime = time.Now().Format(time.RFC3339)
	saveState(s)
	slog("Node started", s)
	fmt.Printf("[INFO] Node started | height: %d | wallets: %d\n", len(s.Chain)-1, len(s.Wallets))

	// --- Daemon: block timer and signal handling ---
	startBlockTimer()
	fmt.Println("[INFO] Node is running. Press Ctrl+C to stop.")

	// Wait for interrupt (Ctrl+C/SIGINT/SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("\n[INFO] Shutting down node...")
	stopBlockTimer()
	slog("Node stopped", s)
	fmt.Println("[INFO] Node stopped")
}

func StopCmd(args []string) {
	s := loadState()
	stopBlockTimer()
	slog("Node stopped", s)
	fmt.Printf("[INFO] Node stopped\n")
}

func RestartCmd(args []string) {
	StopCmd(args)
	StartCmd(args)
}

func NewWalletCmd(args []string) {
	s := loadState()
	if len(args) == 0 || strings.TrimSpace(args[0]) == "" {
		err := "Seed phrase required. Use --seed 'word1 word2 ... word12'"
		logger.Logger.Error("Wallet creation failed", "error", err)
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	words := strings.Fields(args[0])
	if len(words) != 12 {
		err := "Seed phrase must be exactly 12 words."
		logger.Logger.Error("Wallet creation failed", "error", err)
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	seed := sha256.Sum256([]byte(strings.Join(words, " ")))
	pub, _, err := ed25519.GenerateKey(strings.NewReader(string(seed[:])))
	if err != nil {
		logger.Logger.Error("Wallet creation failed", "error", err.Error())
		fmt.Printf("[ERROR] Failed to generate key: %s\n", err.Error())
		return
	}
	address := hex.EncodeToString(pub[:8])
	for _, w := range s.Wallets {
		if string(w.Address) == address {
			fmt.Printf("Wallet already exists!\nPublic address: %s\n(Your wallet is ready to use. Keep your seed phrase safe!)\n", address)
			saveState(s)
			return
		}
	}
	s.Wallets = append(s.Wallets, &wallet.Wallet{
		Wallet: types.Wallet{
			Address: types.Address(address),
			PubKey:  pub,
			Balance: 100.0,
		},
	})
	if s.Balances == nil {
		s.Balances = map[string]float64{}
	}
	s.Balances[address] = 100.0 // Give new wallet a mock balance
	saveState(s)
	slog("Wallet created", s)
	fmt.Printf("Wallet created!\nPublic address: %s\n(Your wallet is ready to use. Keep your seed phrase safe!)\n", address)
}

func SendCmd(args []string) {
	s := loadState()
	if len(args) < 3 {
		err := "Missing required flags: --from, --to, --amount"
		logger.Logger.Error("Send failed", "error", err)
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	from := args[0]
	to := args[1]
	amountStr := args[2]
	if from == "" || to == "" || amountStr == "" {
		err := "All of --from, --to, --amount are required."
		logger.Logger.Error("Send failed", "error", err)
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		errMsg := "Amount must be a positive number."
		logger.Logger.Error("Send failed", "error", errMsg)
		fmt.Printf("[ERROR] %s\n", errMsg)
		return
	}
	fromBal := s.Balances[from]
	if fromBal < amount {
		err := "Insufficient balance."
		logger.Logger.Error("Send failed", "error", err)
		fmt.Printf("[ERROR] %s\n", err)
		return
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
	s.Mempool = append(s.Mempool, tx)
	saveState(s)
	slog("Transaction added to mempool", s)
	fmt.Printf("Sent %.2f tokens from %s to %s\nNew balance:\n  %s: %.2f\n  %s: %.2f\n",
		amount, from, to, from, fromBal, to, toBal)
}

func StatusCmd(args []string) {
	s := loadState()
	now := time.Now()
	uptime := ""
	if s.StartTime != "" {
		if t, err := time.Parse(time.RFC3339, s.StartTime); err == nil {
			uptime = formatUptime(now.Sub(t))
		}
	}
	// --- New: time to next block ---
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
	logger.Logger.Info("Status",
		"height", len(s.Chain)-1,
		"mempool", len(s.Mempool),
		"wallets", len(s.Wallets),
		"time", now.Format(time.RFC3339),
		"uptime", uptime,
		"next_block_in", timeToNextBlock,
	)
	fmt.Printf("Block height: %d\nMempool txs: %d\nWallets: %d\nTime: %s\nUptime: %s\nNext block in: %s\n",
		len(s.Chain)-1, len(s.Mempool), len(s.Wallets), now.Format(time.RFC3339), uptime, timeToNextBlock)
}

func formatUptime(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}

func LoadChainState() *types.Chain {
	c := chain.NewChain()
	_ = chain.Load(c, stateFile)
	return c
}

// GenerateMerkleProof returns the Merkle proof for a transaction in a block
func GenerateMerkleProof(c *types.Chain, blockIdx, txIdx int) ([]string, error) {
	if blockIdx < 0 || blockIdx >= len(c.Blocks) {
		return nil, errors.New("block index out of range")
	}
	txs := c.Blocks[blockIdx].Transactions
	if txIdx < 0 || txIdx >= len(txs) {
		return nil, errors.New("transaction index out of range")
	}
	var proof []string
	hashes := make([][]byte, len(txs))
	for i, tx := range txs {
		hashes[i] = []byte(tx.Hash())
	}
	idx := txIdx
	for len(hashes) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			var left, right []byte
			left = hashes[i]
			if i+1 < len(hashes) {
				right = hashes[i+1]
			} else {
				right = left
			}
			combined := append(left, right...)
			h := types.Hash(fmt.Sprintf("%x", crypto.HashData(combined)))
			nextLevel = append(nextLevel, []byte(h))
			if i == idx || i+1 == idx {
				if i == idx && i+1 < len(hashes) {
					proof = append(proof, string(hashes[i+1]))
					idx = i / 2
				} else if i+1 == idx {
					proof = append(proof, string(hashes[i]))
					idx = i / 2
				}
			}
		}
		hashes = nextLevel
	}
	return proof, nil
}

// VerifyMerkleProof verifies a Merkle proof for a transaction
func VerifyMerkleProof(c *types.Chain, blockIdx int, txHash string, proof []string) (bool, error) {
	if blockIdx < 0 || blockIdx >= len(c.Blocks) {
		return false, errors.New("block index out of range")
	}
	root := c.Blocks[blockIdx].MerkleRoot
	h := []byte(txHash)
	for _, p := range proof {
		combined := append(h, []byte(p)...)
		h = []byte(fmt.Sprintf("%x", crypto.HashData(combined)))
	}
	return string(h) == string(root), nil
}

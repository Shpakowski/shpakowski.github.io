package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"mcp-coop-chain/internal"
	"mcp-coop-chain/internal/blockchain"
	"mcp-coop-chain/internal/contracts/proto"
	"mcp-coop-chain/internal/logging"
	"mcp-coop-chain/internal/storage"
	"mcp-coop-chain/internal/types"
	"mcp-coop-chain/internal/wallet"
)

// loadConfig — загрузка конфига из стандартного пути
func loadConfig() *types.BlockchainConfig {
	cfg, err := internal.LoadConfig("configs/blockchain_config.json")
	if err != nil {
		fmt.Println("[Ошибка конфига]", err)
		os.Exit(1)
	}
	return cfg
}

// newNode — создаёт FullNode с MemoryStorage и логгером
func newNode(cfg *types.BlockchainConfig) *blockchain.FullNode {
	logger, _ := logging.NewLogger(cfg.Logger)
	memStorage := storage.NewMemoryStorage()
	return blockchain.NewFullNode(cfg, memStorage, logger)
}

func main() {
	proto.RegisterAllProtoContracts()
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "init":
		Init()
	case "start":
		Start()
	case "stop":
		Stop()
	case "addwallet":
		addWalletCmd := flag.NewFlagSet("addwallet", flag.ExitOnError)
		mnemonic := addWalletCmd.String("mnemonic", "", "Seed-фраза для генерации (опционально)")
		_ = addWalletCmd.Parse(os.Args[2:])
		AddWallet(*mnemonic)
	case "tx":
		txCmd := flag.NewFlagSet("tx", flag.ExitOnError)
		from := txCmd.String("from", "", "Адрес отправителя")
		to := txCmd.String("to", "", "Адрес получателя")
		amount := txCmd.Uint64("amount", 0, "Сумма перевода (в микро-MCP)")
		_ = txCmd.Parse(os.Args[2:])
		if *from == "" || *to == "" || *amount == 0 {
			fmt.Println("Использование: cli tx --from <адрес> --to <адрес> --amount <число>")
			return
		}
		Transaction(*from, *to, *amount)
	case "deploy":
		deployCmd := flag.NewFlagSet("deploy", flag.ExitOnError)
		name := deployCmd.String("name", "", "Имя контракта")
		body := deployCmd.String("body", "", "Тело контракта")
		author := deployCmd.String("author", "", "Адрес автора")
		_ = deployCmd.Parse(os.Args[2:])
		if *name == "" || *body == "" || *author == "" {
			fmt.Println("Использование: cli deploy --name <имя> --body <код> --author <адрес>")
			return
		}
		DeploySmartContract(*name, *body, *author)
	case "exec":
		execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
		contract := execCmd.String("contract", "", "Имя контракта")
		params := execCmd.String("params", "", "Параметры (через запятую)")
		_ = execCmd.Parse(os.Args[2:])
		if *contract == "" {
			fmt.Println("Использование: cli exec --contract <имя> --params <параметры>")
			return
		}
		paramList := []string{}
		if *params != "" {
			paramList = strings.Split(*params, ",")
		}
		ExecuteSmartContract(*contract, paramList...)
	case "status":
		Status()
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Использование: cli <команда> [флаги]")
	fmt.Println("Доступные команды:")
	fmt.Println("  init")
	fmt.Println("  start")
	fmt.Println("  stop")
	fmt.Println("  addwallet [--mnemonic <seed>]")
	fmt.Println("  tx --from <адрес> --to <адрес> --amount <число>")
	fmt.Println("  deploy --name <имя> --body <код> --author <адрес>")
	fmt.Println("  exec --contract <имя> --params <параметры через запятую>")
	fmt.Println("  status")
}

// Init — инициализация блокчейна (Genesis-блок, первый снапшот)
func Init() {
	cfg := loadConfig()
	node := newNode(cfg)
	err := node.InitGenesisBlock(node.Logger)
	if err != nil {
		fmt.Println("[Ошибка]", err)
		return
	}
	fmt.Println("Genesis-блок создан, снапшот записан.")
}

// Start — запуск FullNode (таймер, mempool, блокогенерация)
func Start() {
	cfg := loadConfig()
	node := newNode(cfg)
	node.Start()
	fmt.Println("FullNode запущен.")
}

// Stop — корректная остановка узла, фиксация снапшота
func Stop() {
	cfg := loadConfig()
	node := newNode(cfg)
	node.Stop()
	fmt.Println("Узел остановлен, снапшот сохранён.")
}

// AddWallet — создать новый кошелёк через FullNode, вернуть адрес
func AddWallet(mnemonic string) {
	cfg := loadConfig()
	node := newNode(cfg)
	priv, mnemonicOut, err := wallet.CreateWalletWithMnemonic(mnemonic)
	if err != nil {
		fmt.Println("[Ошибка]", err)
		return
	}
	pub := wallet.ToPublicWallet(priv)
	if ws, ok := node.Storage.(types.WalletStorage); ok {
		err = ws.AddWallet(pub)
		if err != nil {
			fmt.Println("[Ошибка]", err)
			return
		}
	}
	b, _ := json.MarshalIndent(pub, "", "  ")
	fmt.Println("Кошелёк создан:")
	fmt.Println("  Адрес:", pub.Address)
	fmt.Println("  Публичный ключ:", pub.PublicKey)
	fmt.Println("  Сид-фраза (mnemonic):", mnemonicOut)
	fmt.Println(string(b))
	// Приватный кошелёк можно сохранить локально при необходимости (например, через wallet.SaveWallet)
}

// Transaction — создать, подписать и отправить перевод между кошельками
func Transaction(from, to string, amount uint64) {
	cfg := loadConfig()
	node := newNode(cfg)
	wallets, _ := node.Storage.(types.WalletStorage)
	args := types.TransferArgs{From: from, To: to, Amount: amount}
	b, _ := json.Marshal(args)
	res, err := proto.CallProtoContract("Transfer", wallets, nil, b)
	if err != nil {
		fmt.Println("[Ошибка]", err)
		return
	}
	fmt.Println("Результат перевода:")
	fmt.Println(string(res))
}

// DeploySmartContract — зарегистрировать Proto-контракт
func DeploySmartContract(name, code, author string) {
	cfg := loadConfig()
	node := newNode(cfg)
	contracts, _ := node.Storage.(types.ContractStorage)
	args := types.AddSmartContractArgs{Name: name, Code: code, Author: author}
	b, _ := json.Marshal(args)
	res, err := proto.CallProtoContract("AddSmartContract", nil, contracts, b)
	if err != nil {
		fmt.Println("[Ошибка]", err)
		return
	}
	fmt.Println("Контракт зарегистрирован:")
	fmt.Println(string(res))
}

// ExecuteSmartContract — выполнить зарегистрированный контракт
func ExecuteSmartContract(name string, params ...string) {
	cfg := loadConfig()
	node := newNode(cfg)
	contracts, _ := node.Storage.(types.ContractStorage)
	wallets, _ := node.Storage.(types.WalletStorage)
	b, _ := json.Marshal(params)
	res, err := proto.CallProtoContract(name, wallets, contracts, b)
	if err != nil {
		fmt.Println("[Ошибка]", err)
		return
	}
	fmt.Println("Результат контракта:")
	fmt.Println(string(res))
}

// Status — выводит актуальное состояние FullNode (цепочка, кошельки, mempool)
func Status() {
	cfg := loadConfig()
	node := newNode(cfg)
	snapshot := node.BuildFullSnapshot()
	fmt.Println("=== MCP Coop Chain Status ===")
	fmt.Printf("Блоков в цепочке: %d\n", len(snapshot.Blocks))
	if len(snapshot.Blocks) > 0 {
		fmt.Printf("Высота: %d\n", snapshot.Blocks[len(snapshot.Blocks)-1].Header.Height)
	}
	fmt.Printf("Кошельков: %d\n", len(snapshot.Wallets))
	if len(snapshot.Wallets) > 0 {
		fmt.Println("Адреса кошельков:")
		for _, w := range snapshot.Wallets {
			fmt.Printf("  - %s\n", w.Address)
		}
	}
	fmt.Printf("Mempool: %d транзакций\n", len(snapshot.Mempool))
	fmt.Printf("Время снапшота: %s\n", snapshot.Timestamp.Format(time.RFC3339))
}

package types

// Storage определяет интерфейс для хранения блоков и работы со слепком.
type Storage interface {
	AddBlock(block Block) error
	GetBlockByHash(hash string) (Block, error)
	GetLastBlock() (Block, error)
	HasBlock(hash string) bool
	GetAllBlocks() []Block
	SaveState(data []byte) error
	LoadState() ([]byte, error)
}

// WalletStorage определяет интерфейс для управления кошельками и балансами
// Используется Proto API для работы с MCP
// Все суммы в микро-MCP (uint64)
type WalletStorage interface {
	GetWallet(address string) (Wallet, error)
	AddWallet(wallet Wallet) error
	UpdateBalance(address string, delta int64) error // delta может быть отрицательным
	GetBalance(address string) (uint64, error)
}

// ContractStorage определяет интерфейс для управления пользовательскими контрактами
// Используется Proto API для AddSmartContract и др.
type ContractStorage interface {
	AddContract(name, code, author string) error
	GetContract(name string) (string, error)
}

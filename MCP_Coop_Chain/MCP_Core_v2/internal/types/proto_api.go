package types

// TransferArgs — аргументы для Transfer(from, to, amount)
type TransferArgs struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

type TransferResult struct {
	Success bool   `json:"success"`
	TxID    string `json:"txId,omitempty"`
	Error   string `json:"error,omitempty"`
}

// GetBalanceArgs — аргументы для GetBalance(wallet)
type GetBalanceArgs struct {
	Wallet string `json:"wallet"`
}

type GetBalanceResult struct {
	Balance uint64 `json:"balance"`
}

// CreateWalletArgs — аргументы для CreateWallet(pubKey)
type CreateWalletArgs struct {
	PublicKey string `json:"publicKey"`
}

type CreateWalletResult struct {
	Address string `json:"address"`
}

// AddSmartContractArgs — аргументы для AddSmartContract(name, code, author)
type AddSmartContractArgs struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Author string `json:"author"`
}

type AddSmartContractResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

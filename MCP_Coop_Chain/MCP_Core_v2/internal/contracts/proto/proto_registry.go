package proto

// RegisterAllProtoContracts регистрирует все встроенные proto-контракты
func RegisterAllProtoContracts() {
	RegisterProtoContract("Transfer", &TransferContract{})
	RegisterProtoContract("GetBalance", &GetBalanceContract{})
	RegisterProtoContract("CreateWallet", &CreateWalletContract{})
	RegisterProtoContract("AddSmartContract", &AddSmartContractContract{})
}

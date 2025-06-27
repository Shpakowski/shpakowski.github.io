package services

import (
	"crypto/ed25519"
	"errors"
	"mcp-chain/core/blockchain"
	"mcp-chain/core/mempool"
	"mcp-chain/core/state"
	"mcp-chain/core/transaction"
	"mcp-chain/core/wallet"
	"mcp-chain/crypto"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
	"mcp-chain/types"

	"github.com/tyler-smith/go-bip39"
)

// SendTx инкапсулирует всю бизнес-логику отправки транзакции через ядро
// fromInput — seed-фраза или публичный ключ, to — адрес, amount/fee — uint64
func SendTx(fromInput string, to types.Address, amount, fee uint64) (types.Hash, error) {
	logging.Logger.Info("SendTx", "from", fromInput, "to", to, "amount", amount, "fee", fee)
	st := state.GlobalState
	var privKey ed25519.PrivateKey
	var pubKey ed25519.PublicKey
	var fromAddr types.Address
	if wallet.IsMnemonic(fromInput) {
		seedBytes := bip39.NewSeed(fromInput, "")
		privKey = ed25519.NewKeyFromSeed(seedBytes[:32])
		pubKey = privKey.Public().(ed25519.PublicKey)
		fromAddr = types.Address(crypto.PublicKeyToHex(pubKey))
	} else {
		mgr := wallet.NewManager(wallet.WalletsFile)
		found := false
		for _, w := range mgr.ListWallets() {
			if w.PubKey == fromInput {
				seedBytes := bip39.NewSeed(w.Seed, "")
				privKey = ed25519.NewKeyFromSeed(seedBytes[:32])
				pubKey = privKey.Public().(ed25519.PublicKey)
				fromAddr = types.Address(w.PubKey)
				found = true
				break
			}
		}
		if !found {
			return "", errors.New("Не найден seed для указанного публичного ключа")
		}
	}

	acc := st.GetAccount(fromAddr)
	var nonce uint64
	if acc != nil {
		nonce = acc.Nonce + 1
	} else {
		nonce = 1
	}
	tx := &transaction.Tx{
		From:   fromAddr,
		To:     to,
		Amount: amount,
		Fee:    fee,
		Nonce:  nonce,
	}
	tx.Sign(privKey)

	cfg := config.GetDefaultNetworkConfig()
	if err := transaction.ValidateTx(tx, &cfg, pubKey); err != nil {
		return "", err
	}

	txForMempool := &types.Transaction{
		ID:     tx.ID,
		From:   tx.From,
		To:     tx.To,
		Amount: types.Amount(tx.Amount),
		Fee:    types.Amount(tx.Fee),
		Nonce:  tx.Nonce,
		Sig:    tx.Sig,
	}
	mempool.GlobalTxMempool.AddTx(txForMempool)
	return tx.ID, nil
}

// SendTxAndSave отправляет транзакцию, применяет к state, сохраняет state, логирует
func SendTxAndSave(chain *blockchain.Chain, cfg *config.NetworkConfig, fromInput string, to types.Address, amount, fee uint64) (types.Hash, error) {
	txID, err := SendTx(fromInput, to, amount, fee)
	if err != nil {
		logging.Logger.Error("Ошибка отправки транзакции", "err", err)
		return txID, err
	}
	logging.Logger.Info("SendTxAndSave: транзакция отправлена", "txID", txID, "from", fromInput, "to", to, "amount", amount, "fee", fee)
	return txID, nil
}

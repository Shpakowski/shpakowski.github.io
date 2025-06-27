package transaction

import (
	"crypto/ed25519"
	"errors"
	"mcp-chain/contracts"
	"mcp-chain/core/state"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
)

// ErrTxInvalid возвращает ошибку с кодом причины
func ErrTxInvalid(reason string) error {
	return errors.New("invalid tx: " + reason)
}

// ValidateTx проверяет корректность транзакции
func ValidateTx(tx *Tx, cfg *config.NetworkConfig, pubKey ed25519.PublicKey) error {
	logging.Logger.Debug("ValidateTx", "from", tx.From, "to", tx.To, "amount", tx.Amount, "fee", tx.Fee, "nonce", tx.Nonce)
	// 1. Проверка подписи
	if !tx.Verify(pubKey) {
		return ErrTxInvalid("bad signature")
	}
	// 2. Проверка Nonce
	acc := state.GlobalState.GetAccount(tx.From)
	var nonce uint64
	if acc != nil {
		nonce = acc.Nonce
	} else {
		nonce = 0
	}
	if tx.Nonce != nonce+1 {
		logging.Logger.Warn("ValidateTx: bad nonce", "expected", nonce+1, "got", tx.Nonce)
		return ErrTxInvalid("bad nonce")
	}
	// 3. Проверка баланса
	balance := state.GlobalState.GetBalance(tx.From)
	if uint64(balance) < tx.Amount+tx.Fee {
		logging.Logger.Warn("ValidateTx: insufficient balance", "balance", balance, "need", tx.Amount+tx.Fee)
		return ErrTxInvalid("insufficient balance")
	}
	// 4. Проверка комиссии
	if tx.Fee < uint64(cfg.MinFee) {
		logging.Logger.Warn("ValidateTx: fee too low", "fee", tx.Fee, "min", cfg.MinFee)
		return ErrTxInvalid("fee too low")
	}
	// 5. Проверка Payload (если есть)
	if len(tx.Payload) > 0 {
		// TODO: распарсить имя и параметры контракта из Payload
		_, err := contracts.CallProtoContract("CreateCoop", map[string]interface{}{})
		if err != nil {
			return ErrTxInvalid("payload: " + err.Error())
		}
	}
	// Проверить обращения к Proposal.Votes, если есть: Votes[string(addr)]
	return nil
}

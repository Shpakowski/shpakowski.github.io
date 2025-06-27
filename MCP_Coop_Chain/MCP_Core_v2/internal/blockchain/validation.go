package blockchain

import (
	"encoding/json"
	"errors"
	"mcp-coop-chain/internal/types"
	"mcp-coop-chain/internal/wallet"
)

// ValidateTransactionForMempool проверяет транзакцию перед добавлением в mempool
// state — актуальный state (map[address]balance), mempool — текущий mempool
func ValidateTransactionForMempool(tx *types.Transaction, state map[string]uint64, mempool []*types.Transaction) error {
	// 1. Проверка подписи
	if !wallet.VerifyTransaction(*tx, nil) {
		return errors.New("invalid signature")
	}
	// 2. Проверка формата (валидные адреса, положительная сумма и fee)
	var transfer types.TransferArgs
	if err := json.Unmarshal(tx.Payload, &transfer); err == nil && transfer.From != "" && transfer.To != "" {
		if transfer.Amount == 0 {
			return errors.New("amount must be positive")
		}
		if tx.Fee < 0 {
			return errors.New("fee must be non-negative")
		}
		if !wallet.IsValidAddress(transfer.From) || !wallet.IsValidAddress(transfer.To) {
			return errors.New("invalid address format")
		}
		// 3. Проверка баланса
		fee := uint64(tx.Fee * 1_000_000)
		total := transfer.Amount + fee
		if state[transfer.From] < total {
			return errors.New("insufficient balance")
		}
	} else {
		// Для других типов транзакций — только fee и валидность адреса отправителя
		fee := uint64(tx.Fee * 1_000_000)
		if tx.Origin == "" || !wallet.IsValidAddress(tx.Origin) {
			return errors.New("invalid origin address")
		}
		if state[tx.Origin] < fee {
			return errors.New("insufficient balance for fee")
		}
	}
	// 4. Проверка на дубликат TxID в mempool
	for _, mtx := range mempool {
		if mtx.TxID == tx.TxID {
			return errors.New("duplicate TxID in mempool")
		}
	}
	return nil
}

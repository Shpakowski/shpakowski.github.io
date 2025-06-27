package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"mcp-coop-chain/internal"
	"mcp-coop-chain/internal/types"
	"time"
)

// Transaction содержит методы создания и подписи транзакций перевода MCP
// TODO: реализовать создание и подпись транзакций

// Минимальная комиссия в MCP
// const MinFee = 0.001 // Удалено, теперь из конфига

// NewTransaction создаёт новую транзакцию, рассчитывает TxID, проставляет время и nonce.
func NewTransaction(origin string, payload []byte, contractHashes []string, fee float64, reason string, nonce uint64) (types.Transaction, error) {
	if origin != "User" && origin != "Oracle" {
		return types.Transaction{}, errors.New("origin должен быть 'User' или 'Oracle'")
	}
	if len(payload) == 0 {
		return types.Transaction{}, errors.New("payload не может быть пустым")
	}
	if fee < internal.CurrentConfig.Fees.TxFee {
		return types.Transaction{}, errors.New("комиссия ниже минимальной: см. config")
	}
	// TODO: Проверка уникальности nonce для адреса отправителя (заглушка)
	tx := types.Transaction{
		Origin:              origin,
		Payload:             payload,
		SmartContractHashes: contractHashes,
		Fee:                 fee,
		Reason:              reason,
		Timestamp:           time.Now().UTC(),
		Nonce:               nonce,
	}
	tx.TxID = HashTransaction(tx)
	return tx, nil
}

// SignTransaction подписывает транзакцию приватным ключом (ECDSA) и обновляет Signature.
// Требует, чтобы приватный ключ был валиден и расшифрован.
// Если комиссия ниже минимальной или nonce неуникален — транзакция не должна быть создана (см. NewTransaction).
func SignTransaction(tx *types.Transaction, priv *ecdsa.PrivateKey) error {
	hash := sha256.Sum256(tx.Payload)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return err
	}
	sig := append(r.Bytes(), s.Bytes()...)
	tx.Signature = base64.StdEncoding.EncodeToString(sig)
	return nil
}

// VerifyTransaction проверяет подпись и целостность транзакции.
// Проверяет, что подпись валидна, TxID соответствует содержимому, комиссия не ниже минимума.
// Проверка уникальности nonce — заглушка (реализовать в обработчике блока).
func VerifyTransaction(tx types.Transaction, pub *ecdsa.PublicKey) bool {
	hash := sha256.Sum256(tx.Payload)
	sig, err := base64.StdEncoding.DecodeString(tx.Signature)
	if err != nil || len(sig) < 64 {
		return false
	}
	r := new(big.Int).SetBytes(sig[:len(sig)/2])
	s := new(big.Int).SetBytes(sig[len(sig)/2:])
	if !ecdsa.Verify(pub, hash[:], r, s) {
		return false
	}
	if tx.Fee < internal.CurrentConfig.Fees.TxFee {
		return false
	}
	return tx.TxID == HashTransaction(tx)
}

// HashTransaction сериализует ключевые поля и возвращает SHA256-хеш как hex-строку.
func HashTransaction(tx types.Transaction) string {
	data := struct {
		Origin              string
		Payload             []byte
		SmartContractHashes []string
		Fee                 float64
		Reason              string
		Timestamp           time.Time
	}{
		Origin:              tx.Origin,
		Payload:             tx.Payload,
		SmartContractHashes: tx.SmartContractHashes,
		Fee:                 tx.Fee,
		Reason:              tx.Reason,
		Timestamp:           tx.Timestamp,
	}
	b, _ := json.Marshal(data)
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

// IsOracleTransaction возвращает true, если транзакция от Oracle.
func IsOracleTransaction(tx types.Transaction) bool {
	return tx.Origin == "Oracle"
}

// AffectedContracts возвращает список смарт-контрактов, к которым обращается транзакция.
func AffectedContracts(tx types.Transaction) []string {
	return tx.SmartContractHashes
}

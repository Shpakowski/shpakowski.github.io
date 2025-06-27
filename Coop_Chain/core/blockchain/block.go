package blockchain

import (
	"crypto/ed25519"
	"encoding/json"
	"mcp-chain/core/state"
	"mcp-chain/core/transaction"
	"mcp-chain/crypto"
	"mcp-chain/internal/config"
	"mcp-chain/types"
	"time"
)

// Все типы блока теперь импортируются из types. Функция хэширования вынесена в hash.go.

// ConsensusData — данные для консенсуса (Proof-of-Cooperation)
type ConsensusData struct {
	TotalStake     uint64 // Общий stake в сети
	ProposerRating uint64 // Рейтинг валидатора
	SlotTime       uint64 // Время слота (секунды)
}

// Header — заголовок блока
type Header struct {
	Height        uint64        // Высота блока
	PrevHash      types.Hash    // Хэш предыдущего блока
	TxRoot        types.Hash    // MerkleRoot всех транзакций
	StateRoot     types.Hash    // Хэш состояния после применения блока
	Timestamp     int64         // Время создания блока (unix)
	Proposer      types.Address // Адрес валидатора
	Sign          []byte        // Ed25519 подпись хэша Header
	ConsensusData ConsensusData // Данные для консенсуса
}

// Block — структура блока
// Содержит заголовок и список транзакций
// Хэш блока = SHA-256(JSON(Header))
type Block struct {
	Header Header
	Txs    []transaction.Tx
}

// UpdateRoots пересчитывает TxRoot и StateRoot
func (b *Block) UpdateRoots(cfg *config.NetworkConfig) {
	b.Header.TxRoot = calcTxMerkleRoot(b.Txs)
	b.Header.StateRoot = calcStateRoot()
}

// CalcHash вычисляет SHA-256(JSON(Header))
func (b *Block) CalcHash(cfg *config.NetworkConfig) types.Hash {
	headCopy := b.Header
	headCopy.Sign = nil // Не включаем подпись в хэш
	data, _ := json.Marshal(headCopy)
	return types.Hash(crypto.HashBytes(data))
}

// Sign подписывает хэш Header приватным ключом Ed25519
func (b *Block) SignBlock(privKey ed25519.PrivateKey, cfg *config.NetworkConfig) {
	hash := b.CalcHash(cfg)
	b.Header.Sign = ed25519.Sign(privKey, []byte(hash))
}

// Verify проверяет целостность блока и подпись валидатора
func (b *Block) Verify(prev *Block, cfg *config.NetworkConfig, pubKey ed25519.PublicKey) error {
	if b.Header.PrevHash != prev.CalcHash(cfg) {
		return ErrBlockInvalidPrevHash
	}
	stake := state.GlobalState.Validators[string(b.Header.Proposer)]
	if stake.Stake < types.Amount(cfg.MinStake) {
		return ErrBlockInvalidStake
	}
	if !ed25519.Verify(pubKey, []byte(b.CalcHash(cfg)), b.Header.Sign) {
		return ErrBlockInvalidSignature
	}
	timestamp := time.Now().Unix()
	if b.Header.Timestamp > (timestamp + cfg.MaxDrift) {
		return ErrBlockInvalidTimestamp
	}
	return nil
}

// Size возвращает размер блока в байтах
func (b *Block) Size() int {
	data, _ := json.Marshal(b)
	return len(data)
}

// calcTxMerkleRoot вычисляет MerkleRoot для списка транзакций
func calcTxMerkleRoot(txs []transaction.Tx) types.Hash {
	hashes := make([][]byte, len(txs))
	for i, tx := range txs {
		hashes[i] = []byte(tx.ID)
	}
	root := crypto.MerkleRoot(hashes)
	return types.Hash(root)
}

// calcStateRoot вычисляет SHA-256 от сериализованного состояния
func calcStateRoot() types.Hash {
	data, _ := json.Marshal(state.GlobalState)
	return types.Hash(crypto.HashBytes(data))
}

// Ошибки блока
var (
	ErrBlockInvalidPrevHash  = types.NewBlockError("prev hash mismatch")
	ErrBlockInvalidStake     = types.NewBlockError("validator stake too low")
	ErrBlockInvalidSignature = types.NewBlockError("invalid block signature")
	ErrBlockInvalidTimestamp = types.NewBlockError("block timestamp too far in future")
)

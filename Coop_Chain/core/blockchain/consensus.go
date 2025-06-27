package blockchain

import (
	"errors"
	"mcp-chain/core/state"
	"mcp-chain/crypto"
	"mcp-chain/internal/config"
	mcpErrors "mcp-chain/internal/errors"
	"mcp-chain/types"
)

// Проверяет, что у валидатора достаточно стейка
func HasSufficientStake(addr types.Address) bool {
	stake, ok := state.GlobalState.Validators[string(addr)]
	return ok && float64(stake.Stake) >= config.GetDefaultNetworkConfig().MinValidatorStake
}

// Проверяет подпись блока валидатором
func VerifyBlockSignature(block *types.Block, pubKey []byte, sig []byte) bool {
	// TODO: сериализовать Header и проверить подпись Ed25519
	return crypto.Verify(pubKey, []byte(block.Header.Proposer), sig)
}

// Проверяет кворум (MVP: всегда true, т.к. один валидатор)
func CheckQuorum() bool {
	// В MVP кворум всегда есть
	return true
}

// ValidateBlock выполняет все проверки блока
func ValidateBlock(block *types.Block, pubKey []byte, sig []byte) error {
	if !HasSufficientStake(block.Header.Proposer) {
		return mcpErrors.ErrInsufficientStake
	}
	if !VerifyBlockSignature(block, pubKey, sig) {
		return mcpErrors.ErrInvalidSignature
	}
	if !CheckQuorum() {
		return mcpErrors.ErrInvalidTx // Можно завести отдельную ошибку
	}
	return nil
}

// ConsensusRule — интерфейс для правил консенсуса
// Позволяет внедрять разные алгоритмы (PoC, PoS, ...)
type ConsensusRule interface {
	ValidateBlock(prev, new *Block, cfg *config.NetworkConfig) error
}

// ErrConsensusViolation — ошибка нарушения правил консенсуса
var ErrConsensusViolation = errors.New("consensus rule violation")

// ProofOfCooperationRule — реализация Proof-of-Cooperation
// Единственный валидатор с LOCKED stake ≥ cfg.MinStake
// Height = prev.Height+1
// Блок своевременный, если прошло > AutoMineInterval или txCount ≥ MaxTxPerBlock

type ProofOfCooperationRule struct{}

func (p *ProofOfCooperationRule) ValidateBlock(prev, new *Block, cfg *config.NetworkConfig) error {
	// 1. Проверка валидатора
	stake := state.GlobalState.Validators[string(new.Header.Proposer)]
	if stake.Stake < types.Amount(cfg.MinStake) {
		return ErrConsensusViolation
	}
	// 2. Проверка высоты
	if new.Header.Height != prev.Header.Height+1 {
		return ErrConsensusViolation
	}
	// 3. Проверка своевременности блока
	txCount := len(new.Txs)
	blockInterval := new.Header.Timestamp - prev.Header.Timestamp
	if !(blockInterval > int64(cfg.AutoMineInterval) || txCount >= cfg.MaxTxPerBlock) {
		return ErrConsensusViolation
	}
	return nil
}

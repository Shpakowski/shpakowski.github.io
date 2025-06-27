package blockchain

import (
	"bytes"
	"mcp-chain/core/staking"
	"mcp-chain/core/state"
	"mcp-chain/core/transaction"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
	"mcp-chain/types"
)

// CreateGenesis создает генезис-блок (Height=0)
// Любые изменения параметров генезиса требуют смены ID сети!
func CreateGenesis(cfg *config.NetworkConfig) *Block {
	if state.GlobalState == nil {
		logging.Logger.Info("CreateGenesis: state.GlobalState не инициализирован, создаю новый state")
		state.GlobalState = state.NewState()
	} else {
		logging.Logger.Info("CreateGenesis: state.GlobalState уже инициализирован, не пересоздаю")
	}
	logging.Logger.Info("CreateGenesis: создание genesis-блока")

	// 1. PrevHash = 32 байт нулей
	prevHash := types.Hash(bytes.Repeat([]byte{0}, 32))

	// 2. Одна транзакция GENESIS_REWARD на cfg.GenesisReward ($MCP)
	genesisTx := transaction.Tx{
		From:    types.Address("GENESIS"),
		To:      types.Address(cfg.ValidatorAddress),
		Amount:  uint64(types.Amount(cfg.GenesisReward)),
		Fee:     0,
		Nonce:   1,
		Payload: nil,
	}
	genesisTx.ID = genesisTx.CalcID()

	// 3. Лочим stake валидатора на cfg.MinStake
	state.GlobalState.Accounts[string(cfg.ValidatorAddress)] = types.Account{Balance: types.Amount(cfg.GenesisReward), Nonce: 1}
	staking.LockStake(types.Address(cfg.ValidatorAddress), uint64(cfg.MinStake), 0)

	// 4. Формируем блок
	head := Header{
		Height:    0,
		PrevHash:  prevHash,
		Timestamp: 0,
		Proposer:  types.Address(cfg.ValidatorAddress),
	}
	block := &Block{
		Header: head,
		Txs:    []transaction.Tx{genesisTx},
	}
	block.UpdateRoots(cfg)
	logging.Logger.Info("CreateGenesis: genesis-блок создан")
	return block
}

package staking

import (
	"mcp-chain/core/state"
	"mcp-chain/core/wallet"
	"mcp-chain/internal/config"
	"mcp-chain/logging"
	"mcp-chain/types"
	"time"
)

// LockStake — вычитает amt с адреса и помещает в StakeInfo
func LockStake(addr types.Address, amt uint64, until int64) {
	logging.Logger.Info("LockStake: ДО", "address", addr, "amount", amt, "until", until)
	_ = state.ApplyStateChange(func(st *state.State) {
		acc := st.Accounts[string(addr)]
		if acc.Balance >= types.Amount(amt) {
			acc.Balance -= types.Amount(amt)
			info := st.Validators[string(addr)]
			info.Stake += types.Amount(amt)
			info.LockedUntil = uint64(until)
			st.Validators[string(addr)] = info
			st.Accounts[string(addr)] = acc
			logging.Logger.Info("LockStake: ПОСЛЕ", "account", acc, "validator", info)
		}
	})
}

// IsActive — true, если Amount ≥ cfg.MinStake и SlashCount < 3
func IsActive(addr types.Address, cfg *config.NetworkConfig) bool {
	st := state.GlobalState
	info := st.Validators[string(addr)]
	return info.Stake >= types.Amount(cfg.MinStake) && info.SlashCount < 3
}

// Reward — начисляет cfg.BlockReward валидатору
func Reward(addr types.Address, blockHeight uint64, cfg *config.NetworkConfig) {
	logging.Logger.Info("Reward: ДО", "address", addr, "blockHeight", blockHeight, "amount", cfg.BlockReward)
	_ = state.ApplyStateChange(func(st *state.State) {
		acc := st.Accounts[string(addr)]
		acc.Balance += types.Amount(cfg.BlockReward)
		st.Accounts[string(addr)] = acc
		logging.Logger.Info("Reward: ПОСЛЕ", "account", acc)
	})
}

// GetValidatorStake — получить стейк валидатора
func GetValidatorStake(addr types.Address) types.Amount {
	st := state.GlobalState
	return st.Validators[string(addr)].Stake
}

// CalcBlockReward — рассчитать награду за блок (можно расширять)
func CalcBlockReward(height uint64) types.Amount {
	return types.Amount(config.GetDefaultNetworkConfig().BlockReward)
}

// LockStakeWithInput — разбирает input (seed/pubkey/address), вызывает LockStake
func LockStakeWithInput(input string, amt uint64, durationSec int64) (types.Address, int64, error) {
	manager := wallet.NewManager(wallet.WalletsFile)
	cfg := config.GetDefaultNetworkConfig()
	var addr types.Address
	if wallet.IsMnemonic(input) {
		entry, err := manager.ImportWallet(input, &cfg)
		if err != nil {
			return "", 0, err
		}
		addr = types.Address(entry.PubKey)
	} else {
		addr = types.Address(input)
	}
	until := time.Now().Unix() + durationSec
	LockStake(addr, amt, until)
	return addr, until, nil
}

// GetStakeInfoByInput — возвращает StakeInfo по input (seed/pubkey/address)
func GetStakeInfoByInput(input string) (types.Address, types.StakeInfo, error) {
	manager := wallet.NewManager(wallet.WalletsFile)
	cfg := config.GetDefaultNetworkConfig()
	var addr types.Address
	if wallet.IsMnemonic(input) {
		entry, err := manager.ImportWallet(input, &cfg)
		if err != nil {
			return "", types.StakeInfo{}, err
		}
		addr = types.Address(entry.PubKey)
	} else {
		addr = types.Address(input)
	}
	info := state.GlobalState.Validators[string(addr)]
	stakeInfo := types.StakeInfo{
		Amount:      uint64(info.Stake),
		LockedUntil: int64(info.LockedUntil),
		SlashCount:  info.SlashCount,
	}
	return addr, stakeInfo, nil
}

// LockStakeWithInputAndSave лочит stake, сохраняет state, логирует
func LockStakeWithInputAndSave(cfg *config.NetworkConfig, input string, amt uint64, durationSec int64) (types.Address, int64, error) {
	addr, until, err := LockStakeWithInput(input, amt, durationSec)
	if err != nil {
		return addr, until, err
	}
	logging.Logger.Info("LockStakeWithInputAndSave", "address", addr, "amount", amt, "until", until)
	return addr, until, nil
}

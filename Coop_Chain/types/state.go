// Package state описывает единственный «зеркальный» слепок сети MCP Chain,
// который держится в RAM и периодически сохраняется в data/state.json.
// Никаких "магических чисел": все пороги-комиссии-награды подаются
// через internal/config и не жёстко прописаны в коде.

package types

import (
	"encoding/json"
	"time"
)

/* ---------- 1. Метаданные цепочки ---------- */

type ChainMeta struct {
	Height        uint64    `json:"height"`          // высота последнего блока
	LastBlockHash Hash      `json:"last_block_hash"` // хэш Header'а
	TxRoot        Hash      `json:"tx_root"`         // Merkle-root всех TX
	StateRoot     Hash      `json:"state_root"`      // SHA-256(state после блока)
	Timestamp     time.Time `json:"timestamp"`       // время майна последнего блока
}

/* ---------- 2. Денежная масса MCP ---------- */

type Supply struct {
	TotalSupply Amount `json:"total_supply"` // сгенерировано со старта сети
	Circulating Amount `json:"circulating"`  // на руках у пользователей
	Burned      Amount `json:"burned"`       // навсегда уничтожено
	FeePool     Amount `json:"fee_pool"`     // накопленные комиссии
}

/* ---------- 3. Счета пользователей ---------- */

type Account struct {
	Balance Amount `json:"balance"` // MCP в наносатоши
	Nonce   uint64 `json:"nonce"`   // защита от replay-атак
}

/* ---------- 4. Валидаторы / стейк ---------- */

type ValidatorInfo struct {
	Stake       Amount `json:"stake"`        // залочено MCP
	LockedUntil uint64 `json:"locked_until"` // высота блока, до которой нельзя снять стейк
	Rating      uint64 `json:"rating"`       // сколько блоков добыл
	SlashCount  uint32 `json:"slash_count"`  // число штрафов
	Active      bool   `json:"active"`       // допускается ли к консенсусу
}

/* ---------- 5. Кооперативы ---------- */

type CoopData struct {
	Name            string          `json:"name"`
	Members         []Address       `json:"members"`
	TreasuryBalance Amount          `json:"treasury_balance"`
	GovernanceRules GovernanceRules `json:"governance_rules"`
}

type GovernanceRules struct {
	Quorum       uint8  `json:"quorum"`        // % голосов
	VoteDuration uint64 `json:"vote_duration"` // в блоках
}

/* ---------- 6. Голосования сети ---------- */

type Proposal struct {
	Type    string          `json:"type"`    // upgrade | coop-internal | text
	Target  string          `json:"target"`  // объект голосования
	Payload json.RawMessage `json:"payload"` // произвольные данные
	Votes   map[string]bool `json:"votes"`   // true=за, false=против
	State   string          `json:"state"`   // open | accepted | rejected
}

/* ---------- 7. Soul-Bound Badges ---------- */

type STBBadge struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	IssuedAt time.Time `json:"issued_at"`
}

type STBSet struct {
	Badges []STBBadge `json:"badges"`
}

/* ---------- 8. Хранилище встроенных контрактов ---------- */

type ContractsStorage map[string]map[string][]byte

//   ContractID            ↳  KV-пары внутри контракта

/* ---------- 9. Казна комиссий ---------- */

type FeeTreasury struct {
	Balance           Amount            `json:"balance"`
	DistributionRules DistributionRules `json:"distribution_rules"`
}

type DistributionRules struct {
	ValidatorShare uint8 `json:"validator_share"` // % комиссии валидатору
	CoopShare      uint8 `json:"coop_share"`      // % кооперативу-создателю блока
}

/* ---------- 10. Временный кэш (не в StateRoot) ---------- */

type TxReceipt struct {
	Status    string      `json:"status"`
	GasUsed   uint64      `json:"gas_used"`
	ReturnVal interface{} `json:"return_val,omitempty"`
}

type RuntimeCache struct {
	RecentBlocks []Hash               `json:"-"`
	TxIndex      map[string]TxReceipt `json:"-"`
}

/* ---------- Итоговый слепок сети ---------- */

type State struct {
	ChainMeta        ChainMeta                `json:"chain_meta"`
	Supply           Supply                   `json:"supply"`
	Accounts         map[string]Account       `json:"accounts"`   // ключ — string (hex address)
	Validators       map[string]ValidatorInfo `json:"validators"` // ключ — string (hex address)
	CoopsRegistry    map[string]CoopData      `json:"coops_registry"`
	Governance       map[string]Proposal      `json:"governance"`
	SoulBound        map[string]STBSet        `json:"soul_bound"` // ключ — string (hex address)
	ContractsStorage ContractsStorage         `json:"contracts_storage"`
	FeeTreasury      FeeTreasury              `json:"fee_treasury"`
	Blocks           []Block                  `json:"blocks"`
	Cache            RuntimeCache             `json:"-"`
}

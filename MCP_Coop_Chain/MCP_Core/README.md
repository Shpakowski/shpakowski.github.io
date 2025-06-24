# MCP Coop Chain (MCP_Core)

A modular, SOLID-compliant blockchain core with clean architecture, strict separation of business logic, and a fully refactored wallet and cooperative management system.

## Project Structure

- **cmd/**: Entry points for running the node/CLI.
- **internal/logger/**: Centralized logging for auditability and debugging.
- **pkg/**
  - **chain/**: Blockchain state, block addition, persistence, and all chain logic.
  - **consensus/**: Consensus state, validator logic, quorum, rating, and penalties.
  - **mempool/**: Transaction pools for standard and smart contract transactions, with all logic modularized.
  - **crypto/**: All cryptographic operations (keygen, signing, hashing, mnemonic) via public interfaces.
  - **wallet/**: User wallet management (creation, import/export, signing, metadata, logging).
  - **types/**: Single source of all core data structures, validation, and serialization.
  - **ecoop/**: Electronic cooperative management (see below).
  - **api/cli/**: Command-line interface and modular commands for node and wallet operations.

## Electronic Cooperative Management (`pkg/ecoop/`)

The `ecoop` package provides all core logic for managing electronic cooperatives (co-ops) on the blockchain. It is fully integrated with the chain and exposes public interfaces for use by CLI, chain, and other modules.

### Main Features
- **Coop struct**: Represents a co-op with metadata, members, assets, rules, and event log.
- **Membership management**: Add/remove members, change roles, list members, all with event recording.
- **Voting and proposals**: Create proposals, cast votes, close proposals, tally results, enforce deadlines/quorum.
- **Asset management**: Create assets, transfer/burn tokens, query balances, enforce supply and rules.
- **Rules**: Configurable rules for joining/leaving, proposal types, voting eligibility, asset transfer limits, role permissions, and quorum.
- **Event log**: All major actions (creation, membership, proposals, asset changes, rule changes) are recorded for auditability and traceability.

### File Overview
- `coop.go`: Core Coop struct, creation, metadata, member/asset access, event log.
- `membership.go`: Member struct and all membership management functions.
- `voting.go`: Proposal and Vote structs, proposal/voting logic, result calculation.
- `assets.go`: Asset struct, asset creation/transfer/burn, balance queries.
- `rules.go`: Rules struct, all validation logic for business rules.
- `events.go`: Event struct, event recording and listing functions.

### Usage Example
```go
import "github.com/mcpcoop/chain/pkg/ecoop"

// Create a new co-op
ecoop := ecoop.NewCoop("MyCoop", "A demo cooperative")

// Add a member
ecoop.AddMember(ecoop, "address1", "founder", map[string]string{"display_name": "Alice"})

// Create an asset
ecoop.CreateAsset(ecoop, "MCP", "Coop Token", 1000000)

// Create a proposal
proposal, _ := ecoop.CreateProposal(ecoop, "address1", "Change Rules", "Update quorum", "governance", []string{"yes","no"}, time.Now().Add(48*time.Hour))

// Cast a vote
ecoop.CastVote(ecoop, proposal.ID, "address1", "yes")

// Transfer asset
ecoop.TransferAsset(ecoop, "MCP", "coop_treasury", "address1", 1000)
```

### Integration
- All co-op state changes are auditable and can be persisted as part of the blockchain state.
- All validation functions return errors if business rules are violated.
- No direct file/database writes in ecoop; persistence is managed at the chain level.
- All logic uses existing types and crypto from the relevant packages.

## Key Architectural Principles

- **Strict modularity**: Each package has a single responsibility and clear boundaries.
- **SOLID compliance**: All logic is separated from data structures and interfaces.
- **No hidden coupling**: All inter-package dependencies go through public interfaces only.
- **Centralized types**: All core types and serialization are in `pkg/types`.
- **Auditability**: Logging and error handling are present throughout the codebase.

## Whitepaper Business Logic Coverage

- **Persistent blockchain**: State is restored from JSON, no data loss on restart.
- **Mempool**: Separate pools for standard and smart contract transactions.
- **Deterministic wallets**: 12-word seed, secure signature/validation, and recovery.
- **Consensus**: Validators, quorum, rating/penalty, and block proposal logic.
- **Periodic block creation**: Implemented via timer/goroutine.
- **Strict package boundaries**: No direct access to storage, mempool, or cryptography outside their packages.
- **Centralized types/serialization**: All types are in `pkg/types`.
- **Logging/auditability**: Logging is present throughout, especially in wallet, ecoop, and mempool logic.

## CLI Usage

```
./mcpnode start
./mcpnode stop
./mcpnode restart
./mcpnode status
./mcpnode new-wallet <12-word-seed>
./mcpnode send <from> <to> <amount>
./mcpnode export-wallets <file>
./mcpnode import-wallets <file>
```

### Example: Create a new wallet from a seed phrase
```
$ ./mcpnode new-wallet "alpha beta gamma delta epsilon zeta eta theta iota kappa lambda mu"
[INFO] New wallet created!
Address: <address>
Initial balance: 100.0
```

### Example: Export wallets to a file
```
$ ./mcpnode export-wallets wallets.json
[INFO] Exporting wallets to file wallets.json
```

### Example: Import wallets from a file
```
$ ./mcpnode import-wallets wallets.json
[INFO] Importing wallets from file wallets.json
```

## Package Responsibilities

- **chain**: Blockchain state, block addition, persistence, and all chain logic.
- **consensus**: Consensus state, validator logic, quorum, rating, and penalties.
- **mempool**: Transaction pools for standard and smart contract transactions, with all logic modularized.
- **crypto**: All cryptographic operations (keygen, signing, hashing, mnemonic) via public interfaces.
- **wallet**: User wallet management (creation, import/export, signing, metadata, logging).
- **ecoop**: Electronic cooperative management (co-ops, members, voting, assets, rules, events).
- **types**: Single source of all core data structures, validation, and serialization.
- **api/cli**: Command-line interface and modular commands for node and wallet operations.

## Testing
- Unit tests for all type validation, wallet, and ecoop logic.
- CLI and integration tests for wallet, co-op, and transaction flows.

## Further Expansion
- Smart contract wallet support
- Hardware wallet integration
- Multi-signature wallets
- Advanced metadata and tagging
- On-chain co-op governance and DAO modules

---

For more details, see the code comments and whitepaper references in each package.

go run ./cmd/mcpnode new-wallet --seed "alpha beta gamma delta epsilon zeta eta theta iota kappa lambda mu" - f84a0be7b985aa32
go run ./cmd/mcpnode new-wallet --seed "alpha beta gamma delta epsilon zeta eta theta iota kappa lambda muka" - 7a3525321bf5a543
go run ./cmd/mcpnode send --from f84a0be7b985aa32 --to 7a3525321bf5a543 --amount 42.0
go run ./cmd/mcpnode status
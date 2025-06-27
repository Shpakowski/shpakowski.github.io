package contracts

// ProtoContractSpec — спецификация встроенного контракта
// Описывает поля, gas-cost, события
// (можно расширять по мере необходимости)
type ProtoContractSpec struct {
	Name    string   // Имя контракта
	Fields  []string // Список полей/аргументов
	GasCost uint64   // Стоимость выполнения (gas)
	Events  []string // События, которые может генерировать контракт
}

// Встроенные контракты протокола
var (
	CreateCoopSpec = ProtoContractSpec{
		Name:    "CreateCoop",
		Fields:  []string{"creator", "name", "description"},
		GasCost: 1000,
		Events:  []string{"CoopCreated"},
	}
	VoteSpec = ProtoContractSpec{
		Name:    "Vote",
		Fields:  []string{"voter", "coop_id", "proposal_id", "choice"},
		GasCost: 500,
		Events:  []string{"Voted"},
	}
	IssueTokenSpec = ProtoContractSpec{
		Name:    "IssueToken",
		Fields:  []string{"issuer", "amount", "symbol"},
		GasCost: 2000,
		Events:  []string{"TokenIssued"},
	}
)

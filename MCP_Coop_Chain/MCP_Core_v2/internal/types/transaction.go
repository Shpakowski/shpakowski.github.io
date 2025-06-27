package types

// ContractCall описывает вызов встроенного Proto API контракта в блоке
// Используется для системных операций (Transfer, Balance и др.)
type ContractCall struct {
	Contract string   // Имя или идентификатор встроенного контракта
	Method   string   // Метод Proto API (например, "Transfer", "Balance")
	Args     []string // Аргументы вызова (все в строковом виде для универсальности)
	Caller   string   // Адрес инициатора вызова
	Result   string   // Результат выполнения (опционально, для истории)
}

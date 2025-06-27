package types

import "time"

// SCTx — структура для вызова смарт-контракта
// (можно расширять по мере необходимости)
type SCTx struct {
	ID        string    // Уникальный идентификатор вызова
	Contract  string    // Имя контракта
	Method    string    // Метод контракта
	Args      []string  // Аргументы
	GasLimit  uint64    // Ограничение по gas
	Priority  int       // Приоритет выполнения
	Timestamp time.Time // Время поступления
}

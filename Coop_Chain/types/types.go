package types

import "errors"

// Address — тип для хранения адреса пользователя или контракта
// Обычно это строка в hex-формате
// Пример: "0x1234abcd..."
type Address string

// Hash — тип для хранения хэша блока или транзакции
// Пример: "0xabcdef1234..."
type Hash string

// Amount — тип для хранения суммы MCP (можно использовать float64 или int64)
type Amount float64

// Timestamp — тип для хранения времени (unix-таймштамп в секундах)
type Timestamp int64

// StakeInfo — информация о стейке валидатора
// Amount — залоченный стейк, LockedUntil — unix-время, SlashCount — число штрафов
// Хранится в state.ValidatorsInfo
type StakeInfo struct {
	Amount      uint64
	LockedUntil int64
	SlashCount  uint32
}

// Интерфейс для сериализации/десериализации (можно расширять по мере необходимости)
type Serializable interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

// NewBlockError создаёт ошибку блока с заданным сообщением
func NewBlockError(msg string) error {
	return errors.New(msg)
}

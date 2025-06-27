package types

// Transaction — структура транзакции
// Все поля сериализуются в JSON в стабильном порядке
type Transaction struct {
	ID     Hash    `json:"id"`     // Уникальный идентификатор транзакции (хэш)
	From   Address `json:"from"`   // Отправитель
	To     Address `json:"to"`     // Получатель
	Amount Amount  `json:"amount"` // Сумма перевода
	Fee    Amount  `json:"fee"`    // Комиссия
	Nonce  uint64  `json:"nonce"`  // Нонc для защиты от повторов
	Sig    []byte  `json:"sig"`    // Подпись отправителя (Ed25519)
}

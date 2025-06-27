package types

// Header — заголовок блока
// Содержит метаданные блока
// (можно расширять по мере необходимости)
type Header struct {
	Height    uint64    // Высота блока
	PrevHash  Hash      // Хэш предыдущего блока
	Timestamp Timestamp // Время создания блока
	Proposer  Address   // Адрес валидатора, создавшего блок
}

// Body — тело блока (список транзакций)
type Body struct {
	Txs          []Hash        // Хэши транзакций
	Transactions []Transaction // Сами транзакции
}

// Block — структура блока
// Содержит заголовок, тело и хэш
// Хэш вычисляется по всему блоку (Header+Body)
type Block struct {
	Header Header
	Body   Body
	Hash   Hash
}

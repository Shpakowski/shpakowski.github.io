package db

// Интерфейс для работы с базой данных (заглушка для Prestige DB)
type DB interface {
	Open(path string) error
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	NewTxn() Txn
}

// Txn — интерфейс для транзакций БД (заглушка)
type Txn interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Commit() error
	Discard()
}

// InMemoryDB — простая in-memory реализация для MVP
// (используется до подключения реальной Prestige DB)
type InMemoryDB struct {
	store map[string][]byte
}

// Open инициализирует in-memory хранилище
func (db *InMemoryDB) Open(path string) error {
	db.store = make(map[string][]byte)
	return nil
}

// Put сохраняет значение по ключу
func (db *InMemoryDB) Put(key, value []byte) error {
	db.store[string(key)] = value
	return nil
}

// Get возвращает значение по ключу
func (db *InMemoryDB) Get(key []byte) ([]byte, error) {
	v, ok := db.store[string(key)]
	if !ok {
		return nil, nil // nil если не найдено
	}
	return v, nil
}

// NewTxn возвращает фиктивную транзакцию
func (db *InMemoryDB) NewTxn() Txn {
	return &InMemoryTxn{db: db}
}

// InMemoryTxn — простая транзакция для in-memory БД
// (в реальной реализации будет атомарность)
type InMemoryTxn struct {
	db *InMemoryDB
}

func (txn *InMemoryTxn) Put(key, value []byte) error {
	return txn.db.Put(key, value)
}

func (txn *InMemoryTxn) Get(key []byte) ([]byte, error) {
	return txn.db.Get(key)
}

func (txn *InMemoryTxn) Commit() error {
	return nil // В in-memory ничего не делаем
}

func (txn *InMemoryTxn) Discard() {
	// В in-memory ничего не делаем
}

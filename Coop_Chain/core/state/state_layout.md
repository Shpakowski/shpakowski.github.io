# Паспорт структуры дерева состояния (state tree) для MCP Chain

Дерево состояния MCP Chain — это основной слепок всей сети, который хранится в оперативной памяти и периодически сериализуется на диск. Все параметры (награды, комиссии, лимиты, интервалы) берутся только из internal/config/network_config.go — никаких жёстко прописанных чисел или констант. Все базовые типы (Address, Hash, Amount и др.) берутся из types/state.go для единообразия во всём проекте.

## Секции дерева состояния

1. **ChainMeta** — метаданные цепочки:
   - Высота последнего блока, хэш последнего блока, Merkle-root всех транзакций, StateRoot, время майнинга последнего блока.
   - Нужна для отслеживания текущего состояния цепи и быстрой валидации новых блоков.

2. **Supply** — денежная масса MCP:
   - Общий объём сгенерированных монет, количество монет в обращении, сожжённые монеты, накопленные комиссии.
   - Позволяет отслеживать экономику сети и балансировать эмиссию.

3. **Accounts** — счета пользователей:
   - Баланс пользователя и nonce для защиты от replay-атак.
   - Основная таблица для учёта средств и предотвращения двойных трат.

4. **Validators** — валидаторы и их стейк:
   - Информация о стейке, блокировке, рейтинге, штрафах и активности валидаторов.
   - Используется для консенсуса и распределения наград.

5. **CoopsRegistry** — кооперативы:
   - Данные о кооперативах: название, члены, баланс казны, правила управления.
   - Для реализации децентрализованных организаций и их казны.

6. **Governance** — голосования сети:
   - Активные и завершённые предложения, голоса, состояние голосования.
   - Механизм управления параметрами сети и кооперативов.

7. **SoulBound** — Soul-Bound Badges:
   - Наборы невзаимозаменяемых бейджей, привязанных к адресам.
   - Для репутации, достижений и уникальных прав.

8. **ContractsStorage** — хранилище встроенных контрактов:
   - Ключ-значение для каждого контракта.
   - Для хранения состояния смарт-контрактов.

9. **FeeTreasury** — казна комиссий:
   - Баланс казны и правила распределения комиссий между валидаторами и кооперативами.
   - Для прозрачного распределения доходов сети.

10. **RuntimeCache** — временный кэш (не влияет на StateRoot):
   - Кэш недавних блоков и индексы транзакций.
   - Для ускорения работы узла, не сериализуется в StateRoot.

---

Все поля state оформлены по единому стилю, типы берутся из types/state.go. Если какого-то типа не хватает — его нужно добавить в types/state.go, а не в core/state/state.go.

# State Layout (RAM/JSON)

## Основные компоненты state
- **Accounts**: Балансы и nonce всех кошельков
- **Validators**: Информация о валидаторах
- **Blocks**: Список всех блоков (с транзакциями)
- **CoopsRegistry, Governance, SoulBound, ContractsStorage, FeeTreasury**: служебные структуры

## Важно
- **Mempool** не входит в state.json, это отдельная RAM-очередь (`GlobalTxMempool`).
- **Транзакции применяются к state только после майнинга блока**.
- **state.json** обновляется только после применения блока (балансы, nonce и т.д.).

## Жизненный цикл транзакции
1. Транзакция отправляется через API/CLI и попадает в mempool.
2. Node майнит блок из транзакций mempool.
3. Транзакции из блока применяются к state.
4. State сохраняется в state.json.

## Пример структуры state.json
```json
{
  "Accounts": { ... },
  "Validators": { ... },
  "Blocks": [ ... ],
  ...
}
``` 
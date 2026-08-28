# MercadoBitcoin Golang SDK Reference

Complete API reference for the MercadoBitcoin Golang SDK.


## MercadoBitcoinSDK

### Constructor

```go
func NewMercadoBitcoinSDK(options map[string]any) *MercadoBitcoinSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *MercadoBitcoinSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *MercadoBitcoinSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Balance(data map[string]any) MercadoBitcoinEntity`

Create a new `Balance` entity instance. Pass `nil` for no initial data.

#### `Candle(data map[string]any) MercadoBitcoinEntity`

Create a new `Candle` entity instance. Pass `nil` for no initial data.

#### `DepositAddress(data map[string]any) MercadoBitcoinEntity`

Create a new `DepositAddress` entity instance. Pass `nil` for no initial data.

#### `Order(data map[string]any) MercadoBitcoinEntity`

Create a new `Order` entity instance. Pass `nil` for no initial data.

#### `OrderBook(data map[string]any) MercadoBitcoinEntity`

Create a new `OrderBook` entity instance. Pass `nil` for no initial data.

#### `Ticker(data map[string]any) MercadoBitcoinEntity`

Create a new `Ticker` entity instance. Pass `nil` for no initial data.

#### `Trade(data map[string]any) MercadoBitcoinEntity`

Create a new `Trade` entity instance. Pass `nil` for no initial data.

#### `Withdrawal(data map[string]any) MercadoBitcoinEntity`

Create a new `Withdrawal` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## BalanceEntity

```go
balance := client.Balance(nil)
fmt.Println(balance.GetName()) // "balance"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | `float64` | No | Available balance |
| `currency` | `string` | No | Currency code |
| `locked` | `float64` | No | Locked balance |
| `total` | `float64` | No | Total balance |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Balance(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CandleEntity

```go
candle := client.Candle(nil)
fmt.Println(candle.GetName()) // "candle"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `close` | `float64` | No | Closing price |
| `high` | `float64` | No | Highest price |
| `id` | `string` | No |  |
| `low` | `float64` | No | Lowest price |
| `open` | `float64` | No | Opening price |
| `timestamp` | `int` | No | Candle timestamp in milliseconds |
| `volume` | `float64` | No | Trading volume |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Candle(nil).Load(map[string]any{"id": "candle_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CandleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DepositAddressEntity

```go
depositAddress := client.DepositAddress(nil)
fmt.Println(depositAddress.GetName()) // "deposit_address"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `string` | No | Deposit address |
| `currency` | `string` | No | Cryptocurrency code |
| `qrCode` | `string` | No | QR code for deposit address |
| `tag` | `string` | No | Deposit tag/memo (if applicable) |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.DepositAddress(nil).Load(map[string]any{"currency": "currency"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DepositAddressEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrderEntity

```go
order := client.Order(nil)
fmt.Println(order.GetName()) // "order"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No | Order amount |
| `filled` | `float64` | No | Filled amount |
| `id` | `string` | No | Order ID |
| `price` | `float64` | No | Order price |
| `side` | `string` | No | Order side |
| `status` | `string` | No | Order status |
| `symbol` | `string` | No | Trading pair symbol |
| `timestamp` | `int` | No | Order creation timestamp |
| `type` | `string` | No | Order type |

### Field Usage by Operation

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `amount` | - | - | Yes | - |
| `filled` | - | - | - | - |
| `id` | - | - | - | - |
| `price` | - | - | - | - |
| `side` | - | - | Yes | - |
| `status` | - | - | - | - |
| `symbol` | - | - | Yes | - |
| `timestamp` | - | - | - | - |
| `type` | - | - | Yes | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Order(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Order(nil).Load(map[string]any{"id": "order_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Order(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Order(nil).Remove(map[string]any{"id": "order_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrderEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrderBookEntity

```go
orderBook := client.OrderBook(nil)
fmt.Println(orderBook.GetName()) // "order_book"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `asks` | `[]any` | No | List of ask orders |
| `bids` | `[]any` | No | List of bid orders |
| `timestamp` | `int` | No | Timestamp in milliseconds |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.OrderBook(nil).Load(map[string]any{"symbol": "symbol"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrderBookEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TickerEntity

```go
ticker := client.Ticker(nil)
fmt.Println(ticker.GetName()) // "ticker"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | `float64` | No | Lowest ask price |
| `bid` | `float64` | No | Highest bid price |
| `high` | `float64` | No | 24h high price |
| `id` | `string` | No |  |
| `last` | `float64` | No | Last traded price |
| `low` | `float64` | No | 24h low price |
| `symbol` | `string` | No | Trading pair symbol |
| `timestamp` | `int` | No | Timestamp in milliseconds |
| `volume` | `float64` | No | 24h trading volume |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Ticker(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Ticker(nil).Load(map[string]any{"id": "ticker_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TickerEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TradeEntity

```go
trade := client.Trade(nil)
fmt.Println(trade.GetName()) // "trade"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No | Trade amount |
| `id` | `string` | No | Trade ID |
| `price` | `float64` | No | Trade price |
| `side` | `string` | No | Trade side |
| `timestamp` | `int` | No | Timestamp in milliseconds |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Trade(nil).Load(map[string]any{"id": "trade_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TradeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WithdrawalEntity

```go
withdrawal := client.Withdrawal(nil)
fmt.Println(withdrawal.GetName()) // "withdrawal"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountNumber` | `string` | Yes | Bank account number |
| `accountType` | `string` | No | Account type |
| `address` | `string` | Yes | Destination address |
| `agency` | `string` | Yes | Bank agency |
| `amount` | `float64` | Yes | Withdrawal amount in BRL |
| `bank` | `string` | Yes | Bank code |
| `currency` | `string` | Yes | Cryptocurrency code |
| `tag` | `string` | No | Destination tag/memo (if applicable) |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Withdrawal(nil).Create(map[string]any{
    "accountNumber": "example_accountNumber",
    "address": "example_address",
    "agency": "example_agency",
    "amount": 1,
    "bank": "example_bank",
    "currency": "example_currency",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WithdrawalEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewMercadoBitcoinSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.


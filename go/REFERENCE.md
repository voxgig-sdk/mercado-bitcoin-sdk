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
| `available` | `float64` | No |  |
| `currency` | `string` | No |  |
| `locked` | `float64` | No |  |
| `total` | `float64` | No |  |

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
| `close` | `float64` | No |  |
| `high` | `float64` | No |  |
| `low` | `float64` | No |  |
| `open` | `float64` | No |  |
| `timestamp` | `int` | No |  |
| `volume` | `float64` | No |  |

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
| `address` | `string` | No |  |
| `currency` | `string` | No |  |
| `qr_code` | `string` | No |  |
| `tag` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.DepositAddress(nil).Load(nil, nil)
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
| `amount` | `float64` | No |  |
| `filled` | `float64` | No |  |
| `id` | `string` | No |  |
| `price` | `float64` | No |  |
| `side` | `string` | No |  |
| `status` | `string` | No |  |
| `symbol` | `string` | No |  |
| `timestamp` | `int` | No |  |
| `type` | `string` | No |  |

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
| `ask` | `[]any` | No |  |
| `bid` | `[]any` | No |  |
| `timestamp` | `int` | No |  |

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
| `ask` | `float64` | No |  |
| `bid` | `float64` | No |  |
| `high` | `float64` | No |  |
| `last` | `float64` | No |  |
| `low` | `float64` | No |  |
| `symbol` | `string` | No |  |
| `timestamp` | `int` | No |  |
| `volume` | `float64` | No |  |

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
| `amount` | `float64` | No |  |
| `id` | `string` | No |  |
| `price` | `float64` | No |  |
| `side` | `string` | No |  |
| `timestamp` | `int` | No |  |

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
| `account_number` | `string` | Yes |  |
| `account_type` | `string` | No |  |
| `address` | `string` | Yes |  |
| `agency` | `string` | Yes |  |
| `amount` | `float64` | Yes |  |
| `bank` | `string` | Yes |  |
| `currency` | `string` | Yes |  |
| `tag` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Withdrawal(nil).Create(map[string]any{
    "account_number": "example_account_number",
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


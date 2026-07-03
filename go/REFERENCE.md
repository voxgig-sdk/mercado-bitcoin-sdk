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
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | ``$NUMBER`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `locked` | ``$NUMBER`` | No |  |
| `total` | ``$NUMBER`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Balance(nil).List(nil, nil)
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
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `close` | ``$NUMBER`` | No |  |
| `high` | ``$NUMBER`` | No |  |
| `low` | ``$NUMBER`` | No |  |
| `open` | ``$NUMBER`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |
| `volume` | ``$NUMBER`` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Candle(nil).Load(map[string]any{"id": "candle_id"}, nil)
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
deposit_address := client.DepositAddress(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `qr_code` | ``$STRING`` | No |  |
| `tag` | ``$STRING`` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.DepositAddress(nil).Load(map[string]any{"id": "deposit_address_id"}, nil)
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
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | ``$NUMBER`` | No |  |
| `filled` | ``$NUMBER`` | No |  |
| `id` | ``$STRING`` | No |  |
| `price` | ``$NUMBER`` | No |  |
| `side` | ``$STRING`` | No |  |
| `status` | ``$STRING`` | No |  |
| `symbol` | ``$STRING`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |
| `type` | ``$STRING`` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `amount` | - | - | Yes | - | - |
| `filled` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `price` | - | - | - | - | - |
| `side` | - | - | Yes | - | - |
| `status` | - | - | - | - | - |
| `symbol` | - | - | Yes | - | - |
| `timestamp` | - | - | - | - | - |
| `type` | - | - | Yes | - | - |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Order(nil).Create(map[string]any{
}, nil)
```

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Order(nil).List(nil, nil)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Order(nil).Load(map[string]any{"id": "order_id"}, nil)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Order(nil).Remove(map[string]any{"id": "order_id"}, nil)
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
order_book := client.OrderBook(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | ``$ARRAY`` | No |  |
| `bid` | ``$ARRAY`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.OrderBook(nil).Load(map[string]any{"id": "order_book_id"}, nil)
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
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | ``$NUMBER`` | No |  |
| `bid` | ``$NUMBER`` | No |  |
| `high` | ``$NUMBER`` | No |  |
| `last` | ``$NUMBER`` | No |  |
| `low` | ``$NUMBER`` | No |  |
| `symbol` | ``$STRING`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |
| `volume` | ``$NUMBER`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Ticker(nil).List(nil, nil)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Ticker(nil).Load(map[string]any{"id": "ticker_id"}, nil)
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
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | ``$NUMBER`` | No |  |
| `id` | ``$STRING`` | No |  |
| `price` | ``$NUMBER`` | No |  |
| `side` | ``$STRING`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Trade(nil).Load(map[string]any{"id": "trade_id"}, nil)
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
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number` | ``$STRING`` | Yes |  |
| `account_type` | ``$STRING`` | No |  |
| `address` | ``$STRING`` | Yes |  |
| `agency` | ``$STRING`` | Yes |  |
| `amount` | ``$NUMBER`` | Yes |  |
| `bank` | ``$STRING`` | Yes |  |
| `currency` | ``$STRING`` | Yes |  |
| `tag` | ``$STRING`` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Withdrawal(nil).Create(map[string]any{
    "account_number": /* `$STRING` */,
    "address": /* `$STRING` */,
    "agency": /* `$STRING` */,
    "amount": /* `$NUMBER` */,
    "bank": /* `$STRING` */,
    "currency": /* `$STRING` */,
}, nil)
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


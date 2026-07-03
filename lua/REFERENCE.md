# MercadoBitcoin Lua SDK Reference

Complete API reference for the MercadoBitcoin Lua SDK.


## MercadoBitcoinSDK

### Constructor

```lua
local sdk = require("mercado-bitcoin_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Balance(data)`

Create a new `Balance` entity instance. Pass `nil` for no initial data.

#### `Candle(data)`

Create a new `Candle` entity instance. Pass `nil` for no initial data.

#### `DepositAddress(data)`

Create a new `DepositAddress` entity instance. Pass `nil` for no initial data.

#### `Order(data)`

Create a new `Order` entity instance. Pass `nil` for no initial data.

#### `OrderBook(data)`

Create a new `OrderBook` entity instance. Pass `nil` for no initial data.

#### `Ticker(data)`

Create a new `Ticker` entity instance. Pass `nil` for no initial data.

#### `Trade(data)`

Create a new `Trade` entity instance. Pass `nil` for no initial data.

#### `Withdrawal(data)`

Create a new `Withdrawal` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## BalanceEntity

```lua
local balance = client:Balance(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | ``$NUMBER`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `locked` | ``$NUMBER`` | No |  |
| `total` | ``$NUMBER`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Balance():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CandleEntity

```lua
local candle = client:Candle(nil)
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

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Candle():load({ id = "candle_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CandleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DepositAddressEntity

```lua
local deposit_address = client:DepositAddress(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `qr_code` | ``$STRING`` | No |  |
| `tag` | ``$STRING`` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DepositAddress():load({ id = "deposit_address_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DepositAddressEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrderEntity

```lua
local order = client:Order(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Order():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Order():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Order():load({ id = "order_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Order():remove({ id = "order_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrderEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrderBookEntity

```lua
local order_book = client:OrderBook(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | ``$ARRAY`` | No |  |
| `bid` | ``$ARRAY`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:OrderBook():load({ id = "order_book_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrderBookEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TickerEntity

```lua
local ticker = client:Ticker(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Ticker():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Ticker():load({ id = "ticker_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TickerEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TradeEntity

```lua
local trade = client:Trade(nil)
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

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Trade():load({ id = "trade_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TradeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WithdrawalEntity

```lua
local withdrawal = client:Withdrawal(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Withdrawal():create({
  account_number = --[[ `$STRING` ]],
  address = --[[ `$STRING` ]],
  agency = --[[ `$STRING` ]],
  amount = --[[ `$NUMBER` ]],
  bank = --[[ `$STRING` ]],
  currency = --[[ `$STRING` ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WithdrawalEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```


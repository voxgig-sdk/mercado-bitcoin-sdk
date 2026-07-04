# MercadoBitcoin Ruby SDK Reference

Complete API reference for the MercadoBitcoin Ruby SDK.


## MercadoBitcoinSDK

### Constructor

```ruby
require_relative 'mercado-bitcoin_sdk'

client = MercadoBitcoinSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `MercadoBitcoinSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = MercadoBitcoinSDK.test
```


### Instance Methods

#### `Balance(data = nil)`

Create a new `Balance` entity instance. Pass `nil` for no initial data.

#### `Candle(data = nil)`

Create a new `Candle` entity instance. Pass `nil` for no initial data.

#### `DepositAddress(data = nil)`

Create a new `DepositAddress` entity instance. Pass `nil` for no initial data.

#### `Order(data = nil)`

Create a new `Order` entity instance. Pass `nil` for no initial data.

#### `OrderBook(data = nil)`

Create a new `OrderBook` entity instance. Pass `nil` for no initial data.

#### `Ticker(data = nil)`

Create a new `Ticker` entity instance. Pass `nil` for no initial data.

#### `Trade(data = nil)`

Create a new `Trade` entity instance. Pass `nil` for no initial data.

#### `Withdrawal(data = nil)`

Create a new `Withdrawal` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## BalanceEntity

```ruby
balance = client.Balance
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | ``$NUMBER`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `locked` | ``$NUMBER`` | No |  |
| `total` | ``$NUMBER`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.Balance.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CandleEntity

```ruby
candle = client.Candle
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

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Candle.load({ "id" => "candle_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CandleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DepositAddressEntity

```ruby
deposit_address = client.DepositAddress
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `qr_code` | ``$STRING`` | No |  |
| `tag` | ``$STRING`` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.DepositAddress.load({ "id" => "deposit_address_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DepositAddressEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## OrderEntity

```ruby
order = client.Order
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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Order.create({
})
```

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.Order.list(nil)
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Order.load({ "id" => "order_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Order.remove({ "id" => "order_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `OrderEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## OrderBookEntity

```ruby
order_book = client.OrderBook
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | ``$ARRAY`` | No |  |
| `bid` | ``$ARRAY`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.OrderBook.load({ "id" => "order_book_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `OrderBookEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TickerEntity

```ruby
ticker = client.Ticker
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

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.Ticker.list(nil)
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Ticker.load({ "id" => "ticker_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TickerEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TradeEntity

```ruby
trade = client.Trade
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

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Trade.load({ "id" => "trade_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TradeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## WithdrawalEntity

```ruby
withdrawal = client.Withdrawal
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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Withdrawal.create({
  "account_number" => # `$STRING`,
  "address" => # `$STRING`,
  "agency" => # `$STRING`,
  "amount" => # `$NUMBER`,
  "bank" => # `$STRING`,
  "currency" => # `$STRING`,
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `WithdrawalEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = MercadoBitcoinSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


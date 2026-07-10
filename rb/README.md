# MercadoBitcoin Ruby SDK



The Ruby SDK for the MercadoBitcoin API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Balance` — with named operations (`list`/`load`/`create`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases](https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "MercadoBitcoin_sdk"

client = MercadoBitcoinSDK.new({
  "apikey" => ENV["MERCADO_BITCOIN_APIKEY"],
})
```

### 2. List balance records

```ruby
begin
  # list returns an Array of Balance records — iterate directly.
  balances = client.Balance.list
  balances.each do |item|
    puts "#{item["available"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load an orderbook

OrderBook is nested under symbol, so provide the `symbol`.

```ruby
begin
  # load returns the bare OrderBook record (raises on error).
  orderbook = client.OrderBook.load({ "symbol" => "example_symbol" })
  puts orderbook
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  balances = client.Balance.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = MercadoBitcoinSDK.test

# Entity ops return the bare mock record (raises on error).
balance = client.Balance.list()
puts balance
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = MercadoBitcoinSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
MERCADO_BITCOIN_TEST_LIVE=TRUE
MERCADO_BITCOIN_APIKEY=<your-key>
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### MercadoBitcoinSDK

```ruby
require_relative "MercadoBitcoin_sdk"
client = MercadoBitcoinSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = MercadoBitcoinSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### MercadoBitcoinSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Balance` | `(data) -> BalanceEntity` | Create a Balance entity instance. |
| `Candle` | `(data) -> CandleEntity` | Create a Candle entity instance. |
| `DepositAddress` | `(data) -> DepositAddressEntity` | Create a DepositAddress entity instance. |
| `Order` | `(data) -> OrderEntity` | Create an Order entity instance. |
| `OrderBook` | `(data) -> OrderBookEntity` | Create an OrderBook entity instance. |
| `Ticker` | `(data) -> TickerEntity` | Create a Ticker entity instance. |
| `Trade` | `(data) -> TradeEntity` | Create a Trade entity instance. |
| `Withdrawal` | `(data) -> WithdrawalEntity` | Create a Withdrawal entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `MercadoBitcoinError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Balance

| Field | Description |
| --- | --- |
| `available` |  |
| `currency` |  |
| `locked` |  |
| `total` |  |

Operations: List.

API path: `/accounts/balance`

#### Candle

| Field | Description |
| --- | --- |
| `close` |  |
| `high` |  |
| `low` |  |
| `open` |  |
| `timestamp` |  |
| `volume` |  |

Operations: Load.

API path: `/candles/{symbol}`

#### DepositAddress

| Field | Description |
| --- | --- |
| `address` |  |
| `currency` |  |
| `qr_code` |  |
| `tag` |  |

Operations: Load.

API path: `/deposits/crypto`

#### Order

| Field | Description |
| --- | --- |
| `amount` |  |
| `filled` |  |
| `id` |  |
| `price` |  |
| `side` |  |
| `status` |  |
| `symbol` |  |
| `timestamp` |  |
| `type` |  |

Operations: Create, List, Load, Remove.

API path: `/orders`

#### OrderBook

| Field | Description |
| --- | --- |
| `ask` |  |
| `bid` |  |
| `timestamp` |  |

Operations: Load.

API path: `/orderbook/{symbol}`

#### Ticker

| Field | Description |
| --- | --- |
| `ask` |  |
| `bid` |  |
| `high` |  |
| `last` |  |
| `low` |  |
| `symbol` |  |
| `timestamp` |  |
| `volume` |  |

Operations: List, Load.

API path: `/tickers`

#### Trade

| Field | Description |
| --- | --- |
| `amount` |  |
| `id` |  |
| `price` |  |
| `side` |  |
| `timestamp` |  |

Operations: Load.

API path: `/trades/{symbol}`

#### Withdrawal

| Field | Description |
| --- | --- |
| `account_number` |  |
| `account_type` |  |
| `address` |  |
| `agency` |  |
| `amount` |  |
| `bank` |  |
| `currency` |  |
| `tag` |  |

Operations: Create.

API path: `/withdrawals/brl`



## Entities


### Balance

Create an instance: `balance = client.Balance`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `available` | `Float` |  |
| `currency` | `String` |  |
| `locked` | `Float` |  |
| `total` | `Float` |  |

#### Example: List

```ruby
# list returns an Array of Balance records (raises on error).
balances = client.Balance.list
```


### Candle

Create an instance: `candle = client.Candle`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `close` | `Float` |  |
| `high` | `Float` |  |
| `low` | `Float` |  |
| `open` | `Float` |  |
| `timestamp` | `Integer` |  |
| `volume` | `Float` |  |

#### Example: Load

```ruby
# load returns the bare Candle record (raises on error).
candle = client.Candle.load({ "id" => "candle_id" })
```


### DepositAddress

Create an instance: `deposit_address = client.DepositAddress`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `String` |  |
| `currency` | `String` |  |
| `qr_code` | `String` |  |
| `tag` | `String` |  |

#### Example: Load

```ruby
# load returns the bare DepositAddress record (raises on error).
deposit_address = client.DepositAddress.load()
```


### Order

Create an instance: `order = client.Order`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `filled` | `Float` |  |
| `id` | `String` |  |
| `price` | `Float` |  |
| `side` | `String` |  |
| `status` | `String` |  |
| `symbol` | `String` |  |
| `timestamp` | `Integer` |  |
| `type` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Order record (raises on error).
order = client.Order.load({ "id" => "order_id" })
```

#### Example: List

```ruby
# list returns an Array of Order records (raises on error).
orders = client.Order.list
```

#### Example: Create

```ruby
order = client.Order.create({
})
```


### OrderBook

Create an instance: `order_book = client.OrderBook`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | `Array` |  |
| `bid` | `Array` |  |
| `timestamp` | `Integer` |  |

#### Example: Load

```ruby
# load returns the bare OrderBook record (raises on error).
order_book = client.OrderBook.load({ "symbol" => "symbol" })
```


### Ticker

Create an instance: `ticker = client.Ticker`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | `Float` |  |
| `bid` | `Float` |  |
| `high` | `Float` |  |
| `last` | `Float` |  |
| `low` | `Float` |  |
| `symbol` | `String` |  |
| `timestamp` | `Integer` |  |
| `volume` | `Float` |  |

#### Example: Load

```ruby
# load returns the bare Ticker record (raises on error).
ticker = client.Ticker.load({ "id" => "ticker_id" })
```

#### Example: List

```ruby
# list returns an Array of Ticker records (raises on error).
tickers = client.Ticker.list
```


### Trade

Create an instance: `trade = client.Trade`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `id` | `String` |  |
| `price` | `Float` |  |
| `side` | `String` |  |
| `timestamp` | `Integer` |  |

#### Example: Load

```ruby
# load returns the bare Trade record (raises on error).
trade = client.Trade.load({ "id" => "trade_id" })
```


### Withdrawal

Create an instance: `withdrawal = client.Withdrawal`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `String` |  |
| `account_type` | `String` |  |
| `address` | `String` |  |
| `agency` | `String` |  |
| `amount` | `Float` |  |
| `bank` | `String` |  |
| `currency` | `String` |  |
| `tag` | `String` |  |

#### Example: Create

```ruby
withdrawal = client.Withdrawal.create({
  "account_number" => "example_account_number", # String
  "address" => "example_address", # String
  "agency" => "example_agency", # String
  "amount" => 1, # Float
  "bank" => "example_bank", # String
  "currency" => "example_currency", # String
})
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── MercadoBitcoin_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`MercadoBitcoin_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
balance = client.Balance
balance.list()

# balance.data_get now returns the balance data from the last list
# balance.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.

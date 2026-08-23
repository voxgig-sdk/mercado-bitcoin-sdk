# MercadoBitcoin Lua SDK



The Lua SDK for the MercadoBitcoin API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Balance()` — each with the same small set of operations (`list`, `load`, `create`, `remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("mercado-bitcoin_sdk")

local client = sdk.new({
  apikey = os.getenv("MERCADO_BITCOIN_APIKEY"),
})
```

### 2. List balance records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local balances, err = client:Balance():list()
if err then error(err) end

for _, item in ipairs(balances) do
  print(item["currency"])
end
```

### 3. Load an orderbook

OrderBook is nested under symbol, so provide the `symbol`.

```lua
local orderbook, err = client:OrderBook():load({ symbol = "example_symbol" })
if err then error(err) end
print(orderbook)
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local trade, err = client:Trade():load({ id = "example_id" })
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Trade():load({ id = "test01" })
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### MercadoBitcoinSDK

```lua
local sdk = require("mercado-bitcoin_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### MercadoBitcoinSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local candle, err = client:Candle():load({ id = "example_id" })
    if err then error(err) end
    -- candle is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### Balance

| Field | Description |
| --- | --- |
| `available` | Available balance |
| `currency` | Currency code |
| `locked` | Locked balance |
| `total` | Total balance |

Operations: List.

API path: `/accounts/balance`

#### Candle

| Field | Description |
| --- | --- |
| `close` | Closing price |
| `high` | Highest price |
| `low` | Lowest price |
| `open` | Opening price |
| `timestamp` | Candle timestamp in milliseconds |
| `volume` | Trading volume |

Operations: Load.

API path: `/candles/{symbol}`

#### DepositAddress

| Field | Description |
| --- | --- |
| `address` | Deposit address |
| `currency` | Cryptocurrency code |
| `qrCode` | QR code for deposit address |
| `tag` | Deposit tag/memo (if applicable) |

Operations: Load.

API path: `/deposits/crypto`

#### Order

| Field | Description |
| --- | --- |
| `amount` | Order amount |
| `filled` | Filled amount |
| `id` | Order ID |
| `price` | Order price |
| `side` | Order side |
| `status` | Order status |
| `symbol` | Trading pair symbol |
| `timestamp` | Order creation timestamp |
| `type` | Order type |

Operations: Create, List, Load, Remove.

API path: `/orders`

#### OrderBook

| Field | Description |
| --- | --- |
| `asks` | List of ask orders |
| `bids` | List of bid orders |
| `timestamp` | Timestamp in milliseconds |

Operations: Load.

API path: `/orderbook/{symbol}`

#### Ticker

| Field | Description |
| --- | --- |
| `ask` | Lowest ask price |
| `bid` | Highest bid price |
| `high` | 24h high price |
| `last` | Last traded price |
| `low` | 24h low price |
| `symbol` | Trading pair symbol |
| `timestamp` | Timestamp in milliseconds |
| `volume` | 24h trading volume |

Operations: List, Load.

API path: `/tickers`

#### Trade

| Field | Description |
| --- | --- |
| `amount` | Trade amount |
| `id` | Trade ID |
| `price` | Trade price |
| `side` | Trade side |
| `timestamp` | Timestamp in milliseconds |

Operations: Load.

API path: `/trades/{symbol}`

#### Withdrawal

| Field | Description |
| --- | --- |
| `accountNumber` | Bank account number |
| `accountType` | Account type |
| `address` | Destination address |
| `agency` | Bank agency |
| `amount` | Withdrawal amount in BRL |
| `bank` | Bank code |
| `currency` | Cryptocurrency code |
| `tag` | Destination tag/memo (if applicable) |

Operations: Create.

API path: `/withdrawals/brl`



## Entities


### Balance

Create an instance: `local balance = client:Balance(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `available` | `number` | Available balance |
| `currency` | `string` | Currency code |
| `locked` | `number` | Locked balance |
| `total` | `number` | Total balance |

#### Example: List

```lua
local balances, err = client:Balance():list()
```


### Candle

Create an instance: `local candle = client:Candle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `close` | `number` | Closing price |
| `high` | `number` | Highest price |
| `low` | `number` | Lowest price |
| `open` | `number` | Opening price |
| `timestamp` | `number` | Candle timestamp in milliseconds |
| `volume` | `number` | Trading volume |

#### Example: Load

```lua
local candle, err = client:Candle():load({ id = "candle_id" })
```


### DepositAddress

Create an instance: `local deposit_address = client:DepositAddress(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `string` | Deposit address |
| `currency` | `string` | Cryptocurrency code |
| `qrCode` | `string` | QR code for deposit address |
| `tag` | `string` | Deposit tag/memo (if applicable) |

#### Example: Load

```lua
local deposit_address, err = client:DepositAddress():load()
```


### Order

Create an instance: `local order = client:Order(nil)`

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
| `amount` | `number` | Order amount |
| `filled` | `number` | Filled amount |
| `id` | `string` | Order ID |
| `price` | `number` | Order price |
| `side` | `string` | Order side |
| `status` | `string` | Order status |
| `symbol` | `string` | Trading pair symbol |
| `timestamp` | `number` | Order creation timestamp |
| `type` | `string` | Order type |

#### Example: Load

```lua
local order, err = client:Order():load({ id = "order_id" })
```

#### Example: List

```lua
local orders, err = client:Order():list()
```

#### Example: Create

```lua
local order, err = client:Order():create({
})
```


### OrderBook

Create an instance: `local order_book = client:OrderBook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `asks` | `table` | List of ask orders |
| `bids` | `table` | List of bid orders |
| `timestamp` | `number` | Timestamp in milliseconds |

#### Example: Load

```lua
local order_book, err = client:OrderBook():load({ symbol = "symbol" })
```


### Ticker

Create an instance: `local ticker = client:Ticker(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | `number` | Lowest ask price |
| `bid` | `number` | Highest bid price |
| `high` | `number` | 24h high price |
| `last` | `number` | Last traded price |
| `low` | `number` | 24h low price |
| `symbol` | `string` | Trading pair symbol |
| `timestamp` | `number` | Timestamp in milliseconds |
| `volume` | `number` | 24h trading volume |

#### Example: Load

```lua
local ticker, err = client:Ticker():load({ id = "ticker_id" })
```

#### Example: List

```lua
local tickers, err = client:Ticker():list()
```


### Trade

Create an instance: `local trade = client:Trade(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` | Trade amount |
| `id` | `string` | Trade ID |
| `price` | `number` | Trade price |
| `side` | `string` | Trade side |
| `timestamp` | `number` | Timestamp in milliseconds |

#### Example: Load

```lua
local trade, err = client:Trade():load({ id = "trade_id" })
```


### Withdrawal

Create an instance: `local withdrawal = client:Withdrawal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountNumber` | `string` | Bank account number |
| `accountType` | `string` | Account type |
| `address` | `string` | Destination address |
| `agency` | `string` | Bank agency |
| `amount` | `number` | Withdrawal amount in BRL |
| `bank` | `string` | Bank code |
| `currency` | `string` | Cryptocurrency code |
| `tag` | `string` | Destination tag/memo (if applicable) |

#### Example: Create

```lua
local withdrawal, err = client:Withdrawal():create({
  accountNumber = "example_accountNumber", -- string
  address = "example_address", -- string
  agency = "example_agency", -- string
  amount = 1, -- number
  bank = "example_bank", -- string
  currency = "example_currency", -- string
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── mercado-bitcoin_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`mercado-bitcoin_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```lua
local trade = client:Trade()
trade:load({ id = "example_id" })

-- trade:data_get() now returns the trade data from the last load
-- trade:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.

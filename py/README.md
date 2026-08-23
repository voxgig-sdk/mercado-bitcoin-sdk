# MercadoBitcoin Python SDK



The Python SDK for the MercadoBitcoin API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Balance()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from mercadobitcoin_sdk import MercadoBitcoinSDK

client = MercadoBitcoinSDK({
    "apikey": os.environ.get("MERCADO_BITCOIN_APIKEY"),
})
```

### 2. List balance records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    balances = client.Balance().list()
    for balance in balances:
        print(balance)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load an orderbook

OrderBook is nested under symbol, so provide the `symbol`.
`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    orderbook = client.OrderBook().load({"symbol": "example_symbol"})
    print(orderbook)
except Exception as err:
    print(f"load failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    trade = client.Trade().load({"id": "example_id"})
    print(trade)
except Exception as err:
    print(f"load failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = MercadoBitcoinSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
trade = client.Trade().load({"id": "test01"})
# trade contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = MercadoBitcoinSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### MercadoBitcoinSDK

```python
from mercadobitcoin_sdk import MercadoBitcoinSDK

client = MercadoBitcoinSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = MercadoBitcoinSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### MercadoBitcoinSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `balance = client.Balance()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `available` | `float` | Available balance |
| `currency` | `str` | Currency code |
| `locked` | `float` | Locked balance |
| `total` | `float` | Total balance |

#### Example: List

```python
balances = client.Balance().list()
```


### Candle

Create an instance: `candle = client.Candle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `close` | `float` | Closing price |
| `high` | `float` | Highest price |
| `low` | `float` | Lowest price |
| `open` | `float` | Opening price |
| `timestamp` | `int` | Candle timestamp in milliseconds |
| `volume` | `float` | Trading volume |

#### Example: Load

```python
candle = client.Candle().load({"id": "candle_id"})
```


### DepositAddress

Create an instance: `deposit_address = client.DepositAddress()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `str` | Deposit address |
| `currency` | `str` | Cryptocurrency code |
| `qrCode` | `str` | QR code for deposit address |
| `tag` | `str` | Deposit tag/memo (if applicable) |

#### Example: Load

```python
deposit_address = client.DepositAddress().load()
```


### Order

Create an instance: `order = client.Order()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` | Order amount |
| `filled` | `float` | Filled amount |
| `id` | `str` | Order ID |
| `price` | `float` | Order price |
| `side` | `str` | Order side |
| `status` | `str` | Order status |
| `symbol` | `str` | Trading pair symbol |
| `timestamp` | `int` | Order creation timestamp |
| `type` | `str` | Order type |

#### Example: Load

```python
order = client.Order().load({"id": "order_id"})
```

#### Example: List

```python
orders = client.Order().list()
```

#### Example: Create

```python
order = client.Order().create({
})
```


### OrderBook

Create an instance: `order_book = client.OrderBook()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `asks` | `list` | List of ask orders |
| `bids` | `list` | List of bid orders |
| `timestamp` | `int` | Timestamp in milliseconds |

#### Example: Load

```python
order_book = client.OrderBook().load({"symbol": "symbol"})
```


### Ticker

Create an instance: `ticker = client.Ticker()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | `float` | Lowest ask price |
| `bid` | `float` | Highest bid price |
| `high` | `float` | 24h high price |
| `last` | `float` | Last traded price |
| `low` | `float` | 24h low price |
| `symbol` | `str` | Trading pair symbol |
| `timestamp` | `int` | Timestamp in milliseconds |
| `volume` | `float` | 24h trading volume |

#### Example: Load

```python
ticker = client.Ticker().load({"id": "ticker_id"})
```

#### Example: List

```python
tickers = client.Ticker().list()
```


### Trade

Create an instance: `trade = client.Trade()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` | Trade amount |
| `id` | `str` | Trade ID |
| `price` | `float` | Trade price |
| `side` | `str` | Trade side |
| `timestamp` | `int` | Timestamp in milliseconds |

#### Example: Load

```python
trade = client.Trade().load({"id": "trade_id"})
```


### Withdrawal

Create an instance: `withdrawal = client.Withdrawal()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountNumber` | `str` | Bank account number |
| `accountType` | `str` | Account type |
| `address` | `str` | Destination address |
| `agency` | `str` | Bank agency |
| `amount` | `float` | Withdrawal amount in BRL |
| `bank` | `str` | Bank code |
| `currency` | `str` | Cryptocurrency code |
| `tag` | `str` | Destination tag/memo (if applicable) |

#### Example: Create

```python
withdrawal = client.Withdrawal().create({
    "accountNumber": "example_accountNumber",  # str
    "address": "example_address",  # str
    "agency": "example_agency",  # str
    "amount": 1,  # float
    "bank": "example_bank",  # str
    "currency": "example_currency",  # str
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── mercadobitcoin_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`mercadobitcoin_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
trade = client.Trade()
trade.load({"id": "example_id"})

# trade.data_get() now returns the trade data from the last load
# trade.match_get() returns the last match criteria
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

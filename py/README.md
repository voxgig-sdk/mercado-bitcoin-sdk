# MercadoBitcoin Python SDK

The Python SDK for the MercadoBitcoin API. Provides an entity-oriented interface following Pythonic conventions.


## Install
```bash
pip install mercado-bitcoin-sdk
```

Or install from source:

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
    "apikey": os.environ.get("MERCADO-BITCOIN_APIKEY"),
})
```

### 2. List balances

```python
result, err = client.Balance(None).list(None, None)
if err:
    raise Exception(err)

if isinstance(result, list):
    for item in result:
        d = item.data_get()
        print(d["id"], d["name"])
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
if err:
    raise Exception(err)

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
```

### Prepare a request without sending it

```python
fetchdef, err = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})
if err:
    raise Exception(err)

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = MercadoBitcoinSDK.test(None, None)

result, err = client.MercadoBitcoin(None).load(
    {"id": "test01"}, None
)
# result contains mock response data
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
MERCADO-BITCOIN_TEST_LIVE=TRUE
MERCADO-BITCOIN_APIKEY=<your-key>
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
| `prepare` | `(fetchargs) -> (dict, err)` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> (dict, err)` | Build and send an HTTP request. |
| `Balance` | `(data) -> BalanceEntity` | Create a Balance entity instance. |
| `Candle` | `(data) -> CandleEntity` | Create a Candle entity instance. |
| `DepositAddress` | `(data) -> DepositAddressEntity` | Create a DepositAddress entity instance. |
| `Order` | `(data) -> OrderEntity` | Create a Order entity instance. |
| `OrderBook` | `(data) -> OrderBookEntity` | Create a OrderBook entity instance. |
| `Ticker` | `(data) -> TickerEntity` | Create a Ticker entity instance. |
| `Trade` | `(data) -> TradeEntity` | Create a Trade entity instance. |
| `Withdrawal` | `(data) -> WithdrawalEntity` | Create a Withdrawal entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> (any, err)` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> (any, err)` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> (any, err)` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> (any, err)` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> (any, err)` | Remove an entity. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return `(any, err)`. The first value is a
`dict` with these keys:

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

Create an instance: `const balance = client.Balance()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `available` | ``$NUMBER`` |  |
| `currency` | ``$STRING`` |  |
| `locked` | ``$NUMBER`` |  |
| `total` | ``$NUMBER`` |  |

#### Example: List

```ts
const balances = await client.Balance().list()
```


### Candle

Create an instance: `const candle = client.Candle()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `close` | ``$NUMBER`` |  |
| `high` | ``$NUMBER`` |  |
| `low` | ``$NUMBER`` |  |
| `open` | ``$NUMBER`` |  |
| `timestamp` | ``$INTEGER`` |  |
| `volume` | ``$NUMBER`` |  |

#### Example: Load

```ts
const candle = await client.Candle().load({ id: 'candle_id' })
```


### DepositAddress

Create an instance: `const deposit_address = client.DepositAddress()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$STRING`` |  |
| `currency` | ``$STRING`` |  |
| `qr_code` | ``$STRING`` |  |
| `tag` | ``$STRING`` |  |

#### Example: Load

```ts
const deposit_address = await client.DepositAddress().load({ id: 'deposit_address_id' })
```


### Order

Create an instance: `const order = client.Order()`

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
| `amount` | ``$NUMBER`` |  |
| `filled` | ``$NUMBER`` |  |
| `id` | ``$STRING`` |  |
| `price` | ``$NUMBER`` |  |
| `side` | ``$STRING`` |  |
| `status` | ``$STRING`` |  |
| `symbol` | ``$STRING`` |  |
| `timestamp` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: Load

```ts
const order = await client.Order().load({ id: 'order_id' })
```

#### Example: List

```ts
const orders = await client.Order().list()
```

#### Example: Create

```ts
const order = await client.Order().create({
})
```


### OrderBook

Create an instance: `const order_book = client.OrderBook()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | ``$ARRAY`` |  |
| `bid` | ``$ARRAY`` |  |
| `timestamp` | ``$INTEGER`` |  |

#### Example: Load

```ts
const order_book = await client.OrderBook().load({ id: 'order_book_id' })
```


### Ticker

Create an instance: `const ticker = client.Ticker()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | ``$NUMBER`` |  |
| `bid` | ``$NUMBER`` |  |
| `high` | ``$NUMBER`` |  |
| `last` | ``$NUMBER`` |  |
| `low` | ``$NUMBER`` |  |
| `symbol` | ``$STRING`` |  |
| `timestamp` | ``$INTEGER`` |  |
| `volume` | ``$NUMBER`` |  |

#### Example: Load

```ts
const ticker = await client.Ticker().load({ id: 'ticker_id' })
```

#### Example: List

```ts
const tickers = await client.Ticker().list()
```


### Trade

Create an instance: `const trade = client.Trade()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | ``$NUMBER`` |  |
| `id` | ``$STRING`` |  |
| `price` | ``$NUMBER`` |  |
| `side` | ``$STRING`` |  |
| `timestamp` | ``$INTEGER`` |  |

#### Example: Load

```ts
const trade = await client.Trade().load({ id: 'trade_id' })
```


### Withdrawal

Create an instance: `const withdrawal = client.Withdrawal()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | ``$STRING`` |  |
| `account_type` | ``$STRING`` |  |
| `address` | ``$STRING`` |  |
| `agency` | ``$STRING`` |  |
| `amount` | ``$NUMBER`` |  |
| `bank` | ``$STRING`` |  |
| `currency` | ``$STRING`` |  |
| `tag` | ``$STRING`` |  |

#### Example: Create

```ts
const withdrawal = await client.Withdrawal().create({
  account_number: /* `$STRING` */,
  address: /* `$STRING` */,
  agency: /* `$STRING` */,
  amount: /* `$NUMBER` */,
  bank: /* `$STRING` */,
  currency: /* `$STRING` */,
})
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

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

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller as the second element in the return tuple.

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
moon = client.Moon()
moon.load({"planet_id": "earth", "id": "luna"})

# moon.data_get() now returns the loaded moon data
# moon.match_get() returns the last match criteria
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

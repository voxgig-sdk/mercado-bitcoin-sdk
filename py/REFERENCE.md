# MercadoBitcoin Python SDK Reference

Complete API reference for the MercadoBitcoin Python SDK.


## MercadoBitcoinSDK

### Constructor

```python
from mercadobitcoin_sdk import MercadoBitcoinSDK

client = MercadoBitcoinSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `MercadoBitcoinSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = MercadoBitcoinSDK.test()
```


### Instance Methods

#### `Balance(data=None)`

Create a new `BalanceEntity` instance. Pass `None` for no initial data.

#### `Candle(data=None)`

Create a new `CandleEntity` instance. Pass `None` for no initial data.

#### `DepositAddress(data=None)`

Create a new `DepositAddressEntity` instance. Pass `None` for no initial data.

#### `Order(data=None)`

Create a new `OrderEntity` instance. Pass `None` for no initial data.

#### `OrderBook(data=None)`

Create a new `OrderBookEntity` instance. Pass `None` for no initial data.

#### `Ticker(data=None)`

Create a new `TickerEntity` instance. Pass `None` for no initial data.

#### `Trade(data=None)`

Create a new `TradeEntity` instance. Pass `None` for no initial data.

#### `Withdrawal(data=None)`

Create a new `WithdrawalEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## BalanceEntity

```python
balance = client.Balance()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | `float` | No | Available balance |
| `currency` | `str` | No | Currency code |
| `locked` | `float` | No | Locked balance |
| `total` | `float` | No | Total balance |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Balance().list()
for balance in results:
    print(balance)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BalanceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CandleEntity

```python
candle = client.Candle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `close` | `float` | No | Closing price |
| `high` | `float` | No | Highest price |
| `low` | `float` | No | Lowest price |
| `open` | `float` | No | Opening price |
| `timestamp` | `int` | No | Candle timestamp in milliseconds |
| `volume` | `float` | No | Trading volume |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Candle().load({"id": "candle_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CandleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DepositAddressEntity

```python
deposit_address = client.DepositAddress()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `str` | No | Deposit address |
| `currency` | `str` | No | Cryptocurrency code |
| `qrCode` | `str` | No | QR code for deposit address |
| `tag` | `str` | No | Deposit tag/memo (if applicable) |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.DepositAddress().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DepositAddressEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrderEntity

```python
order = client.Order()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No | Order amount |
| `filled` | `float` | No | Filled amount |
| `id` | `str` | No | Order ID |
| `price` | `float` | No | Order price |
| `side` | `str` | No | Order side |
| `status` | `str` | No | Order status |
| `symbol` | `str` | No | Trading pair symbol |
| `timestamp` | `int` | No | Order creation timestamp |
| `type` | `str` | No | Order type |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Order().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Order().list()
for order in results:
    print(order)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Order().load({"id": "order_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Order().remove({"id": "order_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrderEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrderBookEntity

```python
order_book = client.OrderBook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `asks` | `list` | No | List of ask orders |
| `bids` | `list` | No | List of bid orders |
| `timestamp` | `int` | No | Timestamp in milliseconds |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.OrderBook().load({"symbol": "symbol"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrderBookEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TickerEntity

```python
ticker = client.Ticker()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | `float` | No | Lowest ask price |
| `bid` | `float` | No | Highest bid price |
| `high` | `float` | No | 24h high price |
| `last` | `float` | No | Last traded price |
| `low` | `float` | No | 24h low price |
| `symbol` | `str` | No | Trading pair symbol |
| `timestamp` | `int` | No | Timestamp in milliseconds |
| `volume` | `float` | No | 24h trading volume |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Ticker().list()
for ticker in results:
    print(ticker)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Ticker().load({"id": "ticker_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TickerEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TradeEntity

```python
trade = client.Trade()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No | Trade amount |
| `id` | `str` | No | Trade ID |
| `price` | `float` | No | Trade price |
| `side` | `str` | No | Trade side |
| `timestamp` | `int` | No | Timestamp in milliseconds |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Trade().load({"id": "trade_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TradeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WithdrawalEntity

```python
withdrawal = client.Withdrawal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountNumber` | `str` | Yes | Bank account number |
| `accountType` | `str` | No | Account type |
| `address` | `str` | Yes | Destination address |
| `agency` | `str` | Yes | Bank agency |
| `amount` | `float` | Yes | Withdrawal amount in BRL |
| `bank` | `str` | Yes | Bank code |
| `currency` | `str` | Yes | Cryptocurrency code |
| `tag` | `str` | No | Destination tag/memo (if applicable) |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Withdrawal().create({
    "accountNumber": "example_accountNumber",  # str
    "address": "example_address",  # str
    "agency": "example_agency",  # str
    "amount": 1,  # float
    "bank": "example_bank",  # str
    "currency": "example_currency",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WithdrawalEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = MercadoBitcoinSDK({
    "feature": {
        "test": {"active": True},
    },
})
```


# MercadoBitcoin Python SDK Reference

Complete API reference for the MercadoBitcoin Python SDK.


## MercadoBitcoinSDK

### Constructor

```python
from mercado-bitcoin_sdk import MercadoBitcoinSDK

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
balance = client.balance
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | ``$NUMBER`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `locked` | ``$NUMBER`` | No |  |
| `total` | ``$NUMBER`` | No |  |

### Operations

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.balance.list({})
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
candle = client.candle
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.candle.load({"id": "candle_id"})
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
deposit_address = client.deposit_address
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `qr_code` | ``$STRING`` | No |  |
| `tag` | ``$STRING`` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.deposit_address.load({"id": "deposit_address_id"})
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
order = client.order
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.order.create({
})
```

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.order.list({})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.order.load({"id": "order_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.order.remove({"id": "order_id"})
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
order_book = client.order_book
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | ``$ARRAY`` | No |  |
| `bid` | ``$ARRAY`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.order_book.load({"id": "order_book_id"})
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
ticker = client.ticker
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.ticker.list({})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ticker.load({"id": "ticker_id"})
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
trade = client.trade
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.trade.load({"id": "trade_id"})
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
withdrawal = client.withdrawal
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.withdrawal.create({
    "account_number": # `$STRING`,
    "address": # `$STRING`,
    "agency": # `$STRING`,
    "amount": # `$NUMBER`,
    "bank": # `$STRING`,
    "currency": # `$STRING`,
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


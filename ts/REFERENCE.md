# MercadoBitcoin TypeScript SDK Reference

Complete API reference for the MercadoBitcoin TypeScript SDK.


## MercadoBitcoinSDK

### Constructor

```ts
new MercadoBitcoinSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `MercadoBitcoinSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = MercadoBitcoinSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `MercadoBitcoinSDK` instance in test mode.


### Instance Methods

#### `Balance(data?: object)`

Create a new `Balance` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BalanceEntity` instance.

#### `Candle(data?: object)`

Create a new `Candle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CandleEntity` instance.

#### `DepositAddress(data?: object)`

Create a new `DepositAddress` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DepositAddressEntity` instance.

#### `Order(data?: object)`

Create a new `Order` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrderEntity` instance.

#### `OrderBook(data?: object)`

Create a new `OrderBook` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrderBookEntity` instance.

#### `Ticker(data?: object)`

Create a new `Ticker` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TickerEntity` instance.

#### `Trade(data?: object)`

Create a new `Trade` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TradeEntity` instance.

#### `Withdrawal(data?: object)`

Create a new `Withdrawal` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WithdrawalEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `MercadoBitcoinSDK.test()`.

**Returns:** `MercadoBitcoinSDK` instance in test mode.


---

## BalanceEntity

```ts
const balance = client.Balance()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | `number` | No |  |
| `currency` | `string` | No |  |
| `locked` | `number` | No |  |
| `total` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Balance().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BalanceEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CandleEntity

```ts
const candle = client.Candle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `close` | `number` | No |  |
| `high` | `number` | No |  |
| `low` | `number` | No |  |
| `open` | `number` | No |  |
| `timestamp` | `number` | No |  |
| `volume` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Candle().load({ id: 'candle_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CandleEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DepositAddressEntity

```ts
const deposit_address = client.DepositAddress()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `string` | No |  |
| `currency` | `string` | No |  |
| `qr_code` | `string` | No |  |
| `tag` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.DepositAddress().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DepositAddressEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrderEntity

```ts
const order = client.Order()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `filled` | `number` | No |  |
| `id` | `string` | No |  |
| `price` | `number` | No |  |
| `side` | `string` | No |  |
| `status` | `string` | No |  |
| `symbol` | `string` | No |  |
| `timestamp` | `number` | No |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Order().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Order().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Order().load({ id: 'order_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Order().remove({ id: 'order_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrderEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrderBookEntity

```ts
const order_book = client.OrderBook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | `any[]` | No |  |
| `bid` | `any[]` | No |  |
| `timestamp` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.OrderBook().load({ symbol: 'symbol' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrderBookEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TickerEntity

```ts
const ticker = client.Ticker()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | `number` | No |  |
| `bid` | `number` | No |  |
| `high` | `number` | No |  |
| `last` | `number` | No |  |
| `low` | `number` | No |  |
| `symbol` | `string` | No |  |
| `timestamp` | `number` | No |  |
| `volume` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Ticker().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Ticker().load({ id: 'ticker_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TickerEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TradeEntity

```ts
const trade = client.Trade()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `id` | `string` | No |  |
| `price` | `number` | No |  |
| `side` | `string` | No |  |
| `timestamp` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Trade().load({ id: 'trade_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TradeEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WithdrawalEntity

```ts
const withdrawal = client.Withdrawal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number` | `string` | Yes |  |
| `account_type` | `string` | No |  |
| `address` | `string` | Yes |  |
| `agency` | `string` | Yes |  |
| `amount` | `number` | Yes |  |
| `bank` | `string` | Yes |  |
| `currency` | `string` | Yes |  |
| `tag` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Withdrawal().create({
  account_number: 'example_account_number',
  address: 'example_address',
  agency: 'example_agency',
  amount: 1,
  bank: 'example_bank',
  currency: 'example_currency',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WithdrawalEntity` instance with the same client and
options.

#### `client()`

Return the parent `MercadoBitcoinSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new MercadoBitcoinSDK({
  feature: {
    test: { active: true },
  }
})
```


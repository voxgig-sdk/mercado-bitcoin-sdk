# MercadoBitcoin PHP SDK Reference

Complete API reference for the MercadoBitcoin PHP SDK.


## MercadoBitcoinSDK

### Constructor

```php
require_once __DIR__ . '/mercadobitcoin_sdk.php';

$client = new MercadoBitcoinSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `MercadoBitcoinSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = MercadoBitcoinSDK::test();
```


### Instance Methods

#### `Balance($data = null)`

Create a new `BalanceEntity` instance. Pass `null` for no initial data.

#### `Candle($data = null)`

Create a new `CandleEntity` instance. Pass `null` for no initial data.

#### `DepositAddress($data = null)`

Create a new `DepositAddressEntity` instance. Pass `null` for no initial data.

#### `Order($data = null)`

Create a new `OrderEntity` instance. Pass `null` for no initial data.

#### `OrderBook($data = null)`

Create a new `OrderBookEntity` instance. Pass `null` for no initial data.

#### `Ticker($data = null)`

Create a new `TickerEntity` instance. Pass `null` for no initial data.

#### `Trade($data = null)`

Create a new `TradeEntity` instance. Pass `null` for no initial data.

#### `Withdrawal($data = null)`

Create a new `WithdrawalEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): MercadoBitcoinUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## BalanceEntity

```php
$balance = $client->Balance();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | `float` | No | Available balance |
| `currency` | `string` | No | Currency code |
| `locked` | `float` | No | Locked balance |
| `total` | `float` | No | Total balance |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Balance()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BalanceEntity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CandleEntity

```php
$candle = $client->Candle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `close` | `float` | No | Closing price |
| `high` | `float` | No | Highest price |
| `id` | `string` | No |  |
| `low` | `float` | No | Lowest price |
| `open` | `float` | No | Opening price |
| `timestamp` | `int` | No | Candle timestamp in milliseconds |
| `volume` | `float` | No | Trading volume |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Candle()->load(["id" => "candle_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CandleEntity`

Create a new `CandleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DepositAddressEntity

```php
$deposit_address = $client->DepositAddress();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `string` | No | Deposit address |
| `currency` | `string` | No | Cryptocurrency code |
| `qrCode` | `string` | No | QR code for deposit address |
| `tag` | `string` | No | Deposit tag/memo (if applicable) |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->DepositAddress()->load(["currency" => "currency"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DepositAddressEntity`

Create a new `DepositAddressEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrderEntity

```php
$order = $client->Order();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No | Order amount |
| `filled` | `float` | No | Filled amount |
| `id` | `string` | No | Order ID |
| `price` | `float` | No | Order price |
| `side` | `string` | No | Order side |
| `status` | `string` | No | Order status |
| `symbol` | `string` | No | Trading pair symbol |
| `timestamp` | `int` | No | Order creation timestamp |
| `type` | `string` | No | Order type |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Order()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Order()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Order()->load(["id" => "order_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Order()->remove(["id" => "order_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrderEntity`

Create a new `OrderEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrderBookEntity

```php
$order_book = $client->OrderBook();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `asks` | `array` | No | List of ask orders |
| `bids` | `array` | No | List of bid orders |
| `timestamp` | `int` | No | Timestamp in milliseconds |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->OrderBook()->load(["symbol" => "symbol"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrderBookEntity`

Create a new `OrderBookEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TickerEntity

```php
$ticker = $client->Ticker();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | `float` | No | Lowest ask price |
| `bid` | `float` | No | Highest bid price |
| `high` | `float` | No | 24h high price |
| `id` | `string` | No |  |
| `last` | `float` | No | Last traded price |
| `low` | `float` | No | 24h low price |
| `symbol` | `string` | No | Trading pair symbol |
| `timestamp` | `int` | No | Timestamp in milliseconds |
| `volume` | `float` | No | 24h trading volume |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Ticker()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Ticker()->load(["id" => "ticker_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TickerEntity`

Create a new `TickerEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TradeEntity

```php
$trade = $client->Trade();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No | Trade amount |
| `id` | `string` | No | Trade ID |
| `price` | `float` | No | Trade price |
| `side` | `string` | No | Trade side |
| `timestamp` | `int` | No | Timestamp in milliseconds |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Trade()->load(["id" => "trade_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TradeEntity`

Create a new `TradeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WithdrawalEntity

```php
$withdrawal = $client->Withdrawal();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountNumber` | `string` | Yes | Bank account number |
| `accountType` | `string` | No | Account type |
| `address` | `string` | Yes | Destination address |
| `agency` | `string` | Yes | Bank agency |
| `amount` | `float` | Yes | Withdrawal amount in BRL |
| `bank` | `string` | Yes | Bank code |
| `currency` | `string` | Yes | Cryptocurrency code |
| `tag` | `string` | No | Destination tag/memo (if applicable) |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Withdrawal()->create([
  "accountNumber" => null, // string
  "address" => null, // string
  "agency" => null, // string
  "amount" => null, // float
  "bank" => null, // string
  "currency" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WithdrawalEntity`

Create a new `WithdrawalEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new MercadoBitcoinSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.


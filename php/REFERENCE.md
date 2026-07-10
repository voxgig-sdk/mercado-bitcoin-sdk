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
| `available` | `float` | No |  |
| `currency` | `string` | No |  |
| `locked` | `float` | No |  |
| `total` | `float` | No |  |

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
| `close` | `float` | No |  |
| `high` | `float` | No |  |
| `low` | `float` | No |  |
| `open` | `float` | No |  |
| `timestamp` | `int` | No |  |
| `volume` | `float` | No |  |

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
| `address` | `string` | No |  |
| `currency` | `string` | No |  |
| `qr_code` | `string` | No |  |
| `tag` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->DepositAddress()->load();
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
| `amount` | `float` | No |  |
| `filled` | `float` | No |  |
| `id` | `string` | No |  |
| `price` | `float` | No |  |
| `side` | `string` | No |  |
| `status` | `string` | No |  |
| `symbol` | `string` | No |  |
| `timestamp` | `int` | No |  |
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
| `ask` | `array` | No |  |
| `bid` | `array` | No |  |
| `timestamp` | `int` | No |  |

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
| `ask` | `float` | No |  |
| `bid` | `float` | No |  |
| `high` | `float` | No |  |
| `last` | `float` | No |  |
| `low` | `float` | No |  |
| `symbol` | `string` | No |  |
| `timestamp` | `int` | No |  |
| `volume` | `float` | No |  |

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
| `amount` | `float` | No |  |
| `id` | `string` | No |  |
| `price` | `float` | No |  |
| `side` | `string` | No |  |
| `timestamp` | `int` | No |  |

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
| `account_number` | `string` | Yes |  |
| `account_type` | `string` | No |  |
| `address` | `string` | Yes |  |
| `agency` | `string` | Yes |  |
| `amount` | `float` | Yes |  |
| `bank` | `string` | Yes |  |
| `currency` | `string` | Yes |  |
| `tag` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Withdrawal()->create([
  "account_number" => null, // string
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


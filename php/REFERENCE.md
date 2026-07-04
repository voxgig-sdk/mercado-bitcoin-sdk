# MercadoBitcoin PHP SDK Reference

Complete API reference for the MercadoBitcoin PHP SDK.


## MercadoBitcoinSDK

### Constructor

```php
require_once __DIR__ . '/mercado-bitcoin_sdk.php';

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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

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
$balance = $client->balance();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `available` | ``$NUMBER`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `locked` | ``$NUMBER`` | No |  |
| `total` | ``$NUMBER`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->balance()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): BalanceEntity`

Create a new `BalanceEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## CandleEntity

```php
$candle = $client->candle();
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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->candle()->load(["id" => "candle_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): CandleEntity`

Create a new `CandleEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## DepositAddressEntity

```php
$deposit_address = $client->deposit_address();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `currency` | ``$STRING`` | No |  |
| `qr_code` | ``$STRING`` | No |  |
| `tag` | ``$STRING`` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->deposit_address()->load(["id" => "deposit_address_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): DepositAddressEntity`

Create a new `DepositAddressEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## OrderEntity

```php
$order = $client->order();
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->order()->create([
]);
```

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->order()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->order()->load(["id" => "order_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->order()->remove(["id" => "order_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): OrderEntity`

Create a new `OrderEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## OrderBookEntity

```php
$order_book = $client->order_book();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ask` | ``$ARRAY`` | No |  |
| `bid` | ``$ARRAY`` | No |  |
| `timestamp` | ``$INTEGER`` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->order_book()->load(["id" => "order_book_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): OrderBookEntity`

Create a new `OrderBookEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TickerEntity

```php
$ticker = $client->ticker();
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

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->ticker()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ticker()->load(["id" => "ticker_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TickerEntity`

Create a new `TickerEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TradeEntity

```php
$trade = $client->trade();
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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->trade()->load(["id" => "trade_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TradeEntity`

Create a new `TradeEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## WithdrawalEntity

```php
$withdrawal = $client->withdrawal();
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->withdrawal()->create([
  "account_number" => /* `$STRING` */,
  "address" => /* `$STRING` */,
  "agency" => /* `$STRING` */,
  "amount" => /* `$NUMBER` */,
  "bank" => /* `$STRING` */,
  "currency" => /* `$STRING` */,
]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): WithdrawalEntity`

Create a new `WithdrawalEntity` instance with the same client and
options.

#### `getName(): string`

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


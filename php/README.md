# MercadoBitcoin PHP SDK



The PHP SDK for the MercadoBitcoin API — an entity-oriented client using PHP conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases](https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'mercadobitcoin_sdk.php';

$client = new MercadoBitcoinSDK([
    "apikey" => getenv("MERCADO_BITCOIN_APIKEY"),
]);
```

### 2. List balance records

```php
try {
    // list() returns an array of Balance records — iterate directly.
    $balances = $client->Balance()->list();
    foreach ($balances as $item) {
        echo $item["id"] . " " . $item["name"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    echo "Error: " . $result["err"]->getMessage();
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = MercadoBitcoinSDK::test([
    "entity" => ["balance" => ["test01" => ["id" => "test01"]]],
]);

// load() returns the bare mock record (throws on error).
$balance = $client->Balance()->load(["id" => "test01"]);
print_r($balance);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new MercadoBitcoinSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
MERCADO_BITCOIN_TEST_LIVE=TRUE
MERCADO_BITCOIN_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### MercadoBitcoinSDK

```php
require_once 'mercadobitcoin_sdk.php';
$client = new MercadoBitcoinSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = MercadoBitcoinSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### MercadoBitcoinSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Balance` | `($data): BalanceEntity` | Create a Balance entity instance. |
| `Candle` | `($data): CandleEntity` | Create a Candle entity instance. |
| `DepositAddress` | `($data): DepositAddressEntity` | Create a DepositAddress entity instance. |
| `Order` | `($data): OrderEntity` | Create an Order entity instance. |
| `OrderBook` | `($data): OrderBookEntity` | Create an OrderBook entity instance. |
| `Ticker` | `($data): TickerEntity` | Create a Ticker entity instance. |
| `Trade` | `($data): TradeEntity` | Create a Trade entity instance. |
| `Withdrawal` | `($data): WithdrawalEntity` | Create a Withdrawal entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `($reqmatch, $ctrl): array` | List entities matching the criteria. |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$balance = $client->Balance();`

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

```php
// list() returns an array of Balance records (throws on error).
$balances = $client->Balance()->list();
```


### Candle

Create an instance: `$candle = $client->Candle();`

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

```php
// load() returns the bare Candle record (throws on error).
$candle = $client->Candle()->load(["id" => "candle_id"]);
```


### DepositAddress

Create an instance: `$deposit_address = $client->DepositAddress();`

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

```php
// load() returns the bare DepositAddress record (throws on error).
$deposit_address = $client->DepositAddress()->load(["id" => "deposit_address_id"]);
```


### Order

Create an instance: `$order = $client->Order();`

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

```php
// load() returns the bare Order record (throws on error).
$order = $client->Order()->load(["id" => "order_id"]);
```

#### Example: List

```php
// list() returns an array of Order records (throws on error).
$orders = $client->Order()->list();
```

#### Example: Create

```php
$order = $client->Order()->create([
]);
```


### OrderBook

Create an instance: `$order_book = $client->OrderBook();`

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

```php
// load() returns the bare OrderBook record (throws on error).
$order_book = $client->OrderBook()->load(["id" => "order_book_id"]);
```


### Ticker

Create an instance: `$ticker = $client->Ticker();`

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

```php
// load() returns the bare Ticker record (throws on error).
$ticker = $client->Ticker()->load(["id" => "ticker_id"]);
```

#### Example: List

```php
// list() returns an array of Ticker records (throws on error).
$tickers = $client->Ticker()->list();
```


### Trade

Create an instance: `$trade = $client->Trade();`

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

```php
// load() returns the bare Trade record (throws on error).
$trade = $client->Trade()->load(["id" => "trade_id"]);
```


### Withdrawal

Create an instance: `$withdrawal = $client->Withdrawal();`

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

```php
$withdrawal = $client->Withdrawal()->create([
    "account_number" => null, // `$STRING`
    "address" => null, // `$STRING`
    "agency" => null, // `$STRING`
    "amount" => null, // `$NUMBER`
    "bank" => null, // `$STRING`
    "currency" => null, // `$STRING`
]);
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
error is returned to the caller as the second element in the return array.

### Features and hooks

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── mercadobitcoin_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`mercadobitcoin_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$balance = $client->Balance();
$balance->load(["id" => "example_id"]);

// $balance->dataGet() now returns the loaded balance data
// $balance->matchGet() returns the last match criteria
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

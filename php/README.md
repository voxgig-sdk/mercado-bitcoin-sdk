# MercadoBitcoin PHP SDK



The PHP SDK for the MercadoBitcoin API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Balance()` — with named operations (`list`/`load`/`create`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

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
        echo $item["available"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load an orderbook

OrderBook is nested under symbol, so provide the `symbol`.

```php
try {
    // load() returns the ENTITY — call data_get() for the OrderBook record (throws on error).
    $orderbook = $client->OrderBook()->load(["symbol" => "example_symbol"]);
    print_r($orderbook);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $trade = $client->Trade()->load(["id" => "example_id"]);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
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
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
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
    "entity" => ["trade" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$trade = $client->Trade()->load(["id" => "test01"]);
print_r($trade);
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
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
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
| `id` |  |
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
| `id` |  |
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

Create an instance: `$balance = $client->Balance();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `available` | `float` | Available balance |
| `currency` | `string` | Currency code |
| `locked` | `float` | Locked balance |
| `total` | `float` | Total balance |

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
| `close` | `float` | Closing price |
| `high` | `float` | Highest price |
| `id` | `string` |  |
| `low` | `float` | Lowest price |
| `open` | `float` | Opening price |
| `timestamp` | `int` | Candle timestamp in milliseconds |
| `volume` | `float` | Trading volume |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Candle record (throws on error).
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
| `address` | `string` | Deposit address |
| `currency` | `string` | Cryptocurrency code |
| `qrCode` | `string` | QR code for deposit address |
| `tag` | `string` | Deposit tag/memo (if applicable) |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the DepositAddress record (throws on error).
$deposit_address = $client->DepositAddress()->load();
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
| `amount` | `float` | Order amount |
| `filled` | `float` | Filled amount |
| `id` | `string` | Order ID |
| `price` | `float` | Order price |
| `side` | `string` | Order side |
| `status` | `string` | Order status |
| `symbol` | `string` | Trading pair symbol |
| `timestamp` | `int` | Order creation timestamp |
| `type` | `string` | Order type |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Order record (throws on error).
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
| `asks` | `array` | List of ask orders |
| `bids` | `array` | List of bid orders |
| `timestamp` | `int` | Timestamp in milliseconds |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the OrderBook record (throws on error).
$order_book = $client->OrderBook()->load(["symbol" => "symbol"]);
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
| `ask` | `float` | Lowest ask price |
| `bid` | `float` | Highest bid price |
| `high` | `float` | 24h high price |
| `id` | `string` |  |
| `last` | `float` | Last traded price |
| `low` | `float` | 24h low price |
| `symbol` | `string` | Trading pair symbol |
| `timestamp` | `int` | Timestamp in milliseconds |
| `volume` | `float` | 24h trading volume |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Ticker record (throws on error).
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
| `amount` | `float` | Trade amount |
| `id` | `string` | Trade ID |
| `price` | `float` | Trade price |
| `side` | `string` | Trade side |
| `timestamp` | `int` | Timestamp in milliseconds |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Trade record (throws on error).
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
| `accountNumber` | `string` | Bank account number |
| `accountType` | `string` | Account type |
| `address` | `string` | Destination address |
| `agency` | `string` | Bank agency |
| `amount` | `float` | Withdrawal amount in BRL |
| `bank` | `string` | Bank code |
| `currency` | `string` | Cryptocurrency code |
| `tag` | `string` | Destination tag/memo (if applicable) |

#### Example: Create

```php
$withdrawal = $client->Withdrawal()->create([
    "accountNumber" => null, // string
    "address" => null, // string
    "agency" => null, // string
    "amount" => null, // float
    "bank" => null, // string
    "currency" => null, // string
]);
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
$trade = $client->Trade();
$trade->load(["id" => "example_id"]);

// $trade->data_get() now returns the trade data from the last load
// $trade->match_get() returns the last match criteria
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

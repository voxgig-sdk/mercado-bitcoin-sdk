# MercadoBitcoin Golang SDK



The Golang SDK for the MercadoBitcoin API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Balance(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/mercado-bitcoin-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/mercado-bitcoin-sdk/go=../mercado-bitcoin-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/mercado-bitcoin-sdk/go"
)

func main() {
    client := sdk.NewMercadoBitcoinSDK(map[string]any{
        "apikey": os.Getenv("MERCADO_BITCOIN_APIKEY"),
    })

    // List balance records — the value is the array of records itself.
    balances, err := client.Balance(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range balances.([]any) {
        fmt.Println(item)
    }
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
trade, err := client.Trade(nil).Load(map[string]any{"id": "example_id"}, nil)
if err != nil {
    // handle err
    return
}
_ = trade
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

trade, err := client.Trade(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(trade) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewMercadoBitcoinSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewMercadoBitcoinSDK

```go
func NewMercadoBitcoinSDK(options map[string]any) *MercadoBitcoinSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *MercadoBitcoinSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### MercadoBitcoinSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Balance` | `(data map[string]any) MercadoBitcoinEntity` | Create a Balance entity instance. |
| `Candle` | `(data map[string]any) MercadoBitcoinEntity` | Create a Candle entity instance. |
| `DepositAddress` | `(data map[string]any) MercadoBitcoinEntity` | Create a DepositAddress entity instance. |
| `Order` | `(data map[string]any) MercadoBitcoinEntity` | Create an Order entity instance. |
| `OrderBook` | `(data map[string]any) MercadoBitcoinEntity` | Create an OrderBook entity instance. |
| `Ticker` | `(data map[string]any) MercadoBitcoinEntity` | Create a Ticker entity instance. |
| `Trade` | `(data map[string]any) MercadoBitcoinEntity` | Create a Trade entity instance. |
| `Withdrawal` | `(data map[string]any) MercadoBitcoinEntity` | Create a Withdrawal entity instance. |

### Entity interface (MercadoBitcoinEntity)

All entities implement the `MercadoBitcoinEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    balance, err := client.Balance(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // balance is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Balance

| Field | Description |
| --- | --- |
| `"available"` | Available balance |
| `"currency"` | Currency code |
| `"locked"` | Locked balance |
| `"total"` | Total balance |

Operations: List.

API path: `/accounts/balance`

#### Candle

| Field | Description |
| --- | --- |
| `"close"` | Closing price |
| `"high"` | Highest price |
| `"id"` |  |
| `"low"` | Lowest price |
| `"open"` | Opening price |
| `"timestamp"` | Candle timestamp in milliseconds |
| `"volume"` | Trading volume |

Operations: Load.

API path: `/candles/{symbol}`

#### DepositAddress

| Field | Description |
| --- | --- |
| `"address"` | Deposit address |
| `"currency"` | Cryptocurrency code |
| `"qrCode"` | QR code for deposit address |
| `"tag"` | Deposit tag/memo (if applicable) |

Operations: Load.

API path: `/deposits/crypto`

#### Order

| Field | Description |
| --- | --- |
| `"amount"` | Order amount |
| `"filled"` | Filled amount |
| `"id"` | Order ID |
| `"price"` | Order price |
| `"side"` | Order side |
| `"status"` | Order status |
| `"symbol"` | Trading pair symbol |
| `"timestamp"` | Order creation timestamp |
| `"type"` | Order type |

Operations: Create, List, Load, Remove.

API path: `/orders`

#### OrderBook

| Field | Description |
| --- | --- |
| `"asks"` | List of ask orders |
| `"bids"` | List of bid orders |
| `"timestamp"` | Timestamp in milliseconds |

Operations: Load.

API path: `/orderbook/{symbol}`

#### Ticker

| Field | Description |
| --- | --- |
| `"ask"` | Lowest ask price |
| `"bid"` | Highest bid price |
| `"high"` | 24h high price |
| `"id"` |  |
| `"last"` | Last traded price |
| `"low"` | 24h low price |
| `"symbol"` | Trading pair symbol |
| `"timestamp"` | Timestamp in milliseconds |
| `"volume"` | 24h trading volume |

Operations: List, Load.

API path: `/tickers`

#### Trade

| Field | Description |
| --- | --- |
| `"amount"` | Trade amount |
| `"id"` | Trade ID |
| `"price"` | Trade price |
| `"side"` | Trade side |
| `"timestamp"` | Timestamp in milliseconds |

Operations: Load.

API path: `/trades/{symbol}`

#### Withdrawal

| Field | Description |
| --- | --- |
| `"accountNumber"` | Bank account number |
| `"accountType"` | Account type |
| `"address"` | Destination address |
| `"agency"` | Bank agency |
| `"amount"` | Withdrawal amount in BRL |
| `"bank"` | Bank code |
| `"currency"` | Cryptocurrency code |
| `"tag"` | Destination tag/memo (if applicable) |

Operations: Create.

API path: `/withdrawals/brl`



## Entities


### Balance

Create an instance: `balance := client.Balance(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `available` | `float64` | Available balance |
| `currency` | `string` | Currency code |
| `locked` | `float64` | Locked balance |
| `total` | `float64` | Total balance |

#### Example: List

```go
balances, err := client.Balance(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(balances) // the array of records
```


### Candle

Create an instance: `candle := client.Candle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `close` | `float64` | Closing price |
| `high` | `float64` | Highest price |
| `id` | `string` |  |
| `low` | `float64` | Lowest price |
| `open` | `float64` | Opening price |
| `timestamp` | `int` | Candle timestamp in milliseconds |
| `volume` | `float64` | Trading volume |

#### Example: Load

```go
candle, err := client.Candle(nil).Load(map[string]any{"id": "candle_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(candle) // the loaded record
```


### DepositAddress

Create an instance: `depositAddress := client.DepositAddress(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `string` | Deposit address |
| `currency` | `string` | Cryptocurrency code |
| `qrCode` | `string` | QR code for deposit address |
| `tag` | `string` | Deposit tag/memo (if applicable) |

#### Example: Load

```go
depositAddress, err := client.DepositAddress(nil).Load(map[string]any{"currency": "currency"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(depositAddress) // the loaded record
```


### Order

Create an instance: `order := client.Order(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` | Order amount |
| `filled` | `float64` | Filled amount |
| `id` | `string` | Order ID |
| `price` | `float64` | Order price |
| `side` | `string` | Order side |
| `status` | `string` | Order status |
| `symbol` | `string` | Trading pair symbol |
| `timestamp` | `int` | Order creation timestamp |
| `type` | `string` | Order type |

#### Example: Load

```go
order, err := client.Order(nil).Load(map[string]any{"id": "order_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(order) // the loaded record
```

#### Example: List

```go
orders, err := client.Order(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(orders) // the array of records
```

#### Example: Create

```go
result, err := client.Order(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### OrderBook

Create an instance: `orderBook := client.OrderBook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `asks` | `[]any` | List of ask orders |
| `bids` | `[]any` | List of bid orders |
| `timestamp` | `int` | Timestamp in milliseconds |

#### Example: Load

```go
orderBook, err := client.OrderBook(nil).Load(map[string]any{"symbol": "symbol"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(orderBook) // the loaded record
```


### Ticker

Create an instance: `ticker := client.Ticker(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | `float64` | Lowest ask price |
| `bid` | `float64` | Highest bid price |
| `high` | `float64` | 24h high price |
| `id` | `string` |  |
| `last` | `float64` | Last traded price |
| `low` | `float64` | 24h low price |
| `symbol` | `string` | Trading pair symbol |
| `timestamp` | `int` | Timestamp in milliseconds |
| `volume` | `float64` | 24h trading volume |

#### Example: Load

```go
ticker, err := client.Ticker(nil).Load(map[string]any{"id": "ticker_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(ticker) // the loaded record
```

#### Example: List

```go
tickers, err := client.Ticker(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(tickers) // the array of records
```


### Trade

Create an instance: `trade := client.Trade(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` | Trade amount |
| `id` | `string` | Trade ID |
| `price` | `float64` | Trade price |
| `side` | `string` | Trade side |
| `timestamp` | `int` | Timestamp in milliseconds |

#### Example: Load

```go
trade, err := client.Trade(nil).Load(map[string]any{"id": "trade_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(trade) // the loaded record
```


### Withdrawal

Create an instance: `withdrawal := client.Withdrawal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountNumber` | `string` | Bank account number |
| `accountType` | `string` | Account type |
| `address` | `string` | Destination address |
| `agency` | `string` | Bank agency |
| `amount` | `float64` | Withdrawal amount in BRL |
| `bank` | `string` | Bank code |
| `currency` | `string` | Cryptocurrency code |
| `tag` | `string` | Destination tag/memo (if applicable) |

#### Example: Create

```go
result, err := client.Withdrawal(nil).Create(map[string]any{
    "accountNumber": "example_accountNumber",
    "address": "example_address",
    "agency": "example_agency",
    "amount": 1,
    "bank": "example_bank",
    "currency": "example_currency",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/mercado-bitcoin-sdk/go/
├── mercado-bitcoin.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/mercado-bitcoin-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
trade := client.Trade(nil)
trade.Load(map[string]any{"id": "example_id"}, nil)

// trade.Data() now returns the trade data from the last load
// trade.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.

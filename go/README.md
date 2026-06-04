# MercadoBitcoin Golang SDK

The Golang SDK for the MercadoBitcoin API. Provides an entity-oriented interface using standard Go conventions — no generics required, data flows as `map[string]any`.


## Install
```bash
go get github.com/voxgig-sdk/mercado-bitcoin-sdk/go
```

If the module is not yet published to a registry, use a `replace` directive
in your `go.mod` to point to a local checkout:

```bash
go mod edit -replace github.com/voxgig-sdk/mercado-bitcoin-sdk/go=../path/to/github.com/voxgig-sdk/mercado-bitcoin-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```go
package main

import (
    "fmt"

    sdk "github.com/voxgig-sdk/mercado-bitcoin-sdk/go"
    "github.com/voxgig-sdk/mercado-bitcoin-sdk/go/core"
)

func main() {
    client := sdk.NewMercadoBitcoinSDK(map[string]any{})
```

### 2. List balances

```go
    result, err := client.Balance(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }

    rm := core.ToMapAny(result)
    if rm["ok"] == true {
        for _, item := range rm["data"].([]any) {
            p := core.ToMapAny(item)
            fmt.Println(p["id"], p["name"])
        }
    }
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
client := sdk.TestSDK(nil, nil)

result, err := client.Planet(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
// result contains mock response data
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
MERCADO-BITCOIN_TEST_LIVE=TRUE
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
| `Order` | `(data map[string]any) MercadoBitcoinEntity` | Create a Order entity instance. |
| `OrderBook` | `(data map[string]any) MercadoBitcoinEntity` | Create a OrderBook entity instance. |
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
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(any, error)`. The `any` value is a
`map[string]any` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `"ok"` | `bool` | `true` if the HTTP status is 2xx. |
| `"status"` | `int` | HTTP status code. |
| `"headers"` | `map[string]any` | Response headers. |
| `"data"` | `any` | Parsed JSON response body. |

On error, `"ok"` is `false` and `"err"` contains the error value.

### Entities

#### Balance

| Field | Description |
| --- | --- |
| `"available"` |  |
| `"currency"` |  |
| `"locked"` |  |
| `"total"` |  |

Operations: List.

API path: `/accounts/balance`

#### Candle

| Field | Description |
| --- | --- |
| `"close"` |  |
| `"high"` |  |
| `"low"` |  |
| `"open"` |  |
| `"timestamp"` |  |
| `"volume"` |  |

Operations: Load.

API path: `/candles/{symbol}`

#### DepositAddress

| Field | Description |
| --- | --- |
| `"address"` |  |
| `"currency"` |  |
| `"qr_code"` |  |
| `"tag"` |  |

Operations: Load.

API path: `/deposits/crypto`

#### Order

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"filled"` |  |
| `"id"` |  |
| `"price"` |  |
| `"side"` |  |
| `"status"` |  |
| `"symbol"` |  |
| `"timestamp"` |  |
| `"type"` |  |

Operations: Create, List, Load, Remove.

API path: `/orders`

#### OrderBook

| Field | Description |
| --- | --- |
| `"ask"` |  |
| `"bid"` |  |
| `"timestamp"` |  |

Operations: Load.

API path: `/orderbook/{symbol}`

#### Ticker

| Field | Description |
| --- | --- |
| `"ask"` |  |
| `"bid"` |  |
| `"high"` |  |
| `"last"` |  |
| `"low"` |  |
| `"symbol"` |  |
| `"timestamp"` |  |
| `"volume"` |  |

Operations: List, Load.

API path: `/tickers`

#### Trade

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"id"` |  |
| `"price"` |  |
| `"side"` |  |
| `"timestamp"` |  |

Operations: Load.

API path: `/trades/{symbol}`

#### Withdrawal

| Field | Description |
| --- | --- |
| `"account_number"` |  |
| `"account_type"` |  |
| `"address"` |  |
| `"agency"` |  |
| `"amount"` |  |
| `"bank"` |  |
| `"currency"` |  |
| `"tag"` |  |

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
| `available` | ``$NUMBER`` |  |
| `currency` | ``$STRING`` |  |
| `locked` | ``$NUMBER`` |  |
| `total` | ``$NUMBER`` |  |

#### Example: List

```go
results, err := client.Balance(nil).List(nil, nil)
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
| `close` | ``$NUMBER`` |  |
| `high` | ``$NUMBER`` |  |
| `low` | ``$NUMBER`` |  |
| `open` | ``$NUMBER`` |  |
| `timestamp` | ``$INTEGER`` |  |
| `volume` | ``$NUMBER`` |  |

#### Example: Load

```go
result, err := client.Candle(nil).Load(map[string]any{"id": "candle_id"}, nil)
```


### DepositAddress

Create an instance: `deposit_address := client.DepositAddress(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$STRING`` |  |
| `currency` | ``$STRING`` |  |
| `qr_code` | ``$STRING`` |  |
| `tag` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.DepositAddress(nil).Load(map[string]any{"id": "deposit_address_id"}, nil)
```


### Order

Create an instance: `order := client.Order(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

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

```go
result, err := client.Order(nil).Load(map[string]any{"id": "order_id"}, nil)
```

#### Example: List

```go
results, err := client.Order(nil).List(nil, nil)
```

#### Example: Create

```go
result, err := client.Order(nil).Create(map[string]any{
}, nil)
```


### OrderBook

Create an instance: `order_book := client.OrderBook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ask` | ``$ARRAY`` |  |
| `bid` | ``$ARRAY`` |  |
| `timestamp` | ``$INTEGER`` |  |

#### Example: Load

```go
result, err := client.OrderBook(nil).Load(map[string]any{"id": "order_book_id"}, nil)
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
| `ask` | ``$NUMBER`` |  |
| `bid` | ``$NUMBER`` |  |
| `high` | ``$NUMBER`` |  |
| `last` | ``$NUMBER`` |  |
| `low` | ``$NUMBER`` |  |
| `symbol` | ``$STRING`` |  |
| `timestamp` | ``$INTEGER`` |  |
| `volume` | ``$NUMBER`` |  |

#### Example: Load

```go
result, err := client.Ticker(nil).Load(map[string]any{"id": "ticker_id"}, nil)
```

#### Example: List

```go
results, err := client.Ticker(nil).List(nil, nil)
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
| `amount` | ``$NUMBER`` |  |
| `id` | ``$STRING`` |  |
| `price` | ``$NUMBER`` |  |
| `side` | ``$STRING`` |  |
| `timestamp` | ``$INTEGER`` |  |

#### Example: Load

```go
result, err := client.Trade(nil).Load(map[string]any{"id": "trade_id"}, nil)
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
| `account_number` | ``$STRING`` |  |
| `account_type` | ``$STRING`` |  |
| `address` | ``$STRING`` |  |
| `agency` | ``$STRING`` |  |
| `amount` | ``$NUMBER`` |  |
| `bank` | ``$STRING`` |  |
| `currency` | ``$STRING`` |  |
| `tag` | ``$STRING`` |  |

#### Example: Create

```go
result, err := client.Withdrawal(nil).Create(map[string]any{
    "account_number": /* `$STRING` */,
    "address": /* `$STRING` */,
    "agency": /* `$STRING` */,
    "amount": /* `$NUMBER` */,
    "bank": /* `$STRING` */,
    "currency": /* `$STRING` */,
}, nil)
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
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

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
moon := client.Moon(nil)
moon.Load(map[string]any{"planet_id": "earth", "id": "luna"}, nil)

// moon.Data() now returns the loaded moon data
// moon.Match() returns the last match criteria
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

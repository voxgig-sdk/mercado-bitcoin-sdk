# MercadoBitcoin TypeScript SDK



The TypeScript SDK for the MercadoBitcoin API — a type-safe, entity-oriented client with full async/await support.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases](https://github.com/voxgig-sdk/mercado-bitcoin-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { MercadoBitcoinSDK } from '@voxgig-sdk/mercado-bitcoin'

const client = new MercadoBitcoinSDK({
  apikey: process.env.MERCADO_BITCOIN_APIKEY,
})
```

### 2. List balance records

`list()` resolves to an array of Balance objects — iterate it directly:

```ts
const balances = await client.Balance().list()

for (const balance of balances) {
  console.log(balance)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = MercadoBitcoinSDK.test()

const balance = await client.Balance().load({ id: 'test01' })
// balance is a bare entity populated with mock response data
console.log(balance)
```

You can also use the instance method:

```ts
const client = new MercadoBitcoinSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Balance()

// First call sets internal match
await entity.load({ id: 'example' })

// Subsequent calls reuse the stored match
const data = entity.data()
console.log(data.id) // 'example'
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new MercadoBitcoinSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```


## Reference

### MercadoBitcoinSDK

#### Constructor

```ts
new MercadoBitcoinSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Balance(data?)` | `BalanceEntity` | Create a Balance entity instance. |
| `Candle(data?)` | `CandleEntity` | Create a Candle entity instance. |
| `DepositAddress(data?)` | `DepositAddressEntity` | Create a DepositAddress entity instance. |
| `Order(data?)` | `OrderEntity` | Create an Order entity instance. |
| `OrderBook(data?)` | `OrderBookEntity` | Create an OrderBook entity instance. |
| `Ticker(data?)` | `TickerEntity` | Create a Ticker entity instance. |
| `Trade(data?)` | `TradeEntity` | Create a Trade entity instance. |
| `Withdrawal(data?)` | `WithdrawalEntity` | Create a Withdrawal entity instance. |
| `tester(testopts?, sdkopts?)` | `MercadoBitcoinSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `MercadoBitcoinSDK.test(testopts?, sdkopts?)` | `MercadoBitcoinSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?): any` | Get or set entity data. |
| `match` | `match(match?): any` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): MercadoBitcoinSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Balance

| Field | Description |
| --- | --- |
| `available` |  |
| `currency` |  |
| `locked` |  |
| `total` |  |

Operations: list.

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

Operations: load.

API path: `/candles/{symbol}`

#### DepositAddress

| Field | Description |
| --- | --- |
| `address` |  |
| `currency` |  |
| `qr_code` |  |
| `tag` |  |

Operations: load.

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

Operations: create, list, load, remove.

API path: `/orders`

#### OrderBook

| Field | Description |
| --- | --- |
| `ask` |  |
| `bid` |  |
| `timestamp` |  |

Operations: load.

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

Operations: list, load.

API path: `/tickers`

#### Trade

| Field | Description |
| --- | --- |
| `amount` |  |
| `id` |  |
| `price` |  |
| `side` |  |
| `timestamp` |  |

Operations: load.

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

Operations: create.

API path: `/withdrawals/brl`



## Entities


### Balance

Create an instance: `const balance = client.Balance()`

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

```ts
const balances = await client.Balance().list()
```


### Candle

Create an instance: `const candle = client.Candle()`

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

```ts
const candle = await client.Candle().load({ id: 'candle_id' })
```


### DepositAddress

Create an instance: `const deposit_address = client.DepositAddress()`

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

```ts
const deposit_address = await client.DepositAddress().load({ id: 'deposit_address_id' })
```


### Order

Create an instance: `const order = client.Order()`

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

```ts
const order = await client.Order().load({ id: 'order_id' })
```

#### Example: List

```ts
const orders = await client.Order().list()
```

#### Example: Create

```ts
const order = await client.Order().create({
})
```


### OrderBook

Create an instance: `const order_book = client.OrderBook()`

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

```ts
const order_book = await client.OrderBook().load({ id: 'order_book_id' })
```


### Ticker

Create an instance: `const ticker = client.Ticker()`

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

```ts
const ticker = await client.Ticker().load({ id: 'ticker_id' })
```

#### Example: List

```ts
const tickers = await client.Ticker().list()
```


### Trade

Create an instance: `const trade = client.Trade()`

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

```ts
const trade = await client.Trade().load({ id: 'trade_id' })
```


### Withdrawal

Create an instance: `const withdrawal = client.Withdrawal()`

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

```ts
const withdrawal = await client.Withdrawal().create({
  account_number: /* `$STRING` */,
  address: /* `$STRING` */,
  agency: /* `$STRING` */,
  amount: /* `$NUMBER` */,
  bank: /* `$STRING` */,
  currency: /* `$STRING` */,
})
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
error is returned to the caller.

An unexpected exception triggers the `PreUnexpected` hook before
propagating.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
mercado-bitcoin/
├── src/
│   ├── MercadoBitcoinSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { MercadoBitcoinSDK } from '@voxgig-sdk/mercado-bitcoin'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const balance = client.Balance()
await balance.load({ id: "example_id" })

// balance.data() now returns the loaded balance data
// balance.match() returns { id: "example_id" }
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.

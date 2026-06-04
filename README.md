# MercadoBitcoin SDK

Trade on Brazil's Mercado Bitcoin exchange: market data, orders, deposits, and withdrawals via a REST v4 API

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Mercado Bitcoin API

[Mercado Bitcoin](https://www.mercadobitcoin.com.br/) is one of Brazil's largest cryptocurrency exchanges. Its v4 REST API exposes public market data alongside authenticated endpoints for account management and trading.

What you get from the API:

- Public market data: trading symbols, order books, recent trades, tickers, and OHLC candles.
- Account operations: balances, deposit addresses, and withdrawals for cryptocurrencies and Brazilian Real (BRL).
- Trading: place, query, and cancel orders.

The API is served from `https://api.mercadobitcoin.net/api/v4`. CORS is disabled, so calls must come from a server-side client. Authentication and rate-limit specifics are documented on the official docs site.

## Try it

**TypeScript**
```bash
npm install mercado-bitcoin
```

**Python**
```bash
pip install mercado-bitcoin-sdk
```

**PHP**
```bash
composer require voxgig/mercado-bitcoin-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/mercado-bitcoin-sdk/go
```

**Ruby**
```bash
gem install mercado-bitcoin-sdk
```

**Lua**
```bash
luarocks install mercado-bitcoin-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { MercadoBitcoinSDK } from 'mercado-bitcoin'

const client = new MercadoBitcoinSDK({})

// List all balances
const balances = await client.Balance().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o mercado-bitcoin-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "mercado-bitcoin": {
      "command": "/abs/path/to/mercado-bitcoin-mcp"
    }
  }
}
```

## Entities

The API exposes 8 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Balance** | Account balance resource for the authenticated user's wallets across supported assets. | `/accounts/balance` |
| **Candle** | OHLC (open/high/low/close) candlestick market data for a trading symbol over a chosen interval. | `/candles/{symbol}` |
| **DepositAddress** | Deposit address resource used to fund an account with a given cryptocurrency. | `/deposits/crypto` |
| **Order** | Trading order resource for placing, retrieving, and cancelling buy/sell orders on a market. | `/orders` |
| **OrderBook** | Snapshot of current bids and asks for a trading symbol. | `/orderbook/{symbol}` |
| **Ticker** | Latest price summary (last, bid, ask, volume) for a trading symbol. | `/tickers` |
| **Trade** | Recent public trades executed on a market. | `/trades/{symbol}` |
| **Withdrawal** | Withdrawal request resource for moving crypto or BRL out of an account. | `/withdrawals/brl` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from mercadobitcoin_sdk import MercadoBitcoinSDK

client = MercadoBitcoinSDK({})

# List all balances
balances, err = client.Balance(None).list(None, None)
```

### PHP

```php
<?php
require_once 'mercadobitcoin_sdk.php';

$client = new MercadoBitcoinSDK([]);

// List all balances
[$balances, $err] = $client->Balance(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/mercado-bitcoin-sdk/go"

client := sdk.NewMercadoBitcoinSDK(map[string]any{})

// List all balances
balances, err := client.Balance(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "MercadoBitcoin_sdk"

client = MercadoBitcoinSDK.new({})

# List all balances
balances, err = client.Balance(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("mercado-bitcoin_sdk")

local client = sdk.new({})

-- List all balances
local balances, err = client:Balance(nil):list(nil, nil)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = MercadoBitcoinSDK.test()
const result = await client.Balance().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = MercadoBitcoinSDK.test(None, None)
result, err = client.Balance(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = MercadoBitcoinSDK::test(null, null);
[$result, $err] = $client->Balance(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Balance(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = MercadoBitcoinSDK.test(nil, nil)
result, err = client.Balance(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Balance(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Mercado Bitcoin API

- Upstream: [https://www.mercadobitcoin.com.br/](https://www.mercadobitcoin.com.br/)
- API docs: [https://api.mercadobitcoin.net/api/v4/docs](https://api.mercadobitcoin.net/api/v4/docs)

---

Generated from the Mercado Bitcoin API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

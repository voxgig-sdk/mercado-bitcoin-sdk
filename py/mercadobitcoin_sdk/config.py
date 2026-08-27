# MercadoBitcoin SDK configuration


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "MercadoBitcoin",
            "slug": "mercado-bitcoin",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "transport": "base",
      },
        },
        "options": {
            "base": "https://api.mercadobitcoin.net/api/v4",
            "auth": {
                "prefix": "",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "balance": {},
                "candle": {},
                "deposit_address": {},
                "order": {},
                "order_book": {},
                "ticker": {},
                "trade": {},
                "withdrawal": {},
            },
        },
        "entity": {
      "balance": {
        "fields": [
          {
            "name": "available",
            "short": "Available balance",
            "type": "`$NUMBER`",
          },
          {
            "name": "currency",
            "short": "Currency code",
            "type": "`$STRING`",
          },
          {
            "name": "locked",
            "short": "Locked balance",
            "type": "`$NUMBER`",
          },
          {
            "name": "total",
            "short": "Total balance",
            "type": "`$NUMBER`",
          },
        ],
        "name": "balance",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/accounts/balance",
                "parts": [
                  "accounts",
                  "balance",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.balances`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "candle": {
        "fields": [
          {
            "name": "close",
            "short": "Closing price",
            "type": "`$NUMBER`",
          },
          {
            "name": "high",
            "short": "Highest price",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "low",
            "short": "Lowest price",
            "type": "`$NUMBER`",
          },
          {
            "name": "open",
            "short": "Opening price",
            "type": "`$NUMBER`",
          },
          {
            "name": "timestamp",
            "short": "Candle timestamp in milliseconds",
            "type": "`$INTEGER`",
          },
          {
            "name": "volume",
            "short": "Trading volume",
            "type": "`$NUMBER`",
          },
        ],
        "name": "candle",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "symbol",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "example": "1h",
                      "kind": "query",
                      "name": "interval",
                      "orig": "interval",
                      "type": "`$STRING`",
                    },
                    {
                      "example": 100,
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/candles/{symbol}",
                "parts": [
                  "candles",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "symbol": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "interval",
                    "limit",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "deposit_address": {
        "fields": [
          {
            "name": "address",
            "short": "Deposit address",
            "type": "`$STRING`",
          },
          {
            "name": "currency",
            "short": "Cryptocurrency code",
            "type": "`$STRING`",
          },
          {
            "name": "qrCode",
            "short": "QR code for deposit address",
            "type": "`$STRING`",
          },
          {
            "name": "tag",
            "short": "Deposit tag/memo (if applicable)",
            "type": "`$STRING`",
          },
        ],
        "name": "deposit_address",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "currency",
                      "orig": "currency",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/deposits/crypto",
                "parts": [
                  "deposits",
                  "crypto",
                ],
                "select": {
                  "exist": [
                    "currency",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "order": {
        "fields": [
          {
            "name": "amount",
            "op": {
              "create": {
                "req": True,
                "type": "`$NUMBER`",
              },
            },
            "short": "Order amount",
            "type": "`$NUMBER`",
          },
          {
            "name": "filled",
            "short": "Filled amount",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "short": "Order ID",
            "type": "`$STRING`",
          },
          {
            "name": "price",
            "short": "Order price",
            "type": "`$NUMBER`",
          },
          {
            "name": "side",
            "op": {
              "create": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "short": "Order side",
            "type": "`$STRING`",
          },
          {
            "name": "status",
            "short": "Order status",
            "type": "`$STRING`",
          },
          {
            "name": "symbol",
            "op": {
              "create": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "short": "Trading pair symbol",
            "type": "`$STRING`",
          },
          {
            "name": "timestamp",
            "short": "Order creation timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "type",
            "op": {
              "create": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "short": "Order type",
            "type": "`$STRING`",
          },
        ],
        "name": "order",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/orders",
                "parts": [
                  "orders",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "symbol",
                      "orig": "symbol",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/orders",
                "parts": [
                  "orders",
                ],
                "select": {
                  "exist": [
                    "status",
                    "symbol",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "order_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/orders/{orderId}",
                "parts": [
                  "orders",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "orderId": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "order_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/orders/{orderId}",
                "parts": [
                  "orders",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "orderId": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "order_book": {
        "fields": [
          {
            "name": "asks",
            "short": "List of ask orders",
            "type": "`$ARRAY`",
          },
          {
            "name": "bids",
            "short": "List of bid orders",
            "type": "`$ARRAY`",
          },
          {
            "name": "timestamp",
            "short": "Timestamp in milliseconds",
            "type": "`$INTEGER`",
          },
        ],
        "name": "order_book",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "symbol",
                      "orig": "symbol",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "example": 100,
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/orderbook/{symbol}",
                "parts": [
                  "orderbook",
                  "{symbol}",
                ],
                "select": {
                  "exist": [
                    "limit",
                    "symbol",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "orderbook",
            ],
          ],
        },
      },
      "ticker": {
        "fields": [
          {
            "name": "ask",
            "short": "Lowest ask price",
            "type": "`$NUMBER`",
          },
          {
            "name": "bid",
            "short": "Highest bid price",
            "type": "`$NUMBER`",
          },
          {
            "name": "high",
            "short": "24h high price",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "last",
            "short": "Last traded price",
            "type": "`$NUMBER`",
          },
          {
            "name": "low",
            "short": "24h low price",
            "type": "`$NUMBER`",
          },
          {
            "name": "symbol",
            "short": "Trading pair symbol",
            "type": "`$STRING`",
          },
          {
            "name": "timestamp",
            "short": "Timestamp in milliseconds",
            "type": "`$INTEGER`",
          },
          {
            "name": "volume",
            "short": "24h trading volume",
            "type": "`$NUMBER`",
          },
        ],
        "name": "ticker",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/tickers",
                "parts": [
                  "tickers",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "symbol",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/tickers/{symbol}",
                "parts": [
                  "tickers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "symbol": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "trade": {
        "fields": [
          {
            "name": "amount",
            "short": "Trade amount",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "short": "Trade ID",
            "type": "`$STRING`",
          },
          {
            "name": "price",
            "short": "Trade price",
            "type": "`$NUMBER`",
          },
          {
            "name": "side",
            "short": "Trade side",
            "type": "`$STRING`",
          },
          {
            "name": "timestamp",
            "short": "Timestamp in milliseconds",
            "type": "`$INTEGER`",
          },
        ],
        "name": "trade",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "symbol",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "example": 100,
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/trades/{symbol}",
                "parts": [
                  "trades",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "symbol": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "limit",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "withdrawal": {
        "fields": [
          {
            "name": "accountNumber",
            "req": True,
            "short": "Bank account number",
            "type": "`$STRING`",
          },
          {
            "name": "accountType",
            "short": "Account type",
            "type": "`$STRING`",
          },
          {
            "name": "address",
            "req": True,
            "short": "Destination address",
            "type": "`$STRING`",
          },
          {
            "name": "agency",
            "req": True,
            "short": "Bank agency",
            "type": "`$STRING`",
          },
          {
            "name": "amount",
            "req": True,
            "short": "Withdrawal amount in BRL",
            "type": "`$NUMBER`",
          },
          {
            "name": "bank",
            "req": True,
            "short": "Bank code",
            "type": "`$STRING`",
          },
          {
            "name": "currency",
            "req": True,
            "short": "Cryptocurrency code",
            "type": "`$STRING`",
          },
          {
            "name": "tag",
            "short": "Destination tag/memo (if applicable)",
            "type": "`$STRING`",
          },
        ],
        "name": "withdrawal",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/withdrawals/brl",
                "parts": [
                  "withdrawals",
                  "brl",
                ],
                "select": {
                  "$action": "brl",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/withdrawals/crypto",
                "parts": [
                  "withdrawals",
                  "crypto",
                ],
                "select": {
                  "$action": "crypto",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }

package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "MercadoBitcoin",
			"slug": "mercado-bitcoin",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://api.mercadobitcoin.net/api/v4",
			"auth": map[string]any{
				"prefix": "",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"balance": map[string]any{},
				"candle": map[string]any{},
				"deposit_address": map[string]any{},
				"order": map[string]any{},
				"order_book": map[string]any{},
				"ticker": map[string]any{},
				"trade": map[string]any{},
				"withdrawal": map[string]any{},
			},
		},
		"entity": map[string]any{
			"balance": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "available",
						"short": "Available balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "locked",
						"short": "Locked balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "total",
						"short": "Total balance",
						"type": "`$NUMBER`",
					},
				},
				"name": "balance",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/accounts/balance",
								"segments": []any{
									map[string]any{
										"lit": "accounts",
									},
									map[string]any{
										"lit": "balance",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.balances`",
								},
								"parts": []any{
									"accounts",
									"balance",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"candle": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "close",
						"short": "Closing price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "high",
						"short": "Highest price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "low",
						"short": "Lowest price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "open",
						"short": "Opening price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Candle timestamp in milliseconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "volume",
						"short": "Trading volume",
						"type": "`$NUMBER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "candle",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "symbol",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": "1h",
											"kind": "query",
											"name": "interval",
											"orig": "interval",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/candles/{symbol}",
								"rename": map[string]any{
									"param": map[string]any{
										"symbol": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "candles",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"interval",
										"limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"candles",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"deposit_address": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "address",
						"short": "Deposit address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Cryptocurrency code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "qrCode",
						"short": "QR code for deposit address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tag",
						"short": "Deposit tag/memo (if applicable)",
						"type": "`$STRING`",
					},
				},
				"name": "deposit_address",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/deposits/crypto",
								"segments": []any{
									map[string]any{
										"lit": "deposits",
									},
									map[string]any{
										"lit": "crypto",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"currency",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"deposits",
									"crypto",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"order": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amount",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$NUMBER`",
							},
						},
						"short": "Order amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "filled",
						"short": "Filled amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"short": "Order ID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "price",
						"short": "Order price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "side",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "Order side",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "Order status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "symbol",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "Trading pair symbol",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Order creation timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "Order type",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "order",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/orders",
								"segments": []any{
									map[string]any{
										"lit": "orders",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"orders",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "symbol",
											"orig": "symbol",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/orders",
								"segments": []any{
									map[string]any{
										"lit": "orders",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"status",
										"symbol",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"orders",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "order_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/orders/{orderId}",
								"rename": map[string]any{
									"param": map[string]any{
										"orderId": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "orders",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"orders",
									"{id}",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "order_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/orders/{orderId}",
								"rename": map[string]any{
									"param": map[string]any{
										"orderId": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "orders",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"orders",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"order_book": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "asks",
						"short": "List of ask orders",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "bids",
						"short": "List of bid orders",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Timestamp in milliseconds",
						"type": "`$INTEGER`",
					},
				},
				"name": "order_book",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "symbol",
											"orig": "symbol",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/orderbook/{symbol}",
								"segments": []any{
									map[string]any{
										"lit": "orderbook",
									},
									map[string]any{
										"var": "symbol",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"symbol",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"orderbook",
									"{symbol}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"orderbook",
						},
					},
				},
			},
			"ticker": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ask",
						"short": "Lowest ask price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "bid",
						"short": "Highest bid price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "high",
						"short": "24h high price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last",
						"short": "Last traded price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "low",
						"short": "24h low price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "symbol",
						"short": "Trading pair symbol",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Timestamp in milliseconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "volume",
						"short": "24h trading volume",
						"type": "`$NUMBER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "ticker",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/tickers",
								"segments": []any{
									map[string]any{
										"lit": "tickers",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"tickers",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "symbol",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tickers/{symbol}",
								"rename": map[string]any{
									"param": map[string]any{
										"symbol": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "tickers",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"tickers",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"trade": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amount",
						"short": "Trade amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"short": "Trade ID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "price",
						"short": "Trade price",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "side",
						"short": "Trade side",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Timestamp in milliseconds",
						"type": "`$INTEGER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "trade",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "symbol",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/trades/{symbol}",
								"rename": map[string]any{
									"param": map[string]any{
										"symbol": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "trades",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"trades",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"withdrawal": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountNumber",
						"req": true,
						"short": "Bank account number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountType",
						"short": "Account type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "address",
						"req": true,
						"short": "Destination address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "agency",
						"req": true,
						"short": "Bank agency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "amount",
						"req": true,
						"short": "Withdrawal amount in BRL",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "bank",
						"req": true,
						"short": "Bank code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"req": true,
						"short": "Cryptocurrency code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tag",
						"short": "Destination tag/memo (if applicable)",
						"type": "`$STRING`",
					},
				},
				"name": "withdrawal",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/withdrawals/brl",
								"segments": []any{
									map[string]any{
										"lit": "withdrawals",
									},
									map[string]any{
										"lit": "brl",
									},
								},
								"select": map[string]any{
									"$action": "brl",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"withdrawals",
									"brl",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/withdrawals/crypto",
								"segments": []any{
									map[string]any{
										"lit": "withdrawals",
									},
									map[string]any{
										"lit": "crypto",
									},
								},
								"select": map[string]any{
									"$action": "crypto",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"withdrawals",
									"crypto",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}

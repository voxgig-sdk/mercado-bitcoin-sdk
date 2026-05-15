package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "MercadoBitcoin",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.mercadobitcoin.net/api/v4",
			"auth": map[string]any{
				"prefix": "Bearer",
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
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "currency",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "locked",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "total",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "balance",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"method": "GET",
								"orig": "/accounts/balance",
								"parts": []any{
									"accounts",
									"balance",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "high",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "low",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "open",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "timestamp",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "volume",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 5,
					},
				},
				"name": "candle",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
									"query": []any{
										map[string]any{
											"example": "1h",
											"kind": "query",
											"name": "interval",
											"orig": "interval",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/candles/{symbol}",
								"parts": []any{
									"candles",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"symbol": "id",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "currency",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "qr_code",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "tag",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "deposit_address",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/deposits/crypto",
								"parts": []any{
									"deposits",
									"crypto",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "filled",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "price",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "side",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "status",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "symbol",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "timestamp",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 7,
					},
					map[string]any{
						"name": "type",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 8,
					},
				},
				"name": "order",
				"op": map[string]any{
					"create": map[string]any{
						"name": "create",
						"points": []any{
							map[string]any{
								"method": "POST",
								"orig": "/orders",
								"parts": []any{
									"orders",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "create",
					},
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "symbol",
											"orig": "symbol",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/orders",
								"parts": []any{
									"orders",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/orders/{orderId}",
								"parts": []any{
									"orders",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"orderId": "id",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
					},
					"remove": map[string]any{
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
											"active": true,
										},
									},
								},
								"method": "DELETE",
								"orig": "/orders/{orderId}",
								"parts": []any{
									"orders",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"orderId": "id",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "remove",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"order_book": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ask",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "bid",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "timestamp",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
				},
				"name": "order_book",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
									"query": []any{
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/orderbook/{symbol}",
								"parts": []any{
									"orderbook",
									"{symbol}",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "bid",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "high",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "last",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "low",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "symbol",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "timestamp",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "volume",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 7,
					},
				},
				"name": "ticker",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"method": "GET",
								"orig": "/tickers",
								"parts": []any{
									"tickers",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/tickers/{symbol}",
								"parts": []any{
									"tickers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"symbol": "id",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "price",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "side",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "timestamp",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 4,
					},
				},
				"name": "trade",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
									"query": []any{
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
									},
								},
								"method": "GET",
								"orig": "/trades/{symbol}",
								"parts": []any{
									"trades",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"symbol": "id",
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"withdrawal": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "account_number",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "account_type",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "address",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "agency",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "amount",
						"req": true,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "bank",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "currency",
						"req": true,
						"type": "`$STRING`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "tag",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 7,
					},
				},
				"name": "withdrawal",
				"op": map[string]any{
					"create": map[string]any{
						"name": "create",
						"points": []any{
							map[string]any{
								"method": "POST",
								"orig": "/withdrawals/brl",
								"parts": []any{
									"withdrawals",
									"brl",
								},
								"select": map[string]any{
									"$action": "brl",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"index$": 0,
							},
							map[string]any{
								"method": "POST",
								"orig": "/withdrawals/crypto",
								"parts": []any{
									"withdrawals",
									"crypto",
								},
								"select": map[string]any{
									"$action": "crypto",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"index$": 1,
							},
						},
						"input": "data",
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
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

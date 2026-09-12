<?php
declare(strict_types=1);

// MercadoBitcoin SDK configuration

class MercadoBitcoinConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "MercadoBitcoin",
                "slug" => "mercado-bitcoin",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
          'transport' => 'base',
        ],
            ],
            "options" => [
                "base" => "https://api.mercadobitcoin.net/api/v4",
                "auth" => [
                    "prefix" => "",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "balance" => [],
                    "candle" => [],
                    "deposit_address" => [],
                    "order" => [],
                    "order_book" => [],
                    "ticker" => [],
                    "trade" => [],
                    "withdrawal" => [],
                ],
            ],
            "entity" => [
        'balance' => [
          'fields' => [
            [
              'name' => 'available',
              'short' => 'Available balance',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'currency',
              'short' => 'Currency code',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'locked',
              'short' => 'Locked balance',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'total',
              'short' => 'Total balance',
              'type' => '`$NUMBER`',
            ],
          ],
          'name' => 'balance',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/accounts/balance',
                  'segments' => [
                    [
                      'lit' => 'accounts',
                    ],
                    [
                      'lit' => 'balance',
                    ],
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.balances`',
                  ],
                  'parts' => [
                    'accounts',
                    'balance',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'candle' => [
          'fields' => [
            [
              'name' => 'close',
              'short' => 'Closing price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'high',
              'short' => 'Highest price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'low',
              'short' => 'Lowest price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'open',
              'short' => 'Opening price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'timestamp',
              'short' => 'Candle timestamp in milliseconds',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'volume',
              'short' => 'Trading volume',
              'type' => '`$NUMBER`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'candle',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'symbol',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => '1h',
                        'kind' => 'query',
                        'name' => 'interval',
                        'orig' => 'interval',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/candles/{symbol}',
                  'rename' => [
                    'param' => [
                      'symbol' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'candles',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'interval',
                      'limit',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'candles',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'deposit_address' => [
          'fields' => [
            [
              'name' => 'address',
              'short' => 'Deposit address',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'currency',
              'short' => 'Cryptocurrency code',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'qrCode',
              'short' => 'QR code for deposit address',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'tag',
              'short' => 'Deposit tag/memo (if applicable)',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'deposit_address',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'currency',
                        'orig' => 'currency',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/deposits/crypto',
                  'segments' => [
                    [
                      'lit' => 'deposits',
                    ],
                    [
                      'lit' => 'crypto',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'currency',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'deposits',
                    'crypto',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'order' => [
          'fields' => [
            [
              'name' => 'amount',
              'op' => [
                'create' => [
                  'req' => true,
                  'type' => '`$NUMBER`',
                ],
              ],
              'short' => 'Order amount',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'filled',
              'short' => 'Filled amount',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'id',
              'short' => 'Order ID',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'price',
              'short' => 'Order price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'side',
              'op' => [
                'create' => [
                  'req' => true,
                  'type' => '`$STRING`',
                ],
              ],
              'short' => 'Order side',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status',
              'short' => 'Order status',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'symbol',
              'op' => [
                'create' => [
                  'req' => true,
                  'type' => '`$STRING`',
                ],
              ],
              'short' => 'Trading pair symbol',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'timestamp',
              'short' => 'Order creation timestamp',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'type',
              'op' => [
                'create' => [
                  'req' => true,
                  'type' => '`$STRING`',
                ],
              ],
              'short' => 'Order type',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'order',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/orders',
                  'segments' => [
                    [
                      'lit' => 'orders',
                    ],
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'orders',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'status',
                        'orig' => 'status',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'symbol',
                        'orig' => 'symbol',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/orders',
                  'segments' => [
                    [
                      'lit' => 'orders',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'status',
                      'symbol',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'orders',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'order_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/orders/{orderId}',
                  'rename' => [
                    'param' => [
                      'orderId' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'orders',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'orders',
                    '{id}',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'order_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/orders/{orderId}',
                  'rename' => [
                    'param' => [
                      'orderId' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'orders',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'orders',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'order_book' => [
          'fields' => [
            [
              'name' => 'asks',
              'short' => 'List of ask orders',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'bids',
              'short' => 'List of bid orders',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'timestamp',
              'short' => 'Timestamp in milliseconds',
              'type' => '`$INTEGER`',
            ],
          ],
          'name' => 'order_book',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'symbol',
                        'orig' => 'symbol',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/orderbook/{symbol}',
                  'segments' => [
                    [
                      'lit' => 'orderbook',
                    ],
                    [
                      'var' => 'symbol',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'limit',
                      'symbol',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'orderbook',
                    '{symbol}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'orderbook',
              ],
            ],
          ],
        ],
        'ticker' => [
          'fields' => [
            [
              'name' => 'ask',
              'short' => 'Lowest ask price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'bid',
              'short' => 'Highest bid price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'high',
              'short' => '24h high price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last',
              'short' => 'Last traded price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'low',
              'short' => '24h low price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'symbol',
              'short' => 'Trading pair symbol',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'timestamp',
              'short' => 'Timestamp in milliseconds',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'volume',
              'short' => '24h trading volume',
              'type' => '`$NUMBER`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'ticker',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/tickers',
                  'segments' => [
                    [
                      'lit' => 'tickers',
                    ],
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'tickers',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'symbol',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/tickers/{symbol}',
                  'rename' => [
                    'param' => [
                      'symbol' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'tickers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'tickers',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'trade' => [
          'fields' => [
            [
              'name' => 'amount',
              'short' => 'Trade amount',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'id',
              'short' => 'Trade ID',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'price',
              'short' => 'Trade price',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'side',
              'short' => 'Trade side',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'timestamp',
              'short' => 'Timestamp in milliseconds',
              'type' => '`$INTEGER`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'trade',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'symbol',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/trades/{symbol}',
                  'rename' => [
                    'param' => [
                      'symbol' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'trades',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'limit',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'trades',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'withdrawal' => [
          'fields' => [
            [
              'name' => 'accountNumber',
              'req' => true,
              'short' => 'Bank account number',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'accountType',
              'short' => 'Account type',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'address',
              'req' => true,
              'short' => 'Destination address',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'agency',
              'req' => true,
              'short' => 'Bank agency',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'amount',
              'req' => true,
              'short' => 'Withdrawal amount in BRL',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'bank',
              'req' => true,
              'short' => 'Bank code',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'currency',
              'req' => true,
              'short' => 'Cryptocurrency code',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'tag',
              'short' => 'Destination tag/memo (if applicable)',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'withdrawal',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/withdrawals/brl',
                  'segments' => [
                    [
                      'lit' => 'withdrawals',
                    ],
                    [
                      'lit' => 'brl',
                    ],
                  ],
                  'select' => [
                    '$action' => 'brl',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'withdrawals',
                    'brl',
                  ],
                ],
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/withdrawals/crypto',
                  'segments' => [
                    [
                      'lit' => 'withdrawals',
                    ],
                    [
                      'lit' => 'crypto',
                    ],
                  ],
                  'select' => [
                    '$action' => 'crypto',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'withdrawals',
                    'crypto',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return MercadoBitcoinFeatures::make_feature($name);
    }
}

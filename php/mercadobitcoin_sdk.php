<?php
declare(strict_types=1);

// MercadoBitcoin SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

class MercadoBitcoinSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new MercadoBitcoinUtility();
        $this->_utility = $utility;

        $config = MercadoBitcoinConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features from config.
        $feature_opts = MercadoBitcoinHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $items = Struct::items($feature_opts);
            if ($items) {
                foreach ($items as $item) {
                    $fname = $item[0];
                    $fopts = MercadoBitcoinHelpers::to_map($item[1]);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, MercadoBitcoinFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return MercadoBitcoinUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = MercadoBitcoinHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = MercadoBitcoinHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = MercadoBitcoinHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new MercadoBitcoinSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    public function direct(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = MercadoBitcoinHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = MercadoBitcoinHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }


    private $_balance = null;

    // Canonical facade: $client->Balance()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->balance()
    // resolves here too.
    public function Balance($data = null)
    {
        require_once __DIR__ . '/entity/balance_entity.php';
        if ($data === null) {
            if ($this->_balance === null) {
                $this->_balance = new BalanceEntity($this, null);
            }
            return $this->_balance;
        }
        return new BalanceEntity($this, $data);
    }


    private $_candle = null;

    // Canonical facade: $client->Candle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->candle()
    // resolves here too.
    public function Candle($data = null)
    {
        require_once __DIR__ . '/entity/candle_entity.php';
        if ($data === null) {
            if ($this->_candle === null) {
                $this->_candle = new CandleEntity($this, null);
            }
            return $this->_candle;
        }
        return new CandleEntity($this, $data);
    }


    private $_deposit_address = null;

    // Canonical facade: $client->DepositAddress()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->deposit_address()
    // resolves here too.
    public function DepositAddress($data = null)
    {
        require_once __DIR__ . '/entity/deposit_address_entity.php';
        if ($data === null) {
            if ($this->_deposit_address === null) {
                $this->_deposit_address = new DepositAddressEntity($this, null);
            }
            return $this->_deposit_address;
        }
        return new DepositAddressEntity($this, $data);
    }


    private $_order = null;

    // Canonical facade: $client->Order()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->order()
    // resolves here too.
    public function Order($data = null)
    {
        require_once __DIR__ . '/entity/order_entity.php';
        if ($data === null) {
            if ($this->_order === null) {
                $this->_order = new OrderEntity($this, null);
            }
            return $this->_order;
        }
        return new OrderEntity($this, $data);
    }


    private $_order_book = null;

    // Canonical facade: $client->OrderBook()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->order_book()
    // resolves here too.
    public function OrderBook($data = null)
    {
        require_once __DIR__ . '/entity/order_book_entity.php';
        if ($data === null) {
            if ($this->_order_book === null) {
                $this->_order_book = new OrderBookEntity($this, null);
            }
            return $this->_order_book;
        }
        return new OrderBookEntity($this, $data);
    }


    private $_ticker = null;

    // Canonical facade: $client->Ticker()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->ticker()
    // resolves here too.
    public function Ticker($data = null)
    {
        require_once __DIR__ . '/entity/ticker_entity.php';
        if ($data === null) {
            if ($this->_ticker === null) {
                $this->_ticker = new TickerEntity($this, null);
            }
            return $this->_ticker;
        }
        return new TickerEntity($this, $data);
    }


    private $_trade = null;

    // Canonical facade: $client->Trade()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->trade()
    // resolves here too.
    public function Trade($data = null)
    {
        require_once __DIR__ . '/entity/trade_entity.php';
        if ($data === null) {
            if ($this->_trade === null) {
                $this->_trade = new TradeEntity($this, null);
            }
            return $this->_trade;
        }
        return new TradeEntity($this, $data);
    }


    private $_withdrawal = null;

    // Canonical facade: $client->Withdrawal()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->withdrawal()
    // resolves here too.
    public function Withdrawal($data = null)
    {
        require_once __DIR__ . '/entity/withdrawal_entity.php';
        if ($data === null) {
            if ($this->_withdrawal === null) {
                $this->_withdrawal = new WithdrawalEntity($this, null);
            }
            return $this->_withdrawal;
        }
        return new WithdrawalEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new MercadoBitcoinSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}

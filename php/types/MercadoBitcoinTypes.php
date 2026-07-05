<?php
declare(strict_types=1);

// Typed models for the MercadoBitcoin SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Balance entity data model. */
class Balance
{
    public ?float $available = null;
    public ?string $currency = null;
    public ?float $locked = null;
    public ?float $total = null;
}

/** Request payload for Balance#list. */
class BalanceListMatch
{
    public ?float $available = null;
    public ?string $currency = null;
    public ?float $locked = null;
    public ?float $total = null;
}

/** Candle entity data model. */
class Candle
{
    public ?float $close = null;
    public ?float $high = null;
    public ?float $low = null;
    public ?float $open = null;
    public ?int $timestamp = null;
    public ?float $volume = null;
}

/** Request payload for Candle#load. */
class CandleLoadMatch
{
    public string $id;
}

/** DepositAddress entity data model. */
class DepositAddress
{
    public ?string $address = null;
    public ?string $currency = null;
    public ?string $qr_code = null;
    public ?string $tag = null;
}

/** Request payload for DepositAddress#load. */
class DepositAddressLoadMatch
{
    public ?string $address = null;
    public ?string $currency = null;
    public ?string $qr_code = null;
    public ?string $tag = null;
}

/** Order entity data model. */
class Order
{
    public ?float $amount = null;
    public ?float $filled = null;
    public ?string $id = null;
    public ?float $price = null;
    public ?string $side = null;
    public ?string $status = null;
    public ?string $symbol = null;
    public ?int $timestamp = null;
    public ?string $type = null;
}

/** Request payload for Order#load. */
class OrderLoadMatch
{
    public string $id;
}

/** Request payload for Order#list. */
class OrderListMatch
{
    public ?float $amount = null;
    public ?float $filled = null;
    public ?string $id = null;
    public ?float $price = null;
    public ?string $side = null;
    public ?string $status = null;
    public ?string $symbol = null;
    public ?int $timestamp = null;
    public ?string $type = null;
}

/** Request payload for Order#create. */
class OrderCreateData
{
    public ?float $amount = null;
    public ?float $filled = null;
    public ?string $id = null;
    public ?float $price = null;
    public ?string $side = null;
    public ?string $status = null;
    public ?string $symbol = null;
    public ?int $timestamp = null;
    public ?string $type = null;
}

/** Request payload for Order#remove. */
class OrderRemoveMatch
{
    public string $id;
}

/** OrderBook entity data model. */
class OrderBook
{
    public ?array $ask = null;
    public ?array $bid = null;
    public ?int $timestamp = null;
}

/** Request payload for OrderBook#load. */
class OrderBookLoadMatch
{
    public string $symbol;
}

/** Ticker entity data model. */
class Ticker
{
    public ?float $ask = null;
    public ?float $bid = null;
    public ?float $high = null;
    public ?float $last = null;
    public ?float $low = null;
    public ?string $symbol = null;
    public ?int $timestamp = null;
    public ?float $volume = null;
}

/** Request payload for Ticker#load. */
class TickerLoadMatch
{
    public string $id;
}

/** Request payload for Ticker#list. */
class TickerListMatch
{
    public ?float $ask = null;
    public ?float $bid = null;
    public ?float $high = null;
    public ?float $last = null;
    public ?float $low = null;
    public ?string $symbol = null;
    public ?int $timestamp = null;
    public ?float $volume = null;
}

/** Trade entity data model. */
class Trade
{
    public ?float $amount = null;
    public ?string $id = null;
    public ?float $price = null;
    public ?string $side = null;
    public ?int $timestamp = null;
}

/** Request payload for Trade#load. */
class TradeLoadMatch
{
    public string $id;
}

/** Withdrawal entity data model. */
class Withdrawal
{
    public string $account_number;
    public ?string $account_type = null;
    public string $address;
    public string $agency;
    public float $amount;
    public string $bank;
    public string $currency;
    public ?string $tag = null;
}

/** Request payload for Withdrawal#create. */
class WithdrawalCreateData
{
    public string $account_number;
    public ?string $account_type = null;
    public string $address;
    public string $agency;
    public float $amount;
    public string $bank;
    public string $currency;
    public ?string $tag = null;
}


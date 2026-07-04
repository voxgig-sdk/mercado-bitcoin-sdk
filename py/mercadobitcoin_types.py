# Typed models for the MercadoBitcoin SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Balance:
    available: Optional[float] = None
    currency: Optional[str] = None
    locked: Optional[float] = None
    total: Optional[float] = None


@dataclass
class BalanceListMatch:
    available: Optional[float] = None
    currency: Optional[str] = None
    locked: Optional[float] = None
    total: Optional[float] = None


@dataclass
class Candle:
    close: Optional[float] = None
    high: Optional[float] = None
    low: Optional[float] = None
    open: Optional[float] = None
    timestamp: Optional[int] = None
    volume: Optional[float] = None


@dataclass
class CandleLoadMatch:
    id: str


@dataclass
class DepositAddress:
    address: Optional[str] = None
    currency: Optional[str] = None
    qr_code: Optional[str] = None
    tag: Optional[str] = None


@dataclass
class DepositAddressLoadMatch:
    address: Optional[str] = None
    currency: Optional[str] = None
    qr_code: Optional[str] = None
    tag: Optional[str] = None


@dataclass
class Order:
    amount: Optional[float] = None
    filled: Optional[float] = None
    id: Optional[str] = None
    price: Optional[float] = None
    side: Optional[str] = None
    status: Optional[str] = None
    symbol: Optional[str] = None
    timestamp: Optional[int] = None
    type: Optional[str] = None


@dataclass
class OrderLoadMatch:
    id: str


@dataclass
class OrderListMatch:
    amount: Optional[float] = None
    filled: Optional[float] = None
    id: Optional[str] = None
    price: Optional[float] = None
    side: Optional[str] = None
    status: Optional[str] = None
    symbol: Optional[str] = None
    timestamp: Optional[int] = None
    type: Optional[str] = None


@dataclass
class OrderCreateData:
    amount: Optional[float] = None
    filled: Optional[float] = None
    id: Optional[str] = None
    price: Optional[float] = None
    side: Optional[str] = None
    status: Optional[str] = None
    symbol: Optional[str] = None
    timestamp: Optional[int] = None
    type: Optional[str] = None


@dataclass
class OrderRemoveMatch:
    id: str


@dataclass
class OrderBook:
    ask: Optional[list] = None
    bid: Optional[list] = None
    timestamp: Optional[int] = None


@dataclass
class OrderBookLoadMatch:
    symbol: str


@dataclass
class Ticker:
    ask: Optional[float] = None
    bid: Optional[float] = None
    high: Optional[float] = None
    last: Optional[float] = None
    low: Optional[float] = None
    symbol: Optional[str] = None
    timestamp: Optional[int] = None
    volume: Optional[float] = None


@dataclass
class TickerLoadMatch:
    id: str


@dataclass
class TickerListMatch:
    ask: Optional[float] = None
    bid: Optional[float] = None
    high: Optional[float] = None
    last: Optional[float] = None
    low: Optional[float] = None
    symbol: Optional[str] = None
    timestamp: Optional[int] = None
    volume: Optional[float] = None


@dataclass
class Trade:
    amount: Optional[float] = None
    id: Optional[str] = None
    price: Optional[float] = None
    side: Optional[str] = None
    timestamp: Optional[int] = None


@dataclass
class TradeLoadMatch:
    id: str


@dataclass
class Withdrawal:
    account_number: str
    address: str
    agency: str
    amount: float
    bank: str
    currency: str
    account_type: Optional[str] = None
    tag: Optional[str] = None


@dataclass
class WithdrawalCreateData:
    account_number: Optional[str] = None
    account_type: Optional[str] = None
    address: Optional[str] = None
    agency: Optional[str] = None
    amount: Optional[float] = None
    bank: Optional[str] = None
    currency: Optional[str] = None
    tag: Optional[str] = None


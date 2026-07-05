# Typed models for the MercadoBitcoin SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Balance(TypedDict, total=False):
    available: float
    currency: str
    locked: float
    total: float


class BalanceListMatch(TypedDict, total=False):
    available: float
    currency: str
    locked: float
    total: float


class Candle(TypedDict, total=False):
    close: float
    high: float
    low: float
    open: float
    timestamp: int
    volume: float


class CandleLoadMatch(TypedDict):
    id: str


class DepositAddress(TypedDict, total=False):
    address: str
    currency: str
    qr_code: str
    tag: str


class DepositAddressLoadMatch(TypedDict, total=False):
    address: str
    currency: str
    qr_code: str
    tag: str


class Order(TypedDict, total=False):
    amount: float
    filled: float
    id: str
    price: float
    side: str
    status: str
    symbol: str
    timestamp: int
    type: str


class OrderLoadMatch(TypedDict):
    id: str


class OrderListMatch(TypedDict, total=False):
    amount: float
    filled: float
    id: str
    price: float
    side: str
    status: str
    symbol: str
    timestamp: int
    type: str


class OrderCreateData(TypedDict, total=False):
    amount: float
    filled: float
    id: str
    price: float
    side: str
    status: str
    symbol: str
    timestamp: int
    type: str


class OrderRemoveMatch(TypedDict):
    id: str


class OrderBook(TypedDict, total=False):
    ask: list
    bid: list
    timestamp: int


class OrderBookLoadMatch(TypedDict):
    symbol: str


class Ticker(TypedDict, total=False):
    ask: float
    bid: float
    high: float
    last: float
    low: float
    symbol: str
    timestamp: int
    volume: float


class TickerLoadMatch(TypedDict):
    id: str


class TickerListMatch(TypedDict, total=False):
    ask: float
    bid: float
    high: float
    last: float
    low: float
    symbol: str
    timestamp: int
    volume: float


class Trade(TypedDict, total=False):
    amount: float
    id: str
    price: float
    side: str
    timestamp: int


class TradeLoadMatch(TypedDict):
    id: str


class WithdrawalRequired(TypedDict):
    account_number: str
    address: str
    agency: str
    amount: float
    bank: str
    currency: str


class Withdrawal(WithdrawalRequired, total=False):
    account_type: str
    tag: str


class WithdrawalCreateDataRequired(TypedDict):
    account_number: str
    address: str
    agency: str
    amount: float
    bank: str
    currency: str


class WithdrawalCreateData(WithdrawalCreateDataRequired, total=False):
    account_type: str
    tag: str

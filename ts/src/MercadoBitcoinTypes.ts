// Typed models for the MercadoBitcoin SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Balance {
  available?: number
  currency?: string
  locked?: number
  total?: number
}

export type BalanceListMatch = Partial<Balance>

export interface Candle {
  close?: number
  high?: number
  low?: number
  open?: number
  timestamp?: number
  volume?: number
}

export interface CandleLoadMatch {
  id: string
}

export interface DepositAddress {
  address?: string
  currency?: string
  qr_code?: string
  tag?: string
}

export type DepositAddressLoadMatch = Partial<DepositAddress>

export interface Order {
  amount?: number
  filled?: number
  id?: string
  price?: number
  side?: string
  status?: string
  symbol?: string
  timestamp?: number
  type?: string
}

export interface OrderLoadMatch {
  id: string
}

export type OrderListMatch = Partial<Order>

export type OrderCreateData = Partial<Order>

export interface OrderRemoveMatch {
  id: string
}

export interface OrderBook {
  ask?: any[]
  bid?: any[]
  timestamp?: number
}

export interface OrderBookLoadMatch {
  symbol: string
}

export interface Ticker {
  ask?: number
  bid?: number
  high?: number
  last?: number
  low?: number
  symbol?: string
  timestamp?: number
  volume?: number
}

export interface TickerLoadMatch {
  id: string
}

export type TickerListMatch = Partial<Ticker>

export interface Trade {
  amount?: number
  id?: string
  price?: number
  side?: string
  timestamp?: number
}

export interface TradeLoadMatch {
  id: string
}

export interface Withdrawal {
  account_number: string
  account_type?: string
  address: string
  agency: string
  amount: number
  bank: string
  currency: string
  tag?: string
}

export type WithdrawalCreateData = Partial<Withdrawal>


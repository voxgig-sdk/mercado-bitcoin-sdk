-- Typed models for the MercadoBitcoin SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Balance
---@field available? number
---@field currency? string
---@field locked? number
---@field total? number

---@class BalanceListMatch

---@class Candle
---@field close? number
---@field high? number
---@field low? number
---@field open? number
---@field timestamp? number
---@field volume? number

---@class CandleLoadMatch
---@field id string

---@class DepositAddress
---@field address? string
---@field currency? string
---@field qr_code? string
---@field tag? string

---@class DepositAddressLoadMatch

---@class Order
---@field amount? number
---@field filled? number
---@field id? string
---@field price? number
---@field side? string
---@field status? string
---@field symbol? string
---@field timestamp? number
---@field type? string

---@class OrderLoadMatch
---@field id string

---@class OrderListMatch

---@class OrderCreateData

---@class OrderRemoveMatch
---@field id string

---@class OrderBook
---@field ask? table
---@field bid? table
---@field timestamp? number

---@class OrderBookLoadMatch
---@field symbol string

---@class Ticker
---@field ask? number
---@field bid? number
---@field high? number
---@field last? number
---@field low? number
---@field symbol? string
---@field timestamp? number
---@field volume? number

---@class TickerLoadMatch
---@field id string

---@class TickerListMatch

---@class Trade
---@field amount? number
---@field id? string
---@field price? number
---@field side? string
---@field timestamp? number

---@class TradeLoadMatch
---@field id string

---@class Withdrawal
---@field account_number string
---@field account_type? string
---@field address string
---@field agency string
---@field amount number
---@field bank string
---@field currency string
---@field tag? string

---@class WithdrawalCreateData

local M = {}

return M

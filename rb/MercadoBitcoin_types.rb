# frozen_string_literal: true

# Typed models for the MercadoBitcoin SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Balance entity data model.
#
# @!attribute [rw] available
#   @return [Float, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] locked
#   @return [Float, nil]
#
# @!attribute [rw] total
#   @return [Float, nil]
Balance = Struct.new(
  :available,
  :currency,
  :locked,
  :total,
  keyword_init: true
)

# Match filter for Balance#list (any subset of Balance fields).
#
# @!attribute [rw] available
#   @return [Float, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] locked
#   @return [Float, nil]
#
# @!attribute [rw] total
#   @return [Float, nil]
BalanceListMatch = Struct.new(
  :available,
  :currency,
  :locked,
  :total,
  keyword_init: true
)

# Candle entity data model.
#
# @!attribute [rw] close
#   @return [Float, nil]
#
# @!attribute [rw] high
#   @return [Float, nil]
#
# @!attribute [rw] low
#   @return [Float, nil]
#
# @!attribute [rw] open
#   @return [Float, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
#
# @!attribute [rw] volume
#   @return [Float, nil]
Candle = Struct.new(
  :close,
  :high,
  :low,
  :open,
  :timestamp,
  :volume,
  keyword_init: true
)

# Request payload for Candle#load.
#
# @!attribute [rw] id
#   @return [String]
CandleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# DepositAddress entity data model.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] qr_code
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [String, nil]
DepositAddress = Struct.new(
  :address,
  :currency,
  :qr_code,
  :tag,
  keyword_init: true
)

# Match filter for DepositAddress#load (any subset of DepositAddress fields).
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] qr_code
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [String, nil]
DepositAddressLoadMatch = Struct.new(
  :address,
  :currency,
  :qr_code,
  :tag,
  keyword_init: true
)

# Order entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] filled
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] price
#   @return [Float, nil]
#
# @!attribute [rw] side
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
Order = Struct.new(
  :amount,
  :filled,
  :id,
  :price,
  :side,
  :status,
  :symbol,
  :timestamp,
  :type,
  keyword_init: true
)

# Request payload for Order#load.
#
# @!attribute [rw] id
#   @return [String]
OrderLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Order#list (any subset of Order fields).
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] filled
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] price
#   @return [Float, nil]
#
# @!attribute [rw] side
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
OrderListMatch = Struct.new(
  :amount,
  :filled,
  :id,
  :price,
  :side,
  :status,
  :symbol,
  :timestamp,
  :type,
  keyword_init: true
)

# Match filter for Order#create (any subset of Order fields).
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] filled
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] price
#   @return [Float, nil]
#
# @!attribute [rw] side
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
OrderCreateData = Struct.new(
  :amount,
  :filled,
  :id,
  :price,
  :side,
  :status,
  :symbol,
  :timestamp,
  :type,
  keyword_init: true
)

# Request payload for Order#remove.
#
# @!attribute [rw] id
#   @return [String]
OrderRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# OrderBook entity data model.
#
# @!attribute [rw] ask
#   @return [Array, nil]
#
# @!attribute [rw] bid
#   @return [Array, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
OrderBook = Struct.new(
  :ask,
  :bid,
  :timestamp,
  keyword_init: true
)

# Request payload for OrderBook#load.
#
# @!attribute [rw] symbol
#   @return [String]
OrderBookLoadMatch = Struct.new(
  :symbol,
  keyword_init: true
)

# Ticker entity data model.
#
# @!attribute [rw] ask
#   @return [Float, nil]
#
# @!attribute [rw] bid
#   @return [Float, nil]
#
# @!attribute [rw] high
#   @return [Float, nil]
#
# @!attribute [rw] last
#   @return [Float, nil]
#
# @!attribute [rw] low
#   @return [Float, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
#
# @!attribute [rw] volume
#   @return [Float, nil]
Ticker = Struct.new(
  :ask,
  :bid,
  :high,
  :last,
  :low,
  :symbol,
  :timestamp,
  :volume,
  keyword_init: true
)

# Request payload for Ticker#load.
#
# @!attribute [rw] id
#   @return [String]
TickerLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Ticker#list (any subset of Ticker fields).
#
# @!attribute [rw] ask
#   @return [Float, nil]
#
# @!attribute [rw] bid
#   @return [Float, nil]
#
# @!attribute [rw] high
#   @return [Float, nil]
#
# @!attribute [rw] last
#   @return [Float, nil]
#
# @!attribute [rw] low
#   @return [Float, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
#
# @!attribute [rw] volume
#   @return [Float, nil]
TickerListMatch = Struct.new(
  :ask,
  :bid,
  :high,
  :last,
  :low,
  :symbol,
  :timestamp,
  :volume,
  keyword_init: true
)

# Trade entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] price
#   @return [Float, nil]
#
# @!attribute [rw] side
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
Trade = Struct.new(
  :amount,
  :id,
  :price,
  :side,
  :timestamp,
  keyword_init: true
)

# Request payload for Trade#load.
#
# @!attribute [rw] id
#   @return [String]
TradeLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Withdrawal entity data model.
#
# @!attribute [rw] account_number
#   @return [String]
#
# @!attribute [rw] account_type
#   @return [String, nil]
#
# @!attribute [rw] address
#   @return [String]
#
# @!attribute [rw] agency
#   @return [String]
#
# @!attribute [rw] amount
#   @return [Float]
#
# @!attribute [rw] bank
#   @return [String]
#
# @!attribute [rw] currency
#   @return [String]
#
# @!attribute [rw] tag
#   @return [String, nil]
Withdrawal = Struct.new(
  :account_number,
  :account_type,
  :address,
  :agency,
  :amount,
  :bank,
  :currency,
  :tag,
  keyword_init: true
)

# Match filter for Withdrawal#create (any subset of Withdrawal fields).
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] account_type
#   @return [String, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] agency
#   @return [String, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] bank
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [String, nil]
WithdrawalCreateData = Struct.new(
  :account_number,
  :account_type,
  :address,
  :agency,
  :amount,
  :bank,
  :currency,
  :tag,
  keyword_init: true
)


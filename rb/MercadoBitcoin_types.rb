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

# Request payload for Balance#list.
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
# @!attribute [rw] id
#   @return [String, nil]
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
  :id,
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
#
# @!attribute [rw] interval
#   @return [String, nil]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
CandleLoadMatch = Struct.new(
  :id,
  :interval,
  :limit,
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
# @!attribute [rw] qrCode
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [String, nil]
DepositAddress = Struct.new(
  :address,
  :currency,
  :qrCode,
  :tag,
  keyword_init: true
)

# Request payload for DepositAddress#load.
#
# @!attribute [rw] currency
#   @return [String]
DepositAddressLoadMatch = Struct.new(
  :currency,
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

# Request payload for Order#list.
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
OrderListMatch = Struct.new(
  :status,
  :symbol,
  keyword_init: true
)

# Request payload for Order#create.
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
# @!attribute [rw] asks
#   @return [Array, nil]
#
# @!attribute [rw] bids
#   @return [Array, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
OrderBook = Struct.new(
  :asks,
  :bids,
  :timestamp,
  keyword_init: true
)

# Request payload for OrderBook#load.
#
# @!attribute [rw] symbol
#   @return [String]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
OrderBookLoadMatch = Struct.new(
  :symbol,
  :limit,
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
# @!attribute [rw] id
#   @return [String, nil]
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
  :id,
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

# Request payload for Ticker#list.
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
# @!attribute [rw] id
#   @return [String, nil]
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
  :id,
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
#
# @!attribute [rw] limit
#   @return [Integer, nil]
TradeLoadMatch = Struct.new(
  :id,
  :limit,
  keyword_init: true
)

# Withdrawal entity data model.
#
# @!attribute [rw] accountNumber
#   @return [String]
#
# @!attribute [rw] accountType
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
  :accountNumber,
  :accountType,
  :address,
  :agency,
  :amount,
  :bank,
  :currency,
  :tag,
  keyword_init: true
)

# Request payload for Withdrawal#create.
#
# @!attribute [rw] accountNumber
#   @return [String]
#
# @!attribute [rw] accountType
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
WithdrawalCreateData = Struct.new(
  :accountNumber,
  :accountType,
  :address,
  :agency,
  :amount,
  :bank,
  :currency,
  :tag,
  keyword_init: true
)


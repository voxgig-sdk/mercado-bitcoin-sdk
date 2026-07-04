// Typed models for the MercadoBitcoin SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Balance is the typed data model for the balance entity.
type Balance struct {
	Available *float64 `json:"available,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Locked *float64 `json:"locked,omitempty"`
	Total *float64 `json:"total,omitempty"`
}

// BalanceListMatch mirrors the balance fields as an all-optional match
// filter (Go analog of Partial<Balance>).
type BalanceListMatch struct {
	Available *float64 `json:"available,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Locked *float64 `json:"locked,omitempty"`
	Total *float64 `json:"total,omitempty"`
}

// Candle is the typed data model for the candle entity.
type Candle struct {
	Close *float64 `json:"close,omitempty"`
	High *float64 `json:"high,omitempty"`
	Low *float64 `json:"low,omitempty"`
	Open *float64 `json:"open,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
	Volume *float64 `json:"volume,omitempty"`
}

// CandleLoadMatch is the typed request payload for Candle.LoadTyped.
type CandleLoadMatch struct {
	Id string `json:"id"`
}

// DepositAddress is the typed data model for the deposit_address entity.
type DepositAddress struct {
	Address *string `json:"address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	QrCode *string `json:"qr_code,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// DepositAddressLoadMatch mirrors the deposit_address fields as an all-optional match
// filter (Go analog of Partial<DepositAddress>).
type DepositAddressLoadMatch struct {
	Address *string `json:"address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	QrCode *string `json:"qr_code,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// Order is the typed data model for the order entity.
type Order struct {
	Amount *float64 `json:"amount,omitempty"`
	Filled *float64 `json:"filled,omitempty"`
	Id *string `json:"id,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
}

// OrderLoadMatch is the typed request payload for Order.LoadTyped.
type OrderLoadMatch struct {
	Id string `json:"id"`
}

// OrderListMatch mirrors the order fields as an all-optional match
// filter (Go analog of Partial<Order>).
type OrderListMatch struct {
	Amount *float64 `json:"amount,omitempty"`
	Filled *float64 `json:"filled,omitempty"`
	Id *string `json:"id,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
}

// OrderCreateData mirrors the order fields as an all-optional match
// filter (Go analog of Partial<Order>).
type OrderCreateData struct {
	Amount *float64 `json:"amount,omitempty"`
	Filled *float64 `json:"filled,omitempty"`
	Id *string `json:"id,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
}

// OrderRemoveMatch is the typed request payload for Order.RemoveTyped.
type OrderRemoveMatch struct {
	Id string `json:"id"`
}

// OrderBook is the typed data model for the order_book entity.
type OrderBook struct {
	Ask *[]any `json:"ask,omitempty"`
	Bid *[]any `json:"bid,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
}

// OrderBookLoadMatch is the typed request payload for OrderBook.LoadTyped.
type OrderBookLoadMatch struct {
	Symbol string `json:"symbol"`
}

// Ticker is the typed data model for the ticker entity.
type Ticker struct {
	Ask *float64 `json:"ask,omitempty"`
	Bid *float64 `json:"bid,omitempty"`
	High *float64 `json:"high,omitempty"`
	Last *float64 `json:"last,omitempty"`
	Low *float64 `json:"low,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
	Volume *float64 `json:"volume,omitempty"`
}

// TickerLoadMatch is the typed request payload for Ticker.LoadTyped.
type TickerLoadMatch struct {
	Id string `json:"id"`
}

// TickerListMatch mirrors the ticker fields as an all-optional match
// filter (Go analog of Partial<Ticker>).
type TickerListMatch struct {
	Ask *float64 `json:"ask,omitempty"`
	Bid *float64 `json:"bid,omitempty"`
	High *float64 `json:"high,omitempty"`
	Last *float64 `json:"last,omitempty"`
	Low *float64 `json:"low,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
	Volume *float64 `json:"volume,omitempty"`
}

// Trade is the typed data model for the trade entity.
type Trade struct {
	Amount *float64 `json:"amount,omitempty"`
	Id *string `json:"id,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Side *string `json:"side,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
}

// TradeLoadMatch is the typed request payload for Trade.LoadTyped.
type TradeLoadMatch struct {
	Id string `json:"id"`
}

// Withdrawal is the typed data model for the withdrawal entity.
type Withdrawal struct {
	AccountNumber string `json:"account_number"`
	AccountType *string `json:"account_type,omitempty"`
	Address string `json:"address"`
	Agency string `json:"agency"`
	Amount float64 `json:"amount"`
	Bank string `json:"bank"`
	Currency string `json:"currency"`
	Tag *string `json:"tag,omitempty"`
}

// WithdrawalCreateData mirrors the withdrawal fields as an all-optional match
// filter (Go analog of Partial<Withdrawal>).
type WithdrawalCreateData struct {
	AccountNumber *string `json:"account_number,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
	Address *string `json:"address,omitempty"`
	Agency *string `json:"agency,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Bank *string `json:"bank,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

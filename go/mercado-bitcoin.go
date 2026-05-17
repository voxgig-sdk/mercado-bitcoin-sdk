package voxgigmercadobitcoinsdk

import (
	"github.com/voxgig-sdk/mercado-bitcoin-sdk/go/core"
	"github.com/voxgig-sdk/mercado-bitcoin-sdk/go/entity"
	"github.com/voxgig-sdk/mercado-bitcoin-sdk/go/feature"
	_ "github.com/voxgig-sdk/mercado-bitcoin-sdk/go/utility"
)

// Type aliases preserve external API.
type MercadoBitcoinSDK = core.MercadoBitcoinSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type MercadoBitcoinEntity = core.MercadoBitcoinEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type MercadoBitcoinError = core.MercadoBitcoinError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewBalanceEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewBalanceEntity(client, entopts)
	}
	core.NewCandleEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewCandleEntity(client, entopts)
	}
	core.NewDepositAddressEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewDepositAddressEntity(client, entopts)
	}
	core.NewOrderEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewOrderEntity(client, entopts)
	}
	core.NewOrderBookEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewOrderBookEntity(client, entopts)
	}
	core.NewTickerEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewTickerEntity(client, entopts)
	}
	core.NewTradeEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewTradeEntity(client, entopts)
	}
	core.NewWithdrawalEntityFunc = func(client *core.MercadoBitcoinSDK, entopts map[string]any) core.MercadoBitcoinEntity {
		return entity.NewWithdrawalEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewMercadoBitcoinSDK = core.NewMercadoBitcoinSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature

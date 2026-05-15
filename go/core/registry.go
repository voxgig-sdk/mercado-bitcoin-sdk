package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewBalanceEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewCandleEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewDepositAddressEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewOrderEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewOrderBookEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewTickerEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewTradeEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity

var NewWithdrawalEntityFunc func(client *MercadoBitcoinSDK, entopts map[string]any) MercadoBitcoinEntity


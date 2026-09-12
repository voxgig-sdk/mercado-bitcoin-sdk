import { BalanceEntity } from './entity/BalanceEntity';
import { CandleEntity } from './entity/CandleEntity';
import { DepositAddressEntity } from './entity/DepositAddressEntity';
import { OrderEntity } from './entity/OrderEntity';
import { OrderBookEntity } from './entity/OrderBookEntity';
import { TickerEntity } from './entity/TickerEntity';
import { TradeEntity } from './entity/TradeEntity';
import { WithdrawalEntity } from './entity/WithdrawalEntity';
export type * from './MercadoBitcoinTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { MercadoBitcoinEntityBase } from './MercadoBitcoinEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class MercadoBitcoinSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Balance(entopts?: Record<string, any>): BalanceEntity;
    Candle(entopts?: Record<string, any>): CandleEntity;
    DepositAddress(entopts?: Record<string, any>): DepositAddressEntity;
    Order(entopts?: Record<string, any>): OrderEntity;
    OrderBook(entopts?: Record<string, any>): OrderBookEntity;
    Ticker(entopts?: Record<string, any>): TickerEntity;
    Trade(entopts?: Record<string, any>): TradeEntity;
    Withdrawal(entopts?: Record<string, any>): WithdrawalEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): MercadoBitcoinSDK;
    tester(testopts?: any, sdkopts?: any): MercadoBitcoinSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof MercadoBitcoinSDK;
export { stdutil, config, BaseFeature, MercadoBitcoinEntityBase, MercadoBitcoinSDK, SDK, };

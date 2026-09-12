import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { Trade, TradeLoadMatch } from '../MercadoBitcoinTypes';
declare class TradeEntity extends MercadoBitcoinEntityBase<Trade> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: TradeEntity): TradeEntity;
    load(this: any, reqmatch?: TradeLoadMatch, ctrl?: Control): Promise<TradeEntity>;
}
export { TradeEntity };

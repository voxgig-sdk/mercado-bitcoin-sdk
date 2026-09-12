import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { Ticker, TickerLoadMatch, TickerListMatch } from '../MercadoBitcoinTypes';
declare class TickerEntity extends MercadoBitcoinEntityBase<Ticker> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: TickerEntity): TickerEntity;
    load(this: any, reqmatch?: TickerLoadMatch, ctrl?: Control): Promise<TickerEntity>;
    list(this: any, reqmatch?: TickerListMatch, ctrl?: Control): Promise<TickerEntity[]>;
}
export { TickerEntity };

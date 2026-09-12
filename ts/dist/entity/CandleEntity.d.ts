import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { Candle, CandleLoadMatch } from '../MercadoBitcoinTypes';
declare class CandleEntity extends MercadoBitcoinEntityBase<Candle> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: CandleEntity): CandleEntity;
    load(this: any, reqmatch?: CandleLoadMatch, ctrl?: Control): Promise<CandleEntity>;
}
export { CandleEntity };

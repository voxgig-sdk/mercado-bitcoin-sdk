import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { OrderBook, OrderBookLoadMatch } from '../MercadoBitcoinTypes';
declare class OrderBookEntity extends MercadoBitcoinEntityBase<OrderBook> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: OrderBookEntity): OrderBookEntity;
    load(this: any, reqmatch?: OrderBookLoadMatch, ctrl?: Control): Promise<OrderBookEntity>;
}
export { OrderBookEntity };

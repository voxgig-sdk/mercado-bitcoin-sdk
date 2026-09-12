import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { Order, OrderLoadMatch, OrderListMatch, OrderCreateData, OrderRemoveMatch } from '../MercadoBitcoinTypes';
declare class OrderEntity extends MercadoBitcoinEntityBase<Order> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: OrderEntity): OrderEntity;
    load(this: any, reqmatch?: OrderLoadMatch, ctrl?: Control): Promise<OrderEntity>;
    list(this: any, reqmatch?: OrderListMatch, ctrl?: Control): Promise<OrderEntity[]>;
    create(this: any, reqdata?: OrderCreateData, ctrl?: Control): Promise<OrderEntity>;
    remove(this: any, reqmatch?: OrderRemoveMatch, ctrl?: Control): Promise<OrderEntity>;
}
export { OrderEntity };

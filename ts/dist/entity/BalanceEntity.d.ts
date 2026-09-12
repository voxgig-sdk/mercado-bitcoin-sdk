import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { Balance, BalanceListMatch } from '../MercadoBitcoinTypes';
declare class BalanceEntity extends MercadoBitcoinEntityBase<Balance> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: BalanceEntity): BalanceEntity;
    list(this: any, reqmatch?: BalanceListMatch, ctrl?: Control): Promise<BalanceEntity[]>;
}
export { BalanceEntity };

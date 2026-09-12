import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { DepositAddress, DepositAddressLoadMatch } from '../MercadoBitcoinTypes';
declare class DepositAddressEntity extends MercadoBitcoinEntityBase<DepositAddress> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: DepositAddressEntity): DepositAddressEntity;
    load(this: any, reqmatch?: DepositAddressLoadMatch, ctrl?: Control): Promise<DepositAddressEntity>;
}
export { DepositAddressEntity };

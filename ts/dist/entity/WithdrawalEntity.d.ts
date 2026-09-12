import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase';
import type { MercadoBitcoinSDK } from '../MercadoBitcoinSDK';
import type { Control } from '../types';
import type { Withdrawal, WithdrawalCreateData } from '../MercadoBitcoinTypes';
declare class WithdrawalEntity extends MercadoBitcoinEntityBase<Withdrawal> {
    constructor(client: MercadoBitcoinSDK, entopts: any);
    make(this: WithdrawalEntity): WithdrawalEntity;
    create(this: any, reqdata?: WithdrawalCreateData, ctrl?: Control): Promise<WithdrawalEntity>;
}
export { WithdrawalEntity };

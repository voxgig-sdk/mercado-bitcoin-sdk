
import { inspect } from 'node:util'

import { MercadoBitcoinEntityBase } from '../MercadoBitcoinEntityBase'

import type {
  MercadoBitcoinSDK,
} from '../MercadoBitcoinSDK'


import type {
  Operation,
  Context,
  Control,
} from '../types'

import type {
  DepositAddress,
  DepositAddressLoadMatch,
} from '../MercadoBitcoinTypes'

// TODO: needs Entity superclass
class DepositAddressEntity extends MercadoBitcoinEntityBase<DepositAddress> {

  constructor(client: MercadoBitcoinSDK, entopts: any) {
    super(client, entopts)
    this.name = 'deposit_address'
    this.name_ = 'deposit_address'
    this.Name = 'DepositAddress'
  }


  make(this: DepositAddressEntity) {
    return new DepositAddressEntity(this._client, this.entopts())
  }



  async load(this: any, reqmatch?: DepositAddressLoadMatch, ctrl?: Control): Promise<DepositAddress> {

    const utility = this._utility

    const {
      makeContext,
      done,
      error,
      featureHook,
      makePoint,
      makeRequest,
      makeResponse,
      makeResult,
      makeSpec,
    } = utility

    let fres: Promise<any> | undefined = undefined

    let ctx: Context = makeContext({
      opname: 'load',
      ctrl,
      match: this._match,
      data: this._data,
      reqmatch
    }, this._entctx)

    try {

      fres = featureHook(ctx, 'PrePoint')
      if (fres instanceof Promise) { await fres }

      ctx.out.point = makePoint(ctx)
      if (ctx.out.point instanceof Error) {
        return error(ctx, ctx.out.point)
      }



      fres = featureHook(ctx, 'PreSpec')
      if (fres instanceof Promise) { await fres }

      ctx.out.spec = makeSpec(ctx)
      if (ctx.out.spec instanceof Error) {
        return error(ctx, ctx.out.spec)
      }



      fres = featureHook(ctx, 'PreRequest')
      if (fres instanceof Promise) { await fres }

      ctx.out.request = await makeRequest(ctx)
      if (ctx.out.request instanceof Error) {
        return error(ctx, ctx.out.request)
      }



      fres = featureHook(ctx, 'PreResponse')
      if (fres instanceof Promise) { await fres }

      ctx.out.response = await makeResponse(ctx)
      if (ctx.out.response instanceof Error) {
        return error(ctx, ctx.out.response)
      }



      fres = featureHook(ctx, 'PreResult')
      if (fres instanceof Promise) { await fres }

      ctx.out.result = await makeResult(ctx)
      if (ctx.out.result instanceof Error) {
        return error(ctx, ctx.out.result)
      }



      fres = featureHook(ctx, 'PreDone')
      if (fres instanceof Promise) { await fres }

      if (null != ctx.result) {
        if (null != ctx.result.resmatch) {
          this._match = ctx.result.resmatch
        }

        if (null != ctx.result.resdata) {
          this._data = ctx.result.resdata
        }
      }

      return done(ctx)
    }
    catch (err: any) {

      fres = featureHook(ctx, 'PreUnexpected')
      if (fres instanceof Promise) { await fres }

      err = this._unexpected(ctx, err)

      if (err) {
        throw err
      }
      else {
        // Off-happy-path (throw disabled): typed as any so the method's
        // Promise<DepositAddress> return stays clean under strict null checks.
        return undefined as any
      }
    }
  }






}


export {
  DepositAddressEntity
}

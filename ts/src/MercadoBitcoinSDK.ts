// MercadoBitcoin Ts SDK

import { BalanceEntity } from './entity/BalanceEntity'
import { CandleEntity } from './entity/CandleEntity'
import { DepositAddressEntity } from './entity/DepositAddressEntity'
import { OrderEntity } from './entity/OrderEntity'
import { OrderBookEntity } from './entity/OrderBookEntity'
import { TickerEntity } from './entity/TickerEntity'
import { TradeEntity } from './entity/TradeEntity'
import { WithdrawalEntity } from './entity/WithdrawalEntity'

export type * from './MercadoBitcoinTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { MercadoBitcoinEntityBase } from './MercadoBitcoinEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class MercadoBitcoinSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath
    const items = struct.items

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    items(this._options.feature, (fitem: [string, any]) => {
      const fname = fitem[0]
      const fopts = fitem[1]
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    })

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  async direct(fetchargs?: any) {
    const utility = this._utility
    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
    }
  }



  _balance?: BalanceEntity

  // Idiomatic facade: `client.balance.list()` / `client.balance.load({ id })`.
  get balance(): BalanceEntity {
    return (this._balance ??= new BalanceEntity(this, undefined))
  }

  /** @deprecated Use `client.balance` instead. */
  Balance(data?: any) {
    const self = this
    return new BalanceEntity(self,data)
  }


  _candle?: CandleEntity

  // Idiomatic facade: `client.candle.list()` / `client.candle.load({ id })`.
  get candle(): CandleEntity {
    return (this._candle ??= new CandleEntity(this, undefined))
  }

  /** @deprecated Use `client.candle` instead. */
  Candle(data?: any) {
    const self = this
    return new CandleEntity(self,data)
  }


  _deposit_address?: DepositAddressEntity

  // Idiomatic facade: `client.deposit_address.list()` / `client.deposit_address.load({ id })`.
  get deposit_address(): DepositAddressEntity {
    return (this._deposit_address ??= new DepositAddressEntity(this, undefined))
  }

  /** @deprecated Use `client.deposit_address` instead. */
  DepositAddress(data?: any) {
    const self = this
    return new DepositAddressEntity(self,data)
  }


  _order?: OrderEntity

  // Idiomatic facade: `client.order.list()` / `client.order.load({ id })`.
  get order(): OrderEntity {
    return (this._order ??= new OrderEntity(this, undefined))
  }

  /** @deprecated Use `client.order` instead. */
  Order(data?: any) {
    const self = this
    return new OrderEntity(self,data)
  }


  _order_book?: OrderBookEntity

  // Idiomatic facade: `client.order_book.list()` / `client.order_book.load({ id })`.
  get order_book(): OrderBookEntity {
    return (this._order_book ??= new OrderBookEntity(this, undefined))
  }

  /** @deprecated Use `client.order_book` instead. */
  OrderBook(data?: any) {
    const self = this
    return new OrderBookEntity(self,data)
  }


  _ticker?: TickerEntity

  // Idiomatic facade: `client.ticker.list()` / `client.ticker.load({ id })`.
  get ticker(): TickerEntity {
    return (this._ticker ??= new TickerEntity(this, undefined))
  }

  /** @deprecated Use `client.ticker` instead. */
  Ticker(data?: any) {
    const self = this
    return new TickerEntity(self,data)
  }


  _trade?: TradeEntity

  // Idiomatic facade: `client.trade.list()` / `client.trade.load({ id })`.
  get trade(): TradeEntity {
    return (this._trade ??= new TradeEntity(this, undefined))
  }

  /** @deprecated Use `client.trade` instead. */
  Trade(data?: any) {
    const self = this
    return new TradeEntity(self,data)
  }


  _withdrawal?: WithdrawalEntity

  // Idiomatic facade: `client.withdrawal.list()` / `client.withdrawal.load({ id })`.
  get withdrawal(): WithdrawalEntity {
    return (this._withdrawal ??= new WithdrawalEntity(this, undefined))
  }

  /** @deprecated Use `client.withdrawal` instead. */
  Withdrawal(data?: any) {
    const self = this
    return new WithdrawalEntity(self,data)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new MercadoBitcoinSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return MercadoBitcoinSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'MercadoBitcoin' }
  }

  toString() {
    return 'MercadoBitcoin ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = MercadoBitcoinSDK


export {
  stdutil,

  BaseFeature,
  MercadoBitcoinEntityBase,

  MercadoBitcoinSDK,
  SDK,
}



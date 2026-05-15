
import { Context } from './Context'


class MercadoBitcoinError extends Error {

  isMercadoBitcoinError = true

  sdk = 'MercadoBitcoin'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  MercadoBitcoinError
}


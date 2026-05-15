# MercadoBitcoin SDK utility: make_context
require_relative '../core/context'
module MercadoBitcoinUtilities
  MakeContext = ->(ctxmap, basectx) {
    MercadoBitcoinContext.new(ctxmap, basectx)
  }
end

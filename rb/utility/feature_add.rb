# MercadoBitcoin SDK utility: feature_add
module MercadoBitcoinUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end

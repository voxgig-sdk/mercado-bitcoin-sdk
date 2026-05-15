# MercadoBitcoin SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module MercadoBitcoinFeatures
  def self.make_feature(name)
    case name
    when "base"
      MercadoBitcoinBaseFeature.new
    when "test"
      MercadoBitcoinTestFeature.new
    else
      MercadoBitcoinBaseFeature.new
    end
  end
end

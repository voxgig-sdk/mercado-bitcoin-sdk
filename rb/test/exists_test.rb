# MercadoBitcoin SDK exists test

require "minitest/autorun"
require_relative "../MercadoBitcoin_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = MercadoBitcoinSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end

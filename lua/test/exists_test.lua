-- ProjectName SDK exists test

local sdk = require("mercado-bitcoin_sdk")

describe("MercadoBitcoinSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)

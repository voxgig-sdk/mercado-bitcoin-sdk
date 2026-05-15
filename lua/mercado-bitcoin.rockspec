package = "voxgig-sdk-mercado-bitcoin"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/mercado-bitcoin-sdk.git"
}
description = {
  summary = "MercadoBitcoin SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["mercado-bitcoin_sdk"] = "mercado-bitcoin_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}

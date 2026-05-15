-- MercadoBitcoin SDK error

local MercadoBitcoinError = {}
MercadoBitcoinError.__index = MercadoBitcoinError


function MercadoBitcoinError.new(code, msg, ctx)
  local self = setmetatable({}, MercadoBitcoinError)
  self.is_sdk_error = true
  self.sdk = "MercadoBitcoin"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function MercadoBitcoinError:error()
  return self.msg
end


function MercadoBitcoinError:__tostring()
  return self.msg
end


return MercadoBitcoinError

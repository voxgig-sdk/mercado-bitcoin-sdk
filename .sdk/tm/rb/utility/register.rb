# MercadoBitcoin SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

MercadoBitcoinUtility.registrar = ->(u) {
  u.clean = MercadoBitcoinUtilities::Clean
  u.done = MercadoBitcoinUtilities::Done
  u.make_error = MercadoBitcoinUtilities::MakeError
  u.feature_add = MercadoBitcoinUtilities::FeatureAdd
  u.feature_hook = MercadoBitcoinUtilities::FeatureHook
  u.feature_init = MercadoBitcoinUtilities::FeatureInit
  u.fetcher = MercadoBitcoinUtilities::Fetcher
  u.make_fetch_def = MercadoBitcoinUtilities::MakeFetchDef
  u.make_context = MercadoBitcoinUtilities::MakeContext
  u.make_options = MercadoBitcoinUtilities::MakeOptions
  u.make_request = MercadoBitcoinUtilities::MakeRequest
  u.make_response = MercadoBitcoinUtilities::MakeResponse
  u.make_result = MercadoBitcoinUtilities::MakeResult
  u.make_point = MercadoBitcoinUtilities::MakePoint
  u.make_spec = MercadoBitcoinUtilities::MakeSpec
  u.make_url = MercadoBitcoinUtilities::MakeUrl
  u.param = MercadoBitcoinUtilities::Param
  u.prepare_auth = MercadoBitcoinUtilities::PrepareAuth
  u.prepare_body = MercadoBitcoinUtilities::PrepareBody
  u.prepare_headers = MercadoBitcoinUtilities::PrepareHeaders
  u.prepare_method = MercadoBitcoinUtilities::PrepareMethod
  u.prepare_params = MercadoBitcoinUtilities::PrepareParams
  u.prepare_path = MercadoBitcoinUtilities::PreparePath
  u.prepare_query = MercadoBitcoinUtilities::PrepareQuery
  u.result_basic = MercadoBitcoinUtilities::ResultBasic
  u.result_body = MercadoBitcoinUtilities::ResultBody
  u.result_headers = MercadoBitcoinUtilities::ResultHeaders
  u.transform_request = MercadoBitcoinUtilities::TransformRequest
  u.transform_response = MercadoBitcoinUtilities::TransformResponse
}

<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

MercadoBitcoinUtility::setRegistrar(function (MercadoBitcoinUtility $u): void {
    $u->clean = [MercadoBitcoinClean::class, 'call'];
    $u->done = [MercadoBitcoinDone::class, 'call'];
    $u->make_error = [MercadoBitcoinMakeError::class, 'call'];
    $u->feature_add = [MercadoBitcoinFeatureAdd::class, 'call'];
    $u->feature_hook = [MercadoBitcoinFeatureHook::class, 'call'];
    $u->feature_init = [MercadoBitcoinFeatureInit::class, 'call'];
    $u->fetcher = [MercadoBitcoinFetcher::class, 'call'];
    $u->make_fetch_def = [MercadoBitcoinMakeFetchDef::class, 'call'];
    $u->make_context = [MercadoBitcoinMakeContext::class, 'call'];
    $u->make_options = [MercadoBitcoinMakeOptions::class, 'call'];
    $u->make_request = [MercadoBitcoinMakeRequest::class, 'call'];
    $u->make_response = [MercadoBitcoinMakeResponse::class, 'call'];
    $u->make_result = [MercadoBitcoinMakeResult::class, 'call'];
    $u->make_point = [MercadoBitcoinMakePoint::class, 'call'];
    $u->make_spec = [MercadoBitcoinMakeSpec::class, 'call'];
    $u->make_url = [MercadoBitcoinMakeUrl::class, 'call'];
    $u->param = [MercadoBitcoinParam::class, 'call'];
    $u->prepare_auth = [MercadoBitcoinPrepareAuth::class, 'call'];
    $u->prepare_body = [MercadoBitcoinPrepareBody::class, 'call'];
    $u->prepare_headers = [MercadoBitcoinPrepareHeaders::class, 'call'];
    $u->prepare_method = [MercadoBitcoinPrepareMethod::class, 'call'];
    $u->prepare_params = [MercadoBitcoinPrepareParams::class, 'call'];
    $u->prepare_path = [MercadoBitcoinPreparePath::class, 'call'];
    $u->prepare_query = [MercadoBitcoinPrepareQuery::class, 'call'];
    $u->result_basic = [MercadoBitcoinResultBasic::class, 'call'];
    $u->result_body = [MercadoBitcoinResultBody::class, 'call'];
    $u->result_headers = [MercadoBitcoinResultHeaders::class, 'call'];
    $u->transform_request = [MercadoBitcoinTransformRequest::class, 'call'];
    $u->transform_response = [MercadoBitcoinTransformResponse::class, 'call'];
});

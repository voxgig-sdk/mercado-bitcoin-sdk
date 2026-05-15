<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility: result_body

class MercadoBitcoinResultBody
{
    public static function call(MercadoBitcoinContext $ctx): ?MercadoBitcoinResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}

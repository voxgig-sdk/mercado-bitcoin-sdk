<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility: result_headers

class MercadoBitcoinResultHeaders
{
    public static function call(MercadoBitcoinContext $ctx): ?MercadoBitcoinResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}

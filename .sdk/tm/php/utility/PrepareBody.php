<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility: prepare_body

class MercadoBitcoinPrepareBody
{
    public static function call(MercadoBitcoinContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}

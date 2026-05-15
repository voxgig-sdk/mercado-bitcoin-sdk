<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility: feature_add

class MercadoBitcoinFeatureAdd
{
    public static function call(MercadoBitcoinContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}

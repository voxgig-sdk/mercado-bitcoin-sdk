<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility: feature_hook

class MercadoBitcoinFeatureHook
{
    public static function call(MercadoBitcoinContext $ctx, string $name): void
    {
        if (!$ctx->client) {
            return;
        }
        $features = $ctx->client->features ?? null;
        if (!$features) {
            return;
        }
        foreach ($features as $f) {
            if (method_exists($f, $name)) {
                $f->$name($ctx);
            }
        }
    }
}

<?php
declare(strict_types=1);

// MercadoBitcoin SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class MercadoBitcoinMakeContext
{
    public static function call(array $ctxmap, ?MercadoBitcoinContext $basectx): MercadoBitcoinContext
    {
        return new MercadoBitcoinContext($ctxmap, $basectx);
    }
}

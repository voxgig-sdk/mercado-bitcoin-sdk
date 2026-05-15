<?php
declare(strict_types=1);

// MercadoBitcoin SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class MercadoBitcoinFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new MercadoBitcoinBaseFeature();
            case "test":
                return new MercadoBitcoinTestFeature();
            default:
                return new MercadoBitcoinBaseFeature();
        }
    }
}

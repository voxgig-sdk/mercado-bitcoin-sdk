<?php
declare(strict_types=1);

// MercadoBitcoin SDK base feature

class MercadoBitcoinBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(MercadoBitcoinContext $ctx, array $options): void {}
    public function PostConstruct(MercadoBitcoinContext $ctx): void {}
    public function PostConstructEntity(MercadoBitcoinContext $ctx): void {}
    public function SetData(MercadoBitcoinContext $ctx): void {}
    public function GetData(MercadoBitcoinContext $ctx): void {}
    public function GetMatch(MercadoBitcoinContext $ctx): void {}
    public function SetMatch(MercadoBitcoinContext $ctx): void {}
    public function PrePoint(MercadoBitcoinContext $ctx): void {}
    public function PreSpec(MercadoBitcoinContext $ctx): void {}
    public function PreRequest(MercadoBitcoinContext $ctx): void {}
    public function PreResponse(MercadoBitcoinContext $ctx): void {}
    public function PreResult(MercadoBitcoinContext $ctx): void {}
    public function PreDone(MercadoBitcoinContext $ctx): void {}
    public function PreUnexpected(MercadoBitcoinContext $ctx): void {}
}

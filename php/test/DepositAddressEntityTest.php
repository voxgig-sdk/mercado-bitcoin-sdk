<?php
declare(strict_types=1);

// DepositAddress entity test

require_once __DIR__ . '/../mercadobitcoin_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class DepositAddressEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = MercadoBitcoinSDK::test(null, null);
        $ent = $testsdk->DepositAddress(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = deposit_address_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "deposit_address." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set MERCADO_BITCOIN_TEST_DEPOSIT_ADDRESS_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $deposit_address_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.deposit_address")));
        $deposit_address_ref01_data = null;
        if (count($deposit_address_ref01_data_raw) > 0) {
            $deposit_address_ref01_data = Helpers::to_map($deposit_address_ref01_data_raw[0][1]);
        }

        // LOAD
        $deposit_address_ref01_ent = $client->DepositAddress(null);
        $deposit_address_ref01_match_dt0 = [];
        $deposit_address_ref01_data_dt0_loaded = $deposit_address_ref01_ent->load($deposit_address_ref01_match_dt0, null);
        $this->assertNotNull($deposit_address_ref01_data_dt0_loaded);

    }
}

function deposit_address_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/deposit_address/DepositAddressTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = MercadoBitcoinSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["deposit_address01", "deposit_address02", "deposit_address03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("MERCADO_BITCOIN_TEST_DEPOSIT_ADDRESS_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "MERCADO_BITCOIN_TEST_DEPOSIT_ADDRESS_ENTID" => $idmap,
        "MERCADO_BITCOIN_TEST_LIVE" => "FALSE",
        "MERCADO_BITCOIN_TEST_EXPLAIN" => "FALSE",
        "MERCADO_BITCOIN_APIKEY" => "",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["MERCADO_BITCOIN_TEST_DEPOSIT_ADDRESS_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["MERCADO_BITCOIN_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            Runner::live_client_options(),
            [
                "apikey" => $env["MERCADO_BITCOIN_APIKEY"],
            ],
            // ismap, not a plain "?? []" default: an empty PHP array is a
            // LIST, and a non-map later entry REPLACES the accumulated map in
            // merge - so the no-extras call discarded live_client_options()
            // and the apikey/server map above it.
            Vs::ismap($extra) ? $extra : new \stdClass(),
        ]);
        $client = new MercadoBitcoinSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["MERCADO_BITCOIN_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["MERCADO_BITCOIN_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}

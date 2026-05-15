<?php
declare(strict_types=1);

// Trade entity test

require_once __DIR__ . '/../mercadobitcoin_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class TradeEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = MercadoBitcoinSDK::test(null, null);
        $ent = $testsdk->Trade(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = trade_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "trade." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set MERCADOBITCOIN_TEST_TRADE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $trade_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.trade")));
        $trade_ref01_data = null;
        if (count($trade_ref01_data_raw) > 0) {
            $trade_ref01_data = Helpers::to_map($trade_ref01_data_raw[0][1]);
        }

        // LOAD
        $trade_ref01_ent = $client->Trade(null);
        $trade_ref01_match_dt0 = [
            "id" => $trade_ref01_data["id"],
        ];
        [$trade_ref01_data_dt0_loaded, $err] = $trade_ref01_ent->load($trade_ref01_match_dt0, null);
        $this->assertNull($err);
        $trade_ref01_data_dt0_load_result = Helpers::to_map($trade_ref01_data_dt0_loaded);
        $this->assertNotNull($trade_ref01_data_dt0_load_result);
        $this->assertEquals($trade_ref01_data_dt0_load_result["id"], $trade_ref01_data["id"]);

    }
}

function trade_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/trade/TradeTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = MercadoBitcoinSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["trade01", "trade02", "trade03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("MERCADOBITCOIN_TEST_TRADE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "MERCADOBITCOIN_TEST_TRADE_ENTID" => $idmap,
        "MERCADOBITCOIN_TEST_LIVE" => "FALSE",
        "MERCADOBITCOIN_TEST_EXPLAIN" => "FALSE",
        "MERCADOBITCOIN_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["MERCADOBITCOIN_TEST_TRADE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["MERCADOBITCOIN_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["MERCADOBITCOIN_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new MercadoBitcoinSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["MERCADOBITCOIN_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["MERCADOBITCOIN_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}

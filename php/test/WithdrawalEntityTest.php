<?php
declare(strict_types=1);

// Withdrawal entity test

require_once __DIR__ . '/../mercadobitcoin_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class WithdrawalEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = MercadoBitcoinSDK::test(null, null);
        $ent = $testsdk->Withdrawal(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = withdrawal_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "withdrawal." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $withdrawal_ref01_ent = $client->Withdrawal(null);
        $withdrawal_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.withdrawal"), "withdrawal_ref01"));

        $withdrawal_ref01_data_result = $withdrawal_ref01_ent->create($withdrawal_ref01_data, null);
        $withdrawal_ref01_data = Helpers::to_map($withdrawal_ref01_data_result);
        $this->assertNotNull($withdrawal_ref01_data);

    }
}

function withdrawal_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/withdrawal/WithdrawalTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = MercadoBitcoinSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["withdrawal01", "withdrawal02", "withdrawal03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID" => $idmap,
        "MERCADOBITCOIN_TEST_LIVE" => "FALSE",
        "MERCADOBITCOIN_TEST_EXPLAIN" => "FALSE",
        "MERCADOBITCOIN_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID"]);
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

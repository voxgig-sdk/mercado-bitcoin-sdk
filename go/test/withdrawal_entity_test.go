package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/mercado-bitcoin-sdk"
	"github.com/voxgig-sdk/mercado-bitcoin-sdk/core"

	vs "github.com/voxgig/struct"
)

func TestWithdrawalEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Withdrawal(nil)
		if ent == nil {
			t.Fatal("expected non-nil WithdrawalEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := withdrawalBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "withdrawal." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		withdrawalRef01Ent := client.Withdrawal(nil)
		withdrawalRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "withdrawal"}, setup.data), "withdrawal_ref01"))

		withdrawalRef01DataResult, err := withdrawalRef01Ent.Create(withdrawalRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		withdrawalRef01Data = core.ToMapAny(withdrawalRef01DataResult)
		if withdrawalRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func withdrawalBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "withdrawal", "WithdrawalTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read withdrawal test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse withdrawal test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"withdrawal01", "withdrawal02", "withdrawal03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID": idmap,
		"MERCADOBITCOIN_TEST_LIVE":      "FALSE",
		"MERCADOBITCOIN_TEST_EXPLAIN":   "FALSE",
		"MERCADOBITCOIN_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["MERCADOBITCOIN_TEST_WITHDRAWAL_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["MERCADOBITCOIN_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["MERCADOBITCOIN_APIKEY"],
			},
			extra,
		})
		client = sdk.NewMercadoBitcoinSDK(core.ToMapAny(mergedOpts))
	}

	live := env["MERCADOBITCOIN_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["MERCADOBITCOIN_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}

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

func TestOrderEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Order(nil)
		if ent == nil {
			t.Fatal("expected non-nil OrderEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := orderBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "order." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set MERCADOBITCOIN_TEST_ORDER_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		orderRef01Ent := client.Order(nil)
		orderRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "order"}, setup.data), "order_ref01"))

		orderRef01DataResult, err := orderRef01Ent.Create(orderRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		orderRef01Data = core.ToMapAny(orderRef01DataResult)
		if orderRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if orderRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		orderRef01Match := map[string]any{}

		orderRef01ListResult, err := orderRef01Ent.List(orderRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		orderRef01List, orderRef01ListOk := orderRef01ListResult.([]any)
		if !orderRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", orderRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(orderRef01List), map[string]any{"id": orderRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// LOAD
		orderRef01MatchDt0 := map[string]any{
			"id": orderRef01Data["id"],
		}
		orderRef01DataDt0Loaded, err := orderRef01Ent.Load(orderRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		orderRef01DataDt0LoadResult := core.ToMapAny(orderRef01DataDt0Loaded)
		if orderRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if orderRef01DataDt0LoadResult["id"] != orderRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		orderRef01MatchRm0 := map[string]any{
			"id": orderRef01Data["id"],
		}
		_, err = orderRef01Ent.Remove(orderRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		orderRef01MatchRt0 := map[string]any{}

		orderRef01ListRt0Result, err := orderRef01Ent.List(orderRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		orderRef01ListRt0, orderRef01ListRt0Ok := orderRef01ListRt0Result.([]any)
		if !orderRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", orderRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(orderRef01ListRt0), map[string]any{"id": orderRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func orderBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "order", "OrderTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read order test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse order test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"order01", "order02", "order03"},
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
	entidEnvRaw := os.Getenv("MERCADOBITCOIN_TEST_ORDER_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"MERCADOBITCOIN_TEST_ORDER_ENTID": idmap,
		"MERCADOBITCOIN_TEST_LIVE":      "FALSE",
		"MERCADOBITCOIN_TEST_EXPLAIN":   "FALSE",
		"MERCADOBITCOIN_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["MERCADOBITCOIN_TEST_ORDER_ENTID"])
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

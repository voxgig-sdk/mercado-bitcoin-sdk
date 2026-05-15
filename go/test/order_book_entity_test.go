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

func TestOrderBookEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.OrderBook(nil)
		if ent == nil {
			t.Fatal("expected non-nil OrderBookEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := order_bookBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "order_book." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set MERCADOBITCOIN_TEST_ORDER_BOOK_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		orderBookRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.order_book", setup.data)))
		var orderBookRef01Data map[string]any
		if len(orderBookRef01DataRaw) > 0 {
			orderBookRef01Data = core.ToMapAny(orderBookRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = orderBookRef01Data

		// LOAD
		orderBookRef01Ent := client.OrderBook(nil)
		orderBookRef01MatchDt0 := map[string]any{}
		orderBookRef01DataDt0Loaded, err := orderBookRef01Ent.Load(orderBookRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if orderBookRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func order_bookBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "order_book", "OrderBookTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read order_book test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse order_book test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"order_book01", "order_book02", "order_book03", "orderbook01", "orderbook02", "orderbook03"},
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
	entidEnvRaw := os.Getenv("MERCADOBITCOIN_TEST_ORDER_BOOK_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"MERCADOBITCOIN_TEST_ORDER_BOOK_ENTID": idmap,
		"MERCADOBITCOIN_TEST_LIVE":      "FALSE",
		"MERCADOBITCOIN_TEST_EXPLAIN":   "FALSE",
		"MERCADOBITCOIN_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["MERCADOBITCOIN_TEST_ORDER_BOOK_ENTID"])
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

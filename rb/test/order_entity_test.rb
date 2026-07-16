# Order entity test

require "minitest/autorun"
require "json"
require_relative "../MercadoBitcoin_sdk"
require_relative "runner"

class OrderEntityTest < Minitest::Test
  def test_create_instance
    testsdk = MercadoBitcoinSDK.test(nil, nil)
    ent = testsdk.Order(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "order" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = MercadoBitcoinSDK.test(seed, nil)
    seen = base.Order(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = MercadoBitcoinConfig.make_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = MercadoBitcoinSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.Order(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = order_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "order." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set MERCADOBITCOIN_TEST_ORDER_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    order_ref01_ent = client.Order(nil)
    order_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.order"), "order_ref01"))

    order_ref01_data_result = order_ref01_ent.create(order_ref01_data, nil)
    order_ref01_data = Helpers.to_map(order_ref01_data_result)
    assert !order_ref01_data.nil?
    assert !order_ref01_data["id"].nil?

    # LIST
    order_ref01_match = {}

    order_ref01_list_result = order_ref01_ent.list(order_ref01_match, nil)
    assert order_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(order_ref01_list_result),
      { "id" => order_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # LOAD
    order_ref01_match_dt0 = {
      "id" => order_ref01_data["id"],
    }
    order_ref01_data_dt0_loaded = order_ref01_ent.load(order_ref01_match_dt0, nil)
    order_ref01_data_dt0_load_result = Helpers.to_map(order_ref01_data_dt0_loaded)
    assert !order_ref01_data_dt0_load_result.nil?
    assert_equal order_ref01_data_dt0_load_result["id"], order_ref01_data["id"]

    # REMOVE
    order_ref01_match_rm0 = {
      "id" => order_ref01_data["id"],
    }
    order_ref01_ent.remove(order_ref01_match_rm0, nil)

    # LIST
    order_ref01_match_rt0 = {}

    order_ref01_list_rt0_result = order_ref01_ent.list(order_ref01_match_rt0, nil)
    assert order_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(order_ref01_list_rt0_result),
      { "id" => order_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def order_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "order", "OrderTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = MercadoBitcoinSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["order01", "order02", "order03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["MERCADOBITCOIN_TEST_ORDER_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "MERCADOBITCOIN_TEST_ORDER_ENTID" => idmap,
    "MERCADOBITCOIN_TEST_LIVE" => "FALSE",
    "MERCADOBITCOIN_TEST_EXPLAIN" => "FALSE",
    "MERCADOBITCOIN_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["MERCADOBITCOIN_TEST_ORDER_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["MERCADOBITCOIN_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["MERCADOBITCOIN_APIKEY"],
      },
      extra || {},
    ])
    client = MercadoBitcoinSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["MERCADOBITCOIN_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["MERCADOBITCOIN_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end

package.path = "charts/synapse/scripts/?.lua;" .. package.path

local synapse = require("synapse")

local function assert_equal(actual, expected, message)
    if actual ~= expected then
        error(string.format("%s: expected %q, got %q", message, tostring(expected), tostring(actual)), 2)
    end
end

local function new_headers(values)
    return {
        add = function(_, key, value)
            values[key] = value
        end,
        get = function(_, key)
            return values[key]
        end,
        values = values
    }
end

local function new_request_handle(values, http_call)
    local headers = new_headers(values)
    local logs = {}

    return {
        headers = function()
            return headers
        end,
        httpCall = http_call,
        logWarn = function(_, message)
            table.insert(logs, message)
        end,
        logErr = function(_, message)
            table.insert(logs, message)
        end,
        logs = logs
    }
end

local function test_room_id_from_encoded_path()
    local headers = new_headers({
        [":path"] = "/_matrix/client/v3/rooms/%21room.name-1%3Aexample.org/messages"
    })

    assert_equal(
        synapse.get_room_id_from_request(headers),
        "!room.name-1:example.org",
        "encoded room id should be normalized"
    )
end

local function test_access_token_from_authorization_header()
    local headers = new_headers({
        [":path"] = "/_matrix/client/v3/sync",
        ["authorization"] = "Bearer token-123"
    })

    assert_equal(
        synapse.get_access_token_from_request(headers),
        "token-123",
        "bearer token should be extracted"
    )
end

local function test_access_token_from_query()
    local headers = new_headers({
        [":path"] = "/_matrix/client/r0/events?access_token=query-token"
    })

    assert_equal(
        synapse.get_access_token_from_request(headers),
        "query-token",
        "query access_token should be extracted"
    )
end

local function test_whoami_lookup_returns_localpart()
    local calls = 0
    local request_handle = new_request_handle({
        [":path"] = "/_matrix/client/v3/sync",
        [":authority"] = "matrix.example.org",
        ["authorization"] = "Bearer whoami-token"
    }, function(_, cluster, headers, body, timeout_ms)
        calls = calls + 1
        assert_equal(cluster, "httpd", "whoami cluster")
        assert_equal(headers[":path"], "/_matrix/client/v3/account/whoami", "whoami path")
        assert_equal(headers["authorization"], "Bearer whoami-token", "whoami authorization")
        assert_equal(body, "", "whoami body")
        assert_equal(timeout_ms, 5000, "whoami timeout")

        return { [":status"] = "200" }, '{"user_id":"@alice:example.org"}'
    end)

    assert_equal(
        synapse.get_user_identifier_from_request(request_handle, {}),
        "alice",
        "whoami user_id should resolve to localpart"
    )
    assert_equal(calls, 1, "whoami should be called once")
end

local function test_whoami_lookup_is_cached()
    local request_handle = new_request_handle({
        [":path"] = "/_matrix/client/v3/sync",
        ["authorization"] = "Bearer cached-token"
    }, function()
        return { [":status"] = "200" }, '{"user_id":"@bob:example.org"}'
    end)

    assert_equal(
        synapse.get_user_identifier_from_request(request_handle, { cache_ttl_seconds = 300 }),
        "bob",
        "first whoami lookup should resolve localpart"
    )

    local cached_request_handle = new_request_handle({
        [":path"] = "/_matrix/client/v3/sync",
        ["authorization"] = "Bearer cached-token"
    }, function()
        error("cached token should not call whoami")
    end)

    assert_equal(
        synapse.get_user_identifier_from_request(cached_request_handle, { cache_ttl_seconds = 300 }),
        "bob",
        "second whoami lookup should use cache"
    )
end

local function test_whoami_failure_falls_back_to_token()
    local request_handle = new_request_handle({
        [":path"] = "/_matrix/client/v3/sync",
        ["authorization"] = "Bearer fallback-token"
    }, function()
        return { [":status"] = "401" }, "{}"
    end)

    assert_equal(
        synapse.get_user_identifier_from_request(request_handle, {}),
        "fallback-token",
        "failed whoami lookup should fall back to token"
    )
end

local function test_missing_token_returns_nil()
    local request_handle = new_request_handle({
        [":path"] = "/_matrix/client/v3/sync"
    }, function()
        error("missing token should not call whoami")
    end)

    assert_equal(
        synapse.get_user_identifier_from_request(request_handle, {}),
        nil,
        "missing token should return nil"
    )
end

local function test_hash_fallback_logs_and_sets_request_id()
    local request_handle = new_request_handle({
        [":method"] = "GET",
        [":path"] = "/_matrix/client/v3/sync",
        [":authority"] = "matrix.example.org",
        ["x-request-id"] = "request-1"
    }, function()
        error("fallback should not call whoami")
    end)
    local headers = request_handle:headers()

    synapse.set_request_id_hash_key_with_fallback_log(request_handle, headers, "user")

    assert_equal(headers.values["X-Hash-Key"], "request-1", "fallback should set X-Hash-Key")
    assert_equal(
        request_handle.logs[1],
        "synapse_envoy_user_hash_fallback: method=GET path=/_matrix/client/v3/sync authority=matrix.example.org request_id=request-1",
        "fallback should log request details"
    )
end

local tests = {
    test_room_id_from_encoded_path,
    test_access_token_from_authorization_header,
    test_access_token_from_query,
    test_whoami_lookup_returns_localpart,
    test_whoami_lookup_is_cached,
    test_whoami_failure_falls_back_to_token,
    test_missing_token_returns_nil,
    test_hash_fallback_logs_and_sets_request_id
}

for _, test in ipairs(tests) do
    test()
end

print(string.format("ok - %d synapse lua tests", #tests))

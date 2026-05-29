local room_id_pattern = "(![A-Za-z0-9._=%%%-/]+:[A-Za-z0-9.%-]+)"
local whoami_cache = {}

local function normalize_matrix_room_separators(path)
    return path:gsub("%%21", "!"):gsub("%%3[Aa]", ":")
end

local function get_room_id_from_namespace_path(path, namespace)
    local _, _, room_id = string.find(path, "^/_matrix/" .. namespace .. "/.-" .. room_id_pattern)

    return room_id
end

local function get_room_id_from_path(path)
    local normalized_path = normalize_matrix_room_separators(path)

    local key = get_room_id_from_namespace_path(normalized_path, "client")
    if key ~= nil then
        return key
    end

    key = get_room_id_from_namespace_path(normalized_path, "federation")
    if key ~= nil then
        return key
    end

    local _, _, key = string.find(path, "/_matrix/client/v3/rooms/([^/]+)/messages")

    return key
end

local function get_room_id_from_request(headers)
    local path = headers:get(":path")

    return get_room_id_from_path(path)
end

local function get_access_token(auth_header, path)
    if auth_header ~= nil and string.len(auth_header) > 0 then
        local _, _, bearer_token = string.find(auth_header, "^[Bb]earer%s+(.+)$")
        if bearer_token ~= nil then
            return bearer_token
        end

        return auth_header
    end

    local _, _, token_param = string.find(path, "access_token=([^&]+)")
    if token_param ~= nil then
        return token_param
    end

    return auth_header
end

local function get_access_token_from_request(headers)
    local path = headers:get(":path")

    return get_access_token(headers:get("authorization"), path)
end

local function log_hash_fallback(request_handle, headers, fallback_type, request_id)
    request_handle:logWarn(
        "synapse_envoy_" .. fallback_type .. "_hash_fallback: method="
        .. tostring(headers:get(":method"))
        .. " path="
        .. tostring(headers:get(":path"))
        .. " authority="
        .. tostring(headers:get(":authority"))
        .. " request_id="
        .. tostring(request_id)
    )
end

local function set_request_id_hash_key_with_fallback_log(request_handle, headers, fallback_type)
    local request_id = headers:get("x-request-id")

    log_hash_fallback(request_handle, headers, fallback_type, request_id)
    headers:add("X-Hash-Key", request_id)
end

local function get_option(options, key, default)
    if options ~= nil and options[key] ~= nil then
        return options[key]
    end

    return default
end

local function log(request_handle, options, level, message)
    if not get_option(options, "logging_enabled", false) then
        return
    end

    local prefixed_message = "whoami_sync_worker_router: " .. message
    if level == "error" then
        request_handle:logErr(prefixed_message)
        return
    end

    request_handle:logWarn(prefixed_message)
end

local function truncate_token(token, options)
    local token_length = get_option(options, "logging_token_length", 8)
    if token == nil or string.len(token) <= token_length then
        return token
    end

    return string.sub(token, 1, token_length) .. "..."
end

local function extract_localpart(user_id)
    if user_id == nil or string.sub(user_id, 1, 1) ~= "@" then
        return nil
    end

    local colon_index = string.find(user_id, ":", 2, true)
    if colon_index == nil then
        return nil
    end

    return string.sub(user_id, 2, colon_index - 1)
end

local function extract_user_id_from_whoami_body(body)
    local _, _, user_id = string.find(body, '"user_id"%s*:%s*"([^"]+)"')

    return user_id
end

local function get_cached_username(token, options)
    local entry = whoami_cache[token]
    if entry == nil then
        return nil
    end

    if entry.expires_at > os.time() then
        return entry.username
    end

    whoami_cache[token] = nil
    return nil
end

local function cache_username(token, username, options)
    local ttl_seconds = get_option(options, "cache_ttl_seconds", 300)
    whoami_cache[token] = {
        username = username,
        expires_at = os.time() + ttl_seconds
    }
end

local function lookup_whoami(request_handle, token, options)
    local headers = request_handle:headers()
    local authority = headers:get(":authority")
    if authority == nil then
        authority = "synapse-client-reader-headless"
    end

    log(request_handle, options, "warn", "performing whoami lookup for token " .. truncate_token(token, options))

    local call_headers = {
        [":method"] = "GET",
        [":path"] = get_option(options, "whoami_path", "/_matrix/client/v3/account/whoami"),
        [":authority"] = authority,
        ["authorization"] = "Bearer " .. token,
        ["x-forwarded-proto"] = "https"
    }
    local xff = headers:get("x-forwarded-for")
    if xff ~= nil then
        call_headers["x-forwarded-for"] = xff
    end

    local ok, response_headers, response_body = pcall(function()
        return request_handle:httpCall(
            get_option(options, "whoami_cluster", "httpd"),
            call_headers,
            "",
            get_option(options, "timeout_ms", 5000)
        )
    end)

    if not ok then
        log(request_handle, options, "error", "whoami lookup failed: " .. tostring(response_headers))
        return nil
    end

    local status = response_headers[":status"]
    if status ~= "200" then
        if status == "401" then
            log(request_handle, options, "warn", "whoami lookup returned 401 for token " .. truncate_token(token, options))
        else
            log(request_handle, options, "error", "whoami lookup returned status " .. tostring(status))
        end
        return nil
    end

    local user_id = extract_user_id_from_whoami_body(response_body)
    local username = extract_localpart(user_id)
    if username ~= nil then
        log(request_handle, options, "warn", "whoami lookup success: " .. user_id .. " -> " .. username)
    end

    return username
end

local function get_user_identifier_from_request(request_handle, options)
    local headers = request_handle:headers()
    local token = get_access_token_from_request(headers)
    if token == nil or string.len(token) == 0 then
        log(request_handle, options, "warn", "no token found in request")
        return nil
    end

    local cached_username = get_cached_username(token, options)
    if cached_username ~= nil then
        log(request_handle, options, "warn", "cache hit for token " .. truncate_token(token, options) .. " -> " .. cached_username)
        return cached_username
    end

    local username = lookup_whoami(request_handle, token, options)
    if username ~= nil then
        cache_username(token, username, options)
        return username
    end

    log(request_handle, options, "warn", "whoami lookup failed, falling back to token-based routing")
    return token
end

return {
    get_access_token_from_request = get_access_token_from_request,
    get_room_id_from_request = get_room_id_from_request,
    set_request_id_hash_key_with_fallback_log = set_request_id_hash_key_with_fallback_log,
    get_user_identifier_from_request = get_user_identifier_from_request
}

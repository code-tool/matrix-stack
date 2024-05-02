local function get_hash_key_from_path(path)
    local _, _, key = string.find(path, "/_matrix/client/v3/rooms/([^/]+)/messages")

    return key
end

local function parse_username_from_token(token)
    local _, _, username = string.find(token, "[^_]+_([^+]+)_.*$")
    if username ~= nil then
        return username
    end

    return token
end

local function get_auth_token(auth_header, path)
    if auth_header ~= nil and string.len(auth_header) > 0 then
        return parse_username_from_token(auth_header)
    end

    local _, _, token_param = string.find(path, "access_token=([^&]+)")
    if token_param ~= nil then
        return parse_username_from_token(auth_header)
    end

    return auth_header
end

local function get_hash_key_from_request(headers)
    local path = headers:get(":path")

    local result = get_hash_key_from_path(path)
    if result ~= nil then
        return result
    end

    return get_auth_token(headers:get("authorization"), path)
end

return {
    get_hash_key_from_request = get_hash_key_from_request
}

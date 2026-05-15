local room_id_pattern = "(![A-Za-z0-9._=%%%-/]+:[A-Za-z0-9.%-]+)"

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

return {
    get_access_token_from_request = get_access_token_from_request,
    get_room_id_from_request = get_room_id_from_request
}

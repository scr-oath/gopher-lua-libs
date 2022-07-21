local plugin = require("plugin")
local time = require("time")
local ioutil = require("ioutil")

function Test_curl_plugin(t)

    local plugin_body_1 = [[

local http = require("http")
local time = require("time")

local i = 1
local client = http.client()
local request = http.request("GET", "http://google.com")
os.remove("./test/file.txt")

while true do
    local result, err = client:do_request(request)
    if err then error(err) end
    time.sleep(0.1)
    local file = io.open("./test/file.txt", "w")
    file:write(tostring(i), "\n")
    file:close()
    i = i + 1
end

]]

    local curl_plugin = plugin.do_string(plugin_body_1)

    curl_plugin:run()
    time.sleep(2)
    curl_plugin:stop()
    time.sleep(1)

    local running = curl_plugin:is_running()
    assert(not running, "already running")

    local data, err = ioutil.read_file("./test/file.txt")
    assert(not err, err)

    t:Log(data)
    local i = tonumber(data)
    if i < 1 then
        error("i < 1")
    end
end

function Test_with_payload(t)
    local plugin_body_2 = [[
    local ioutil = require("ioutil")
    local err = ioutil.write_file("./test/payload.txt", payload)
    if err then error(err) end
]]
    local test_payload = "OK"
    local plugin_with_payload = plugin.do_string_with_payload(plugin_body_2, test_payload)
    plugin_with_payload:run()

    time.sleep(1)
    local i = 0
    while plugin_with_payload:is_running() do
        time.sleep(1)
        i = i + 1
        assert(i <= 3, "timeout")
    end

    local data, err = ioutil.read_file("./test/payload.txt")
    assert(not err, err)
    assert(data == test_payload, string.format("expected %s; got %s", test_payload, data))
end

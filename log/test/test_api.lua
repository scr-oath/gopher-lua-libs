local log = require("log")
local ioutil = require("ioutil")

os.remove("./test/test.log")

function Test_log(t)
    local info, err = log.new("./test/test.log")
    assert(not err, err)

    info:println("1", 2)

    -- check set prefix
    info:set_prefix("[INFO] ")
    info:printf("%s %.2f", "1", 2.2)

    -- check flags
    info:set_flags({ longfile = true })
    info:printf("ok\n")
    info:print("ok")
    info:println("ok")

    local err = info:close()
    assert(not err, err)

    -- check result
    local except_result = [[1 2
[INFO] 1 2.20
[INFO] ./test/test_api.lua:18: ok
[INFO] ./test/test_api.lua:19: ok
[INFO] ./test/test_api.lua:20: ok
]]
    local get_result, err = ioutil.read_file("./test/test.log")
    assert(not err, err)
    assert(except_result == get_result,
            string.format([[expected "%s"; got "%s"]], except_result, get_result))
end

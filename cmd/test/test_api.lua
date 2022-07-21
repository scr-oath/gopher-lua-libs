local cmd = require("cmd")
local runtime = require("runtime")

function Test_cmd_exec(t)
    t:Run("no timeout", function(t)
        local command = runtime.goos() == "windows" and "timeout 1" or "sleep 1"
        local result, err = cmd.exec(command)
        assert(not err, err)
        t:Log(result.stdout)
    end)

    t:Run("timeout", function(t)
        -- Test timeout
        local command = runtime.goos() == "windows" and "timeout 5" or "sleep 5"
        local result, err = cmd.exec(command, 1)
        assert(err, "timeout expected but did not occur")
        assert(err == "execute timeout", "expected 'execute timeout' but instead got '" .. err .. "'")
    end)
end

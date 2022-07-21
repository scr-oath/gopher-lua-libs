local filepath = require("filepath")

function Test_filepath(t)
    t:Run("join and separator", function(t)
        local path = "1"
        local need_path = path .. filepath.separator() .. "2" .. filepath.separator() .. "3"
        path = filepath.join(path, "2", "3")
        assert(path == need_path, "filepath.join()")
    end)

    t:Run("glob", function(t)
        local results = filepath.glob("test" .. filepath.separator() .. "*")
        assert(#results == 1)

        expected = "test" .. filepath.separator() .. "test_api.lua"
        got = results[1]
        assert(got == expected, string.format("expected %s; got %s", expected, got))
    end)
end

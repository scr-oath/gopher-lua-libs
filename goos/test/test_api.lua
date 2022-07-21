local goos = require("goos")

function Test_stat(t)
    local info, err = goos.stat("./test/test.file")
    assert(not err, err)
    assert(info.is_dir == false, "is dir")
    assert(0 == info.size, "size")
    assert(info.mod_time > 0, "mod_time")
    assert(info.mode > "", "mode")
end

function Test_hostname(t)
    local hostname, err = goos.hostname()
    assert(not err, err)
    assert(hostname > "", "hostname")
end

function Test_get_pagesize(t)
    assert(goos.get_pagesize() > 0, "pagesize")
end

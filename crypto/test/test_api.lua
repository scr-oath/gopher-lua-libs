local crypto = require("crypto")

function TestCrypto(t)
    tests = {
        {
            name = "md5(1)",
            input = "1\n",
            method = crypto.md5,
            expected = "b026324c6904b2a9cb4b88d6d61c81d1",
        },
        {
            name = "sha256(1)",
            input = "1\n",
            method = crypto.sha256,
            expected = "4355a46b19d348dc2f57c046f8ef63d4538ebb936000f3c9ee954a27460dd865",
        },
    }
    for _, tt in ipairs(tests) do
        t:Run(tt.name, function(t)
            got = tt.method(tt.input)
            assert(got == tt.expected, string.format("expected %s; got %s", tt.expected, got))
        end)
    end
end

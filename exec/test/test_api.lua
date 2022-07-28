local strings = require 'strings'
local inspect = require 'inspect'
local exec = require 'exec'

function assert_equal(expected, got)
    assert(got == expected, string.format([[expected "%s"; got "%s"]], expected, got))
end

function TestCommand(t)
    command = exec.command('mypath', "abc", "def")
    t:Log(inspect(command))
    assert_equal("mypath abc def", tostring(command))
end

function TestRun_echo(t)
    command = exec.command("echo", "foo bar baz")
    local sb = strings.new_builder()
    command.stdout = sb
    t:Log(inspect(command))
    local err = command:run()
    assert(not err, err)
    assert_equal("foo bar baz\n", sb:string())
end

function TestRun_printf(t)
    command = exec.command("printf", "%s\\n", "foo", "bar baz")
    local sb = strings.new_builder()
    command.stdout = sb
    t:Log(inspect(command))
    local err = command:run()
    assert(not err, err)
    assert_equal("foo\nbar baz\n", sb:string())
end

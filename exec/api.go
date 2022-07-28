// Package cmd implements golang cmd functionality for lua.
package exec

import (
	lio "github.com/vadv/gopher-lua-libs/io"
	lua "github.com/yuin/gopher-lua"
	"os/exec"
	"strings"
)

const (
	cmdType = "exec.Cmd"
)

func cmdCommand(L *lua.LState) int {
	path := L.CheckString(1)
	top := L.GetTop()
	args := L.CreateTable(top-1, 0)
	for i := 2; i <= top; i++ {
		args.Append(lua.LString(L.CheckString(i)))
	}
	command := L.CreateTable(0, 12)
	L.SetField(command, "path", lua.LString(path))
	L.SetField(command, "args", args)
	L.SetMetatable(command, L.GetTypeMetatable(cmdType))
	L.Push(command)
	return 1
}

func cmdString(L *lua.LState) int {
	command := L.CheckTable(1)
	path := lua.LVAsString(command.RawGetString("path"))
	args := command.RawGetString("args").(*lua.LTable)
	var sb strings.Builder
	sb.WriteString(path)
	args.ForEach(func(_ lua.LValue, val lua.LValue) {
		sb.WriteRune(' ')
		sb.WriteString(lua.LVAsString(L.ToStringMeta(val)))
	})
	L.Push(lua.LString(sb.String()))
	return 1
}

func cmdRun(L *lua.LState) int {
	command := L.CheckTable(1)
	path := lua.LVAsString(command.RawGetString("path"))
	var arg []string
	args := command.RawGetString("args").(*lua.LTable)
	args.ForEach(func(_ lua.LValue, value lua.LValue) {
		arg = append(arg, lua.LVAsString(value))
	})

	cmd := exec.Command(path, arg...)

	if env := command.RawGetString("env"); env != lua.LNil {
		L.Push(env)
		L.CheckTable(-1).ForEach(func(_ lua.LValue, value lua.LValue) {
			cmd.Env = append(cmd.Env, lua.LVAsString(value))
		})
	}
	if dir := command.RawGetString("dir"); dir != lua.LNil {
		L.Push(dir)
		cmd.Dir = L.CheckString(-1)
	}
	if stdin := command.RawGetString("stdin"); stdin != lua.LNil {
		L.Push(stdin)
		cmd.Stdin = lio.CheckIOReader(L, -1)
	}
	if stdout := command.RawGetString("stdout"); stdout != lua.LNil {
		L.Push(stdout)
		cmd.Stdout = lio.CheckIOWriter(L, -1)
	}
	if stderr := command.RawGetString("stdout"); stderr != lua.LNil {
		L.Push(stderr)
		cmd.Stderr = lio.CheckIOWriter(L, -1)
	}

	if err := cmd.Run(); err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}

func cmdStart(L *lua.LState) int {
	panic("implement me")
}

func cmdWait(L *lua.LState) int {
	panic("implement me")
}

func registerCmd(L *lua.LState) {
	mt := L.NewTypeMetatable(cmdType)
	L.SetGlobal(cmdType, mt)
	L.SetField(mt, "__tostring", L.NewFunction(cmdString))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"run":   cmdRun,
		"start": cmdStart,
		"wait":  cmdWait,
	}))
}

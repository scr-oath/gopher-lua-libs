// Package cmd implements golang cmd functionality for lua.
package exec

import (
	lio "github.com/vadv/gopher-lua-libs/io"
	lua "github.com/yuin/gopher-lua"
	"os/exec"
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
	cmd := checkCmd(L, 1)
	L.Push(lua.LString(cmd.String()))
	return 1
}

func checkCmd(L *lua.LState, n int) *exec.Cmd {
	command := L.CheckTable(n)

	// First check if this command was started
	if ud, ok := command.RawGetString("cmd").(*lua.LUserData); ok {
		if cmd, ok := ud.Value.(*exec.Cmd); ok {
			return cmd
		}
	}

	L.Push(command.RawGetString("path"))
	defer L.Pop(1)
	path := L.CheckString(-1)

	var arg []string
	L.Push(command.RawGetString("args"))
	defer L.Pop(1)
	args := L.CheckTable(-1)
	args.ForEach(func(_ lua.LValue, value lua.LValue) {
		arg = append(arg, lua.LVAsString(value))
	})

	cmd := exec.Command(path, arg...)

	if env := command.RawGetString("env"); env != lua.LNil {
		L.Push(env)
		defer L.Pop(1)
		L.CheckTable(-1).ForEach(func(_ lua.LValue, value lua.LValue) {
			cmd.Env = append(cmd.Env, lua.LVAsString(value))
		})
	}
	if dir := command.RawGetString("dir"); dir != lua.LNil {
		L.Push(dir)
		defer L.Pop(1)
		cmd.Dir = L.CheckString(-1)
	}
	if stdin := command.RawGetString("stdin"); stdin != lua.LNil {
		L.Push(stdin)
		defer L.Pop(1)
		cmd.Stdin = lio.CheckIOReader(L, -1)
	}
	if stdout := command.RawGetString("stdout"); stdout != lua.LNil {
		L.Push(stdout)
		defer L.Pop(1)
		cmd.Stdout = lio.CheckIOWriter(L, -1)
	}
	if stderr := command.RawGetString("stdout"); stderr != lua.LNil {
		L.Push(stderr)
		defer L.Pop(1)
		cmd.Stderr = lio.CheckIOWriter(L, -1)
	}

	return cmd
}

func lvCmd(L *lua.LState, cmd *exec.Cmd) lua.LValue {
	ud := L.NewUserData()
	ud.Value = cmd
	return ud
}

func cmdRun(L *lua.LState) int {
	cmd := checkCmd(L, 1)
	if err := cmd.Run(); err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}

func cmdStart(L *lua.LState) int {
	command := L.CheckTable(1)
	cmd := checkCmd(L, 1)
	command.RawSetString("cmd", lvCmd(L, cmd))

	if err := cmd.Start(); err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}

func cmdWait(L *lua.LState) int {
	cmd := checkCmd(L, 1)
	if err := cmd.Wait(); err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
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

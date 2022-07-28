package exec

import (
	lua "github.com/yuin/gopher-lua"
)

// Preload exec library
func Preload(L *lua.LState) {
	L.PreloadModule("exec", Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {
	registerCmd(L)
	exec := L.SetFuncs(L.NewTable(), api)
	L.Push(exec)
	return 1
}

var api = map[string]lua.LGFunction{
	"command": cmdCommand,
}

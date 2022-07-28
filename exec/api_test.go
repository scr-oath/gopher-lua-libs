package exec

import (
	"github.com/stretchr/testify/assert"
	"github.com/vadv/gopher-lua-libs/inspect"
	"github.com/vadv/gopher-lua-libs/strings"
	"github.com/vadv/gopher-lua-libs/tests"
	"testing"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		strings.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}

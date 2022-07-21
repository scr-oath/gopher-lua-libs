package cert_util

import (
	"github.com/stretchr/testify/assert"
	"github.com/vadv/gopher-lua-libs/tests"
	"log"
	"net/http"
	"testing"
	"time"
)

func runHttps(addr string) {
	err := http.ListenAndServeTLS(addr, "./test/cert.pem", "./test/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func httpRouterGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`OK`))
}

func TestApi(t *testing.T) {

	http.HandleFunc("/get", httpRouterGet)
	go runHttps(":1443")
	time.Sleep(time.Second)

	assert.NotZero(t, tests.RunLuaTestFile(t, Preload, "./test/test_api.lua"))
}

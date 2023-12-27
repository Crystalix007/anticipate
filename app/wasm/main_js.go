//go:build js && wasm

package main

import (
	"net/http"

	"github.com/Crystalix007/anticipate/app"
	"github.com/Crystalix007/anticipate/lib"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	app := app.New(app.WithLogger(lib.NoopLogger))

	routes := app.DeclareRoutes()
	mux := routes.Router()

	release := wasmhttp.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	}), wasmhttp.WithStripPrefix(false))

	defer release()

	select {}
}

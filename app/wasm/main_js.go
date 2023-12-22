//go:build js && wasm

package main

import (
	"net/http"

	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	wasmhttp.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World from Go!"))
	}))

	select {}
}

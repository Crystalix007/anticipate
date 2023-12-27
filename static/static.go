package static

import (
	"embed"

	"github.com/Crystalix007/anticipate/lib"
)

//go:embed js/*
var JSFS embed.FS
var JS = lib.ServeStatic("js", JSFS)

//go:embed favicon.ico
var FaviconFS embed.FS
var Favicon = lib.ServeStatic("", FaviconFS)

//go:embed wasm/*
var WASMFS embed.FS
var WASM = lib.ServeStatic("wasm", WASMFS)

// Package templates provides a filesystem for all the templates.
package templates

import (
	"embed"
)

//go:embed *
var FS embed.FS

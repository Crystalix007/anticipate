package app

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/Crystalix007/anticipate/lib"
)

// ServeAPI serves the API by rendering the specified HTML template with the
// given data.
// It takes an http.ResponseWriter, an http.Request, and a map[string]any as
// parameters.
// The path of the HTML template is derived from the URL path of the request.
// It returns an error if there is an issue serving the template.
func (a *App) ServeAPI(w http.ResponseWriter, r *http.Request, bind map[string]any) error {
	path := strings.TrimPrefix(r.URL.Path, "/api/") + ".html"

	a.logger.InfoContext(r.Context(), "serving api", slog.String("path", path))

	return lib.ServeTemplateHTML(w, path, bind)
}

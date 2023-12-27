package app

import (
	"net/http"

	"github.com/Crystalix007/anticipate/lib"
)

// ServeIndex serves the index page.
// It takes a http.ResponseWriter and a http.Request as parameters.
// It returns an error if there was a problem serving the template.
// The function uses the lib.ServeTemplateHTML function to render the
// "index.html" template.
// The template is rendered with no additional data.
func ServeIndex(w http.ResponseWriter, r *http.Request) error {
	return lib.ServeTemplateHTML(w, "index.html", nil)
}

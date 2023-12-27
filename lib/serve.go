package lib

import (
	"embed"
	"fmt"
	"net/http"

	"html/template"

	"github.com/Crystalix007/anticipate/templates"
	"github.com/go-chi/chi"
)

// ServeTemplateHTML renders a template based on the request path and binds it
// with the provided data.
// It takes a fiber.Ctx object and a bind parameter as input.
// The request path is used to determine the template to render.
// The bind parameter is used to bind data to the template.
// It returns an error if there is any issue rendering the template.
func ServeTemplateHTML(w http.ResponseWriter, path string, bind any) error {
	template, err := template.ParseFS(templates.FS, path)
	if err != nil {
		return fmt.Errorf(
			"lib: couldn't parse template from templates static FS: %w",
			err,
		)
	}

	if err := template.Execute(w, bind); err != nil {
		return fmt.Errorf(
			"lib: couldn't execute template: %w",
			err,
		)
	}

	return nil
}

// ServeStatic returns an http.Handler that serves static files on the
// specified path prefix, under the provided subdirectory, using the provided
// embed.FS.
func ServeStatic(
	subdirectory string,
	fs embed.FS,
) http.Handler {
	fileserver := http.FileServer(http.FS(fs))
	mux := chi.NewMux()

	mux.Use(AddPrefix(subdirectory))
	mux.Get("/*", fileserver.ServeHTTP)

	return mux
}

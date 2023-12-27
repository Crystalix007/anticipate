package serve

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Crystalix007/anticipate/app"
	"github.com/Crystalix007/anticipate/lib"
	"github.com/Crystalix007/anticipate/static"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

const (
	// Subcommand is the name of the subcommand.
	Subcommand = "serve"

	// Short is the short description of the subcommand.
	Short = "Serve the application"
)

// Command returns a new instance of the serve command.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   Subcommand,
		Short: Short,
		RunE:  Serve,
	}

	cmd.Flags().BoolP("verbose", "v", false, "verbose output")

	return cmd
}

// Serve is a function that starts the server and handles incoming requests.
// It initializes the HTML template engine, sets up routes, and listens on port 8080.
// The function returns an error if there is an issue starting the server.
func Serve(cmd *cobra.Command, _ []string) error {
	mux := chi.NewMux()

	if verbose, err := cmd.Flags().GetBool("verbose"); err == nil && verbose {
		mux.Use(middleware.Logger)
	}

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.CleanPath)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		lib.ErrHandler(w, r, lib.ServeTemplateHTML(w, "index.html", nil))
	})

	mux.Mount("/static/js/", http.StripPrefix("/static/js/", static.JS))
	mux.Mount("/static/wasm/", http.StripPrefix("/static/wasm/", static.WASM))

	mux.Get("/favicon.ico", static.Favicon.ServeHTTP)
	mux.Get("/sw.js", static.JS.ServeHTTP)

	logger := slog.New(slog.NewTextHandler(cmd.ErrOrStderr(), nil))
	errLogger := lib.LogErrors(logger)
	app := app.New(app.WithLogger(logger))

	for _, route := range app.DeclareRoutes() {
		mux.Method(route.Method, route.Path, lib.FuncErrHandler(
			errLogger(route.Handler),
		))
	}

	defer app.Close()

	fmt.Fprintf(
		cmd.ErrOrStderr(),
		"Listening on port 8080 with %d routes\n\n",
		len(mux.Routes()),
	)

	return http.ListenAndServe(":8080", mux)
}

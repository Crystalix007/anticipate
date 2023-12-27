package app

import (
	"log/slog"
	"net/http"

	"github.com/Crystalix007/anticipate/lib"
)

type App struct {
	comments []string

	logger *slog.Logger
}

// Ensure [App] implements [lib.App].
var _ lib.App = &App{}

// New creates a new instance of the app with the provided options.
// It accepts optional configuration options as variadic arguments.
func New(opts ...Option) *App {
	a := &App{}

	for _, opt := range opts {
		opt(a)
	}

	a.setDefaults()

	return a
}

// DeclareRoutes declares the routes for the App.
// It returns a slice of lib.Route containing the routes.
func (a *App) DeclareRoutes() lib.Routes {
	return lib.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/api/comments",
			Handler: a.ShowComments,
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/comments",
			Handler: a.AddComment,
		},
	}
}

// Close is a method of the App struct that is responsible for closing the application.
// It returns an error if there was a problem during the closing process.
func (*App) Close() error {
	return nil
}

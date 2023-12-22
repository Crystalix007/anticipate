package app

import (
	"net/http"

	"github.com/Crystalix007/anticipate/lib"
)

type App struct {
	comments []string
}

var _ lib.App = &App{}

// DeclareRoutes declares the routes for the App.
// It returns a slice of lib.Route containing the routes.
func (a *App) DeclareRoutes() []lib.Route {
	return []lib.Route{
		{
			Method:  http.MethodGet,
			Path:    "/comments",
			Handler: a.ShowComments,
		},
		{
			Method:  http.MethodPost,
			Path:    "/comments",
			Handler: a.AddComment,
		},
		{
			Method:  http.MethodGet,
			Path:    "/sw.js",
			Handler: a.ServeSW,
		},
	}
}

// Close is a method of the App struct that is responsible for closing the application.
// It returns an error if there was a problem during the closing process.
func (*App) Close() error {
	return nil
}

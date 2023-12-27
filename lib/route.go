package lib

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Handler represents a function that handles HTTP requests.
// It takes in a http.ResponseWriter and a http.Request as parameters,
// and returns an error if any occurred during the handling process.
type Handler func(w http.ResponseWriter, r *http.Request) error

// Method represents the HTTP method of the route.
// Path represents the URL path of the route.
// Handler is the function that handles the route request.
// Middleware is an optional list of middleware functions that are applied to
// the route.
type Route struct {
	Method     string
	Path       string
	Handler    Handler
	Middleware []Middleware
}

// Mount mounts the Route onto the provided chi.Router.
// If the Route has no middleware, it directly registers the handler function
// with the specified method and path.
// If the Route has middleware, it creates a new router group, applies the
// middleware, and then registers the handler function with the specified
// method and path.
func (r Route) Mount(router chi.Router) {
	if len(r.Middleware) == 0 {
		router.Method(r.Method, r.Path, FuncErrHandler(r.Handler))
	} else {
		router.Group(func(router chi.Router) {
			router.Use(r.Middleware...)
			router.Method(r.Method, r.Path, FuncErrHandler(r.Handler))
		})
	}
}

// Routes represents a collection of individual [Route]s.
type Routes []Route

// Router returns a chi.Router that is configured with the routes defined in
// the Routes slice.
// It iterates over each route in the Routes slice and mounts them onto the
// router.
// The configured router is then returned.
func (r Routes) Router() chi.Router {
	router := chi.NewRouter()

	for _, route := range r {
		route.Mount(router)
	}

	return router
}

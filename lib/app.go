package lib

// App represents an application interface.
type App interface {
	// DeclareRoutes declares the routes for the application.
	DeclareRoutes() []Route

	// Close closes the application.
	Close() error
}

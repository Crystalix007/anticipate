package lib

import (
	"log/slog"
	"net/http"
	"path"
)

// Middleware represents a function that wraps an http.Handler with additional
// functionality.
type Middleware = func(http.Handler) http.Handler

// ErrMiddleware adds error handling to the [Middleware] type.
type ErrMiddleware = func(Handler) Handler

// StripPrefix returns a middleware that strips the given prefix from the
// request URL path before passing it to the next handler.
// The prefix parameter specifies the prefix to be stripped.
// The returned middleware function takes an http.Handler as input and returns
// a new http.Handler that performs the prefix stripping.
func StripPrefix(prefix string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.StripPrefix(prefix, h)
	}
}

// AddPrefix is a middleware that adds a prefix to the URL path of incoming
// requests.
// It takes a prefix string as input and returns a Middleware function.
func AddPrefix(prefix string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = path.Join(prefix, r.URL.Path)

			h.ServeHTTP(w, r)
		})
	}
}

// LogErrors is a middleware function that logs any errors that occur during
// request handling.
// The modified Handler function logs any errors that occur during request
// handling using the provided logger.
// If an error occurs, it is passed to the logger.ErrorContext method along
// with additional context information.
// The original error is then returned.
func LogErrors(logger *slog.Logger) ErrMiddleware {
	return func(h Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			err := h(w, r)
			if err != nil {
				logger.ErrorContext(r.Context(), "error handling request", slog.Any("error", err))
			}

			return err
		}
	}
}

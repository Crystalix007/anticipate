package lib

import (
	"net/http"

	"github.com/go-chi/render"
)

// FuncErrHandler converts a Handler into a http.Handler, by wrapping it with a
// function that handles errors and converts them to logged errors, and 500
// HTTP responses.
func FuncErrHandler(handler Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrHandler(w, r, handler(w, r))
	})
}

// ErrHandler handles errors by writing an HTTP 500 status code and the error
// message to the response.
// It takes in the response writer, the request, and the error as parameters.
// If the error is not nil, it writes the error message to the response writer.
func ErrHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
	}
}

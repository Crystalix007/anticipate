package lib

import (
	"context"
	"log/slog"
)

// NoopHandler represents a logger handler implementation that does nothing.
type NoopHandler struct {
	slog.TextHandler
}

// Ensure [PNoopHandler] implements [slog.Handler].
var _ slog.Handler = &NoopHandler{}

// Enabled returns a boolean value indicating whether logging is enabled for
// the given context and log level.
// It always returns false for the NoopLogger implementation.
func (NoopHandler) Enabled(context.Context, slog.Level) bool {
	return false
}

// NoopLogger is a logger that discards all log messages.
var NoopLogger = slog.New(&NoopHandler{})

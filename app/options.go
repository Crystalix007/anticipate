package app

import (
	"log/slog"
	"os"
)

// Option is a function that can be used to configure an App.
type Option = func(a *App)

// WithLogger sets the logger for the App.
// It takes a pointer to a slog.Logger as input and returns an Option function.
// The Option function sets the logger of the App to the provided logger.
// Example usage:
//
//	app := NewApp(WithLogger(logger))
func WithLogger(logger *slog.Logger) Option {
	return func(a *App) {
		a.logger = logger
	}
}

// setDefaults sets the default values for the App struct.
// If the logger is not set, it creates a new logger with a default text
// handler that writes to os.Stderr.
func (a *App) setDefaults() {
	if a.logger == nil {
		a.logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
}

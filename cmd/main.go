package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Crystalix007/anticipate/cmd/serve"
	"github.com/spf13/cobra"
)

const (
	// Version is the current version of the application.
	Version = "0.0.1"

	// Name is the name of the application.
	Name = "anticipate"
)

// ErrNoCommand is an error indicating that no command was provided.
var ErrNoCommand = errors.New("no command provided")

func main() {
	cmd := cobra.Command{
		Use:     Name,
		Version: Version,
		RunE:    Fallback,
	}

	cmd.AddCommand(serve.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)

		os.Exit(1)
	}
}

// Fallback is a function that is called when no command is provided.
// It returns an error of type ErrNoCommand.
func Fallback(_ *cobra.Command, _ []string) error {
	return ErrNoCommand
}

package serve

import (
	"fmt"
	"net/http"

	"github.com/Crystalix007/anticipate/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
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
	engine := html.New("./templates", ".tmpl")
	mux := fiber.New(fiber.Config{
		Views: engine,
	})

	if verbose, err := cmd.Flags().GetBool("verbose"); err == nil && verbose {
		mux.Use(logger.New())
	}

	mux.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	mux.Get("/static/+", func(c *fiber.Ctx) error {
		return c.SendFile(fmt.Sprintf("./static/%s", c.Params("+")))
	})

	mux.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.Redirect("/static/favicon.ico", http.StatusMovedPermanently)
	})

	app := &app.App{}

	for _, route := range app.DeclareRoutes() {
		mux.Add(route.Method, route.Path, route.Handler)
	}

	defer app.Close()

	return mux.Listen(":8080")
}

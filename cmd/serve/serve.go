package serve

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	return &cobra.Command{
		Use:   Subcommand,
		Short: Short,
		RunE:  Serve,
	}
}

// Serve is a function that starts the server and handles incoming requests.
// It initializes the HTML template engine, sets up routes, and listens on port 8080.
// The function returns an error if there is an issue starting the server.
func Serve(_ *cobra.Command, _ []string) error {
	engine := html.New("./templates", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/static/+", func(c *fiber.Ctx) error {
		return c.SendFile(fmt.Sprintf("./static/%s", c.Params("+")))
	})

	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.Redirect("/static/favicon.ico", http.StatusMovedPermanently)
	})

	return app.Listen(":8080")
}

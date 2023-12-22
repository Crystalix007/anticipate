package lib

import "github.com/gofiber/fiber/v2"

// Method represents the HTTP method of the route.
// Path represents the URL path of the route.
// Handler is the function that handles the route request.
type Route struct {
	Method  string
	Path    string
	Handler func(c *fiber.Ctx) error
}

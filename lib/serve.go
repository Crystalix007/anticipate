package lib

import (
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Serve(c *fiber.Ctx, bind any) error {
	reqPath := strings.TrimPrefix(c.Path(), "/")
	template := strings.TrimSuffix(reqPath, path.Ext(reqPath))

	return c.Render(template, bind)
}

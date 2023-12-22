package app

import (
	"strings"

	"github.com/Crystalix007/anticipate/lib"
	"github.com/gofiber/fiber/v2"
)

func (a *App) ShowComments(c *fiber.Ctx) error {
	return lib.Serve(c, map[string]any{
		"comments": a.comments,
	})
}

func (a *App) AddComment(c *fiber.Ctx) error {
	comment := strings.Clone(c.FormValue("comment"))
	a.comments = append(a.comments, comment)

	return a.ShowComments(c)
}

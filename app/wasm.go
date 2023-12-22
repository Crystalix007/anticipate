package app

import "github.com/gofiber/fiber/v2"

func (a *App) ServeSW(c *fiber.Ctx) error {
	return c.SendFile("static/js/sw.js")
}

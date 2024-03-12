package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Routes(app *fiber.App) {
	// default route web
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
			"author": "Nopaleon Bonaparte",
			"feature": c.Route().Name,
		})
	}).Name("Home web Endpoint")

	app.Get("/stats", monitor.New(
		monitor.Config{
			Title: "Turu API monitor page",
		},
	)).Name("Monitor API")
}
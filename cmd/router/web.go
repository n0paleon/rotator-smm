package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Web(app *fiber.App) {
	// default route web
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"title": "hello world",
		}, "layouts/main")
	}).Name("Home web Endpoint")

	app.Get("/stats", monitor.New(
		monitor.Config{
			Title: "Turu API monitor page",
			Refresh: 1 * time.Second,
		},
	)).Name("Monitor API")
}
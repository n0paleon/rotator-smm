package api

import "github.com/gofiber/fiber/v2"

type TestMessage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"response_code"`
}

func Routes(app *fiber.App) {
	// inisialisasikan /api sebagai endpoint masuk route group 'api'
	api := app.Group("/api")

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(TestMessage{
			Message: "Welcome to /api routes!",
			Status: "success",
			Code: 200,
		})
	}).Name("Test API routes")
}
package router

import (
	"RotatorSMM/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TestMessage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"response_code"`
}

type Users struct {
	ID uint
	Name string
	Age int
}

func Api(app *fiber.App) {
	// inisialisasikan /api sebagai endpoint masuk route group 'api'
	api := app.Group("/api")

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(TestMessage{
			Message: "Welcome to /api routes!",
			Status: "success",
			Code: 200,
		})
	}).Name("Test API routes")

	api.Get("/users", func(c *fiber.Ctx) error {
		var users []Users

		if err := database.DB.Find(&users).Error; err != nil {
			log.Printf("Failed to fetch users: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to fetch users",
			})
		}

		return c.JSON(fiber.Map{
			"users": users,
		})
	})
}
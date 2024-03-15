package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"

	services "RotatorSMM/cmd/router"
	"RotatorSMM/configs"
	"RotatorSMM/pkg/database"
	"RotatorSMM/pkg/database/migrations"
)


func Start () {
	// load database
	database.Init()
	// migrate database
	migrations.Users()
	// close database
	defer database.Close()

	// set up new fiber instance
	app := fiber.New(configs.Fiber())

	// set static route and directory
	app.Static("/static", "./assets")

	// web service handler 
	services.Web(app)
	// api service handler
	services.Api(app)

	// start application
	log.Fatal(app.Listen(":3000"))
}
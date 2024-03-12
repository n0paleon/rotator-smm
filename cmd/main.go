package cmd

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"RotatorSMM/api"
	"RotatorSMM/web"
)


func Start () {
	app := fiber.New(fiber.Config{
		AppName: "TuruAPI v1.1",
		ServerHeader: "Golang",
		CaseSensitive: true,
		EnablePrintRoutes: true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api.Routes(app)
	web.Routes(app)

	log.Fatal(app.Listen(":3000"))
}
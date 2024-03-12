package configs

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)


func Fiber () fiber.Config {
	// set up new fiber instance
	return fiber.Config{
		AppName: "TuruAPI v1.1",
		ServerHeader: "Golang",
		CaseSensitive: true,
		EnablePrintRoutes: true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork: false, // set true if you want the app run in multi process
		StrictRouting: true,
		WriteTimeout: 30 * time.Second,
		Views: handlebars.New("./web/templates", ".hbs"),
	}
}
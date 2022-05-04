package main

import (
	"log"
	"time"

	"github.com/faizan-glitch/url-sh/pkg/config"
	"github.com/faizan-glitch/url-sh/pkg/database"
	"github.com/faizan-glitch/url-sh/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	config.Load()

	database.Connect()
}

func main() {

	app := fiber.New(fiber.Config{
		Prefork:                      true,
		BodyLimit:                    0.5 * 1024 * 1024,
		DisableKeepalive:             true,
		DisableDefaultDate:           true,
		AppName:                      "url-sh",
		DisablePreParseMultipartForm: true,
		WriteTimeout:                 time.Millisecond * 100,
		ReadTimeout:                  time.Millisecond * 100,
	})

	routes.RegisterUrlRoutes(app)

	log.Fatal(app.Listen(viper.GetString("PORT")))
}

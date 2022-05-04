package routes

import (
	"github.com/faizan-glitch/url-sh/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUrlRoutes(app *fiber.App) {

	app.Post("/shorten", controllers.Shorten)

	app.Get("/:short", controllers.Resolve)

}

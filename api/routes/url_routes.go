package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jesseinvent/go-fiber-redis-url-shortener/controllers"
)

func UrlRoutes(app *fiber.App)  {
	app.Post("/api/v1/create_link", controllers.ShortenUrl);
	app.Get("/:url", controllers.ResolveUrl);
}

package controllers

import(
	"github.com/gofiber/fiber/v2"
	"github.com/jesseinvent/go-fiber-redis-url-shortener/database"
)

func ResolveUrl(c *fiber.Ctx) error {

	url := c.Params("url");

	value, err := database.Get(url);

	if value == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Short not found"});
	} else if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Cannot connect to DB"});
	}

	return c.Redirect(value, 301);
}
package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jesseinvent/go-fiber-redis-url-shortener/routes"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New();

	err := godotenv.Load(); 

	if err != nil {
		log.Fatal("Error loading env file:", err);
	}

	app.Use(logger.New());

	routes.UrlRoutes(app);

	PORT := os.Getenv("PORT");

	err2 := app.Listen(PORT);

	if err2 != nil {
		log.Fatal(err);
	}

}
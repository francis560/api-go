package main

import (
	"log"

	"github.com/francis560/api-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	routes.RandUsers(app)

	log.Fatal(app.Listen(":3000"))

}

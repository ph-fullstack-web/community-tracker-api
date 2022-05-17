package main

import (
	"github.com/RegieCanadaRC/community-tracker-api/database"
	"github.com/RegieCanadaRC/community-tracker-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 8000
	app.Listen(":8000")
}

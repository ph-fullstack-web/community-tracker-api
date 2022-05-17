package router

import (
	communityRoutes "github.com/RegieCanadaRC/community-tracker-api/internal/routes/community"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	communityRoutes.SetupCommunityRoutes(api)
}

package communityRoutes

import (
	communityHandler "github.com/RegieCanadaRC/community-tracker-api/internal/handlers/community"
	"github.com/gofiber/fiber/v2"
)

func SetupCommunityRoutes(router fiber.Router) {
	community := router.Group("/community")
	// Create a community
	community.Post("/", communityHandler.AddCommunity)
	// Read all community
	community.Get("/", communityHandler.GetCommunity)
	// Read one community
	community.Get("/:communityid", communityHandler.GetCommunityById)
	// Update one community
	//community.Put("/:communityid", func(c *fiber.Ctx) error {})
	// Delete one community
	//community.Delete("/:communityid", func(c *fiber.Ctx) error {})
}

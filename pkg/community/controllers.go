package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	communityRoutes := app.Group("/api/community", middleware.AuthMiddleware)
	communityRoutes.Get("/job-level", h.GetJobLevel)
	communityRoutes.Get("/", h.GetCommunity)
	communityRoutes.Get("/percentage", h.GetCommunityWithmembersPercentage)
	communityRoutes.Post("/", h.AddCommunity)
	communityRoutes.Put("/:communityid", h.UpdateCommunity)
	communityRoutes.Get("/:communityid", h.GetCommunityById)
}

package community

import (
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
	communityRoutes := app.Group("/api/community")
	//communityRoutes.Get("/job-level", h.GetJobLevel)
	communityRoutes.Get("/", h.GetCommunity)
	communityRoutes.Get("/:Communityid", h.GetCommunityId)
	communityRoutes.Post("/", h.AddCommunity)
	communityRoutes.Put("/:UpdateCommunity", h.UpdateCommunity)
	communityRoutes.Put("/:DeleteCommunity", h.DeleteCommunity)
}

package community_managers

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
	routes := app.Group("/api/managers", middleware.AuthMiddleware)
	routes.Get("/", h.GetManagers)
	routes.Get("/community", h.GetCommunityManagers)
}

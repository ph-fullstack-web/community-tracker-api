package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/VncntDzn/community-tracker-api/pkg/middleware"
	"gorm.io/gorm"
	
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	route := app.Group("/api/admin", middleware.AuthMiddleware)
	route.Post("/", h.CreateAdmin)
	route.Put("/:communityadminandmanagerid", h.UpdateAdminDetails)
	route.Put("/:communityadminandmanagerid/password", h.UpdatePassword)
}

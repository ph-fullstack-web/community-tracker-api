package projects

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

	projectRoutes := app.Group("/api/projects", middleware.AuthMiddleware)
	projectRoutes.Post("/", h.CreateProject)
	projectRoutes.Put("/:projectid", h.UpdateProject)
	projectRoutes.Get("/", h.GetProjects)
	projectRoutes.Get("/:projectid", h.GetProject)
	projectRoutes.Delete("/:projectid", h.DeleteProject)
}

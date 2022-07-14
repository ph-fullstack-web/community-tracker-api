package projects

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

	projectRoutes := app.Group("/api/projects")
	projectRoutes.Post("/", h.CreateProject)
	projectRoutes.Get("/", h.GetProjects)
}

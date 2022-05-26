package community_managers

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
	citiesRoutes := app.Group("/api/managers")
	citiesRoutes.Get("/", h.GetManagers)
}

package cities

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
	citiesRoutes := app.Group("/api/cities", middleware.AuthMiddleware)
	citiesRoutes.Get("/", h.GetCities)
	citiesRoutes.Post("/", h.AddCity)
	citiesRoutes.Delete("/:id", h.DeleteCity)
	citiesRoutes.Put("/:city_id", h.UpdateCity)
}

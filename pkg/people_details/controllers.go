package people_details

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
	route := app.Group("/api/peopledetails", middleware.AuthMiddleware)
	route.Get("/description", h.GetPeopleDetailsDesc)
	route.Get("/", h.GetPeopleDetails)
}

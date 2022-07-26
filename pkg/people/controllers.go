package people

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
	peopleRoutes := app.Group("/api/people")
	peopleRoutes.Get("/", h.GetPeople)
	peopleRoutes.Get("/:people_id", h.GetPeopleById)
	peopleRoutes.Post("/", middleware.AuthMiddleware, h.AddPeople)
	peopleRoutes.Put("/:peopleid", middleware.AuthMiddleware, h.UpdatePeople)
	peopleRoutes.Post("/skills", middleware.AuthMiddleware, h.GetPeopleBySkills)
}

package people

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
	communityRoutes := app.Group("/api/people")
	communityRoutes.Get("/", h.GetPeople)
	//communityRoutes.Post("/", h.AddPeople)
	//communityRoutes.Put("/:peopleid", h.UpdatePeople)
	communityRoutes.Get("/:people_id", h.GetPeopleById)
}

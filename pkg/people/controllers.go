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
	peopleRoutes := app.Group("/api/people")
	peopleRoutes.Get("/", h.GetPeople)
	peopleRoutes.Get("/:people_id", h.GetPeopleById)
	peopleRoutes.Post("/skills", h.GetPeopleBySkills);
	peopleRoutes.Post("/", h.AddPeople)
	peopleRoutes.Put("/:peopleid", h.UpdatePeople)

	peopleskillsRoutes := app.Group("/api/peopleskills")
	peopleskillsRoutes.Post("/", h.AddPeopleSkills)
	peopleskillsRoutes.Get("/", h.GetPeopleSkills)
	peopleskillsRoutes.Put("/:peopleskillsid", h.DeletePeople)
	peopleskillsRoutes.Delete("/:peopleskillsid", h.DeleteSkills)
}

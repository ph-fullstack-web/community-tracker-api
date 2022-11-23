package people_skills

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
	route := app.Group("/api/peopleskills", middleware.AuthMiddleware)
	route.Post("/", h.AddPeopleSkills)
	route.Get("/", h.GetPeopleSkills)
	route.Put("/:peopleskillsid", h.UpdateSkill)
	route.Delete("/:peopleskillsid", h.DeleteSkills)
}

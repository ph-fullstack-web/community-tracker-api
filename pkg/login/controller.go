package login

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
	route := app.Group("/api")
	route.Post("/login", h.Login)
	route.Post("/googleLogin", h.GoogleLogin)
}

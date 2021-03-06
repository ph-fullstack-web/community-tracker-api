package cities

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddCityRequestBody struct {
	Name string `json:"name"`
}

func (h handler) AddCity(c *fiber.Ctx) error {
	body := AddCityRequestBody{
		Name: "",
	}

	// parse body, attach to AddCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var city models.AddCity

	city.Name = body.Name

	// insert new db entry
	if result := h.DB.Create(&city); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success! Added Data!", "data": &city})
}

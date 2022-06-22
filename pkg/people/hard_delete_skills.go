package people

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteCity(c *fiber.Ctx) error {
	id := c.Params("peopleskillsid")

	var peopleskills models.Hard_Delete_Skills

	trim_id := strings.TrimLeft(id, "peopleskillsid=")

	if result := h.DB.First(&peopleskills, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete city from db
	h.DB.Where("peopleskillsid = ?", trim_id).Delete(&peopleskills)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success Deleted!", "data": &peopleskills})
}

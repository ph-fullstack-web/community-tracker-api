package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCommunityById(c *fiber.Ctx) error {
	var Community models.Community

	id := c.Params("communityid")

	if result := h.DB.Where("isactive = ?", true).Preload("Manager").First(&Community, "communityid = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": &Community})
}

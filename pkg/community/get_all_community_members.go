package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAllCommunityMembers(c *fiber.Ctx) error {
	var member_data []models.People

	if result := h.DB.Find(&member_data); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&member_data)
}

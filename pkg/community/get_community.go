package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCommunity(c *fiber.Ctx) error {
	var community_data []models.Community

	if result := h.DB.Order("lower(communityname)").Where("isactive = ?", true).Preload("Manager").Find(&community_data); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&community_data)
}

package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCommunityId(c *fiber.Ctx) error {
	id := c.Params("Communityid")

	var CommunityId models.Community

	h.DB.Find(&CommunityId, "Communityid = ?", id)

	if result := h.DB.First(&CommunityId, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Data Found", "data": CommunityId})
}

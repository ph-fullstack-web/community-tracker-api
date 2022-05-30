package people

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPeopleById(c *fiber.Ctx) error {
	var PeopleId models.People

	id := c.Params("people_id")

	h.DB.First(&PeopleId, "peopleid = ?", id)

	if result := h.DB.First(&PeopleId, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Data Found", "data": PeopleId})
}

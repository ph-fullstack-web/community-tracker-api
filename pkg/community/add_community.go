package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddCommunityRequestBody struct {
	Communityname        string `json:"communityname"`
	Communitymgrpeopleid int    `json:"communitymanagerpeopleid"`
	Isactive             bool   `json:"isactive"`
}

func (h handler) AddCommunity(c *fiber.Ctx) error {
	body := AddCommunityRequestBody{
		Communityname:        "",
		Communitymgrpeopleid: 0,
		Isactive:             false,
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var community models.Community

	community.Communityname = body.Communityname
	community.Communitymgrpeopleid = body.Communitymgrpeopleid
	community.Isactive = body.Isactive

	if result := h.DB.Create(&community); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success! Added Data!", "data": &community})

}

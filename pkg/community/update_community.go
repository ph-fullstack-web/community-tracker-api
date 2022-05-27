package community

import (
	"fmt"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateCommunityRequestBody struct {
	Communityname        string `json:"communityname"`
	Communitymgrpeopleid int    `json:"communitymanagerpeopleid"`
	Isactive             bool   `json:"isactive"`
}

func (h handler) UpdateCommunity(c *fiber.Ctx) error {
	id := c.Params("UpdateCommunity")
	body := UpdateCommunityRequestBody{}

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var community models.Community

	community.Communityname = body.Communityname
	community.Communitymgrpeopleid = body.Communitymgrpeopleid
	community.Isactive = body.Isactive

	if result := h.DB.First(&community, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	} else {
		community.Communityname = body.Communityname
		community.Communitymgrpeopleid = body.Communitymgrpeopleid
		community.Isactive = body.Isactive
		h.DB.Save(&community)
		fmt.Println(result)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &community})
	}

}

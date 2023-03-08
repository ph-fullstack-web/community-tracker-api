package community_members

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) GetCommunityMembersSearch(c *fiber.Ctx) error {
	var community_data models.PeopleUnderCommunitySearch
	// var rows = strings.TrimSpace(c.Query("rows"))
	rows, rowsError := strconv.Atoi(strings.TrimSpace(c.Query("rows")))
	if rowsError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}

	// var page = strings.TrimSpace(c.Query("page"))
	page, pageError := strconv.Atoi(strings.TrimSpace(c.Query("page")))
	if pageError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}

	communityId, communityIdError := strconv.Atoi(c.Params("communityId"))
	if communityIdError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}

	fmt.Println("Rows", rows)
	fmt.Println("Page", page)
	result := h.DB.Order("c.communityid, p.fullname").Limit(rows).Offset((page- 1) * rows).Table("people p").Select("p.peopleid, p.fullname, p.communityid, c.communityname, c.communityid").Joins("left join community c on c.communityid = p.communityid").Where(&models.PeopleUnderCommunitySearch{Communityid: communityId})

	// show 404 error if no community found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Community does not exist."})
	}

	// check for other errors
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &community_data})
}

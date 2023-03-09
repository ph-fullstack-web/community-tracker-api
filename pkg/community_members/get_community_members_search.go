package community_members

import (
	"errors"
	"strconv"
	"strings"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) GetCommunityMembersSearch(c *fiber.Ctx) error {
	var community_data models.CommunityMembersSearch
	var members_count []models.PeopleUnderCommunitySearch
	var count int64

	communityId, communityIdError := strconv.Atoi(c.Params("communityId"))
	if communityIdError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Path Parameter is invalid"})
	}

	var rowsStr = strings.TrimSpace(c.Query("rows"))
	rows, rowsError := strconv.Atoi(rowsStr)
	if rowsError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Query Parameter is invalid"})
	}
	
	var pageStr = strings.TrimSpace(c.Query("page"))
	page, pageError := strconv.Atoi(pageStr)
	if pageError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Query Parameter is invalid"})
	}

	h.DB.Where(&models.PeopleUnderCommunitySearch{Communityid: communityId, Isactive: true}).Find(&members_count).Count(&count)
	result := h.DB.Where(&models.CommunityMembersSearch{CommunityID: communityId}).Preload("Members", func(tx *gorm.DB) *gorm.DB { return tx.Limit(rows).Offset((page - 1) * rows).Where(models.PeopleUnderCommunitySearch{Isactive: true})}).Preload("Manager").First(&community_data)
	
	// show 404 error if no community found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Community does not exist."})
	}

	// check for other errors
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &community_data, "totalCount": &count})
}

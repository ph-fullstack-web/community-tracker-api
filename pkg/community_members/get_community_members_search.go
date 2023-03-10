package community_members

import (
	"errors"
	"strconv"
	"strings"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)    

type CommunityMembersSearchResponse struct {
	CurrentPage int `gorm:"-" json:"current_page"`
	LastPage int `gorm:"-" json:"last_page"`
	Community models.CommunityMembersSearch `gorm:"-" json:"community"`
}

func getParameterValue(input string)(output string) {
	if(input == "") {
		output = "%"
	}
	output = "%" + input + "%";
	return
}

func (h handler) GetCommunityMembersSearch(c *fiber.Ctx) error {
	var response_data CommunityMembersSearchResponse
	var community_data models.CommunityMembersSearch
	var members_data []models.PeopleUnderCommunitySearch
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

	full_name := getParameterValue(strings.TrimSpace(c.Query("full_name")))
	csv_email := getParameterValue(strings.TrimSpace(c.Query("csv_email")))

	h.DB.Where("communityid = ? and isactive = ? and (fullname ilike ? or csvemail ilike ?)", communityId, true, full_name, csv_email).Find(&members_data).Count(&count)
	result := h.DB.Where(&models.CommunityMembersSearch{CommunityID: communityId}).Preload("Members", func(tx *gorm.DB) *gorm.DB { return tx.Limit(rows).Offset((page - 1) * rows).Order("communityid").Order("fullname").Where("isactive = ? and (fullname ilike ? or csvemail ilike ?)", true, full_name, csv_email)}).Preload("Manager").First(&community_data)
	
	// show 404 error if no community found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Community does not exist."})
	}

	// check for other errors
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}
    
	var lastPageFloat float64 = float64(count) / float64(rows)
	response_data.LastPage = int(math.Ceil(lastPageFloat))
	response_data.CurrentPage = page

	community_data.TotalMembers = int(count)	
	response_data.Community = community_data

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &response_data})
}

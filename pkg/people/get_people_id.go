package people

import (
	"strconv"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPeopleById(c *fiber.Ctx) error {
	var PeopleId models.People
	var skills []models.PeoplePrimarySkills
	var memberSkills []models.SkillSet
	var details []models.PeopleDetails
	var detailsDescriptions []models.PeopleDetailsDesc

	id := c.Params("people_id")

	h.DB.First(&PeopleId, "peopleid = ?", id)

	parsedId, parseError := strconv.Atoi(id)
	if parseError != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Not Found"})
	}

	// Get Member Info
	if result := h.DB.Preload("Community", "isactive = ?", true).Preload("Community.Manager").Where("isactive = ?", true).First(&PeopleId, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Not Found"})
	}

	// Get Skill Ids
	h.DB.Where(&models.PeoplePrimarySkills{PeopleId: parsedId}).Find(&skills)
	var skillsListIds []int
	for _, skillItem := range skills {
		skillsListIds = append(skillsListIds, skillItem.PeopleSkill)
	}

	// Get skills descriptions
	h.DB.Where("peopleskillsid IN ?", skillsListIds).Find(&memberSkills)


	h.DB.Where(&models.PeopleDetails{PeopleId: parsedId}).Find(&details)
	var detailsListIds []int
	for _, detailItem := range details {
		detailsListIds = append(detailsListIds, detailItem.PeopleDetailsDescId)
	}

	h.DB.Where("peopledetailsdescid IN ?", detailsListIds).Find(&detailsDescriptions)


	responseData := models.PeopleWithSkills{
		Peopleid:       PeopleId.Peopleid,
		Cognizantid:    PeopleId.Cognizantid,
		Fullname:       PeopleId.Fullname,
		Csvemail:       PeopleId.Csvemail,
		Hireddate:      PeopleId.Hireddate,
		Communityid:    PeopleId.Communityid,
		Workstateid:    PeopleId.Workstateid,
		Joblevelid:     PeopleId.Joblevelid,
		Projectid:      PeopleId.Projectid,
		Isactive:       PeopleId.Isactive,
		Isprobationary: PeopleId.Isprobationary,
		Community:      PeopleId.Community,
		Skill:          memberSkills,
		Details:		detailsDescriptions,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": responseData})
}

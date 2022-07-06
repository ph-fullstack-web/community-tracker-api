package people

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type GetPeopleRequestBody struct {
	Skills []int `json:"skills"`
}

type PeopleWithSkill struct {
	Fullname	string	`gorm:"column:fullname"`
	PeopleSkillsId	int	`gorm:"column:peopleskillsid"`
	PeopleSkillsDesc	string	`gorm:"column:peopleskillsdesc"`
}

func (h handler) GetPeopleBySkills(c *fiber.Ctx) error {
	body := GetPeopleRequestBody{
		Skills: nil,
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var people []PeopleWithSkill

	
	//subquery between peopleprimaryskills and peopleskills
	sub := h.DB.Table("peopleprimaryskills").Select("peopleprimaryskills.peopleid, peopleskills.peopleskillsid, peopleskills.peopleskillsdesc").Joins("inner join peopleskills on peopleprimaryskills.peopleskillsid = peopleskills.peopleskillsid")

	if result := h.DB.Table("people").Select("fullname, sub.peopleskillsid, sub.peopleskillsdesc").Joins("inner join (?) as sub on sub.peopleid = people.peopleid", sub).Where("sub.peopleskillsid IN (1,2)").Scan(&people); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success!", "data": &people})
}

package people_skills

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type AddPeopleskillsBody struct {
	Peopleskillsdesc string `validate:"required" gorm:"column:peopleskillsdesc" json:"peopleskills_desc"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (h handler) AddPeopleSkills(c *fiber.Ctx) error {

	body := AddPeopleskillsBody{
		Peopleskillsdesc: "",
		IsActive:         true,
	}

	// parse body, attach to struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var peopleskills models.Add_People_Skills

	peopleskills.Peopleskillsdesc = body.Peopleskillsdesc
	peopleskills.IsActive = body.IsActive

	// insert new records
	if result := h.DB.Create(&peopleskills); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Added New Data!", "data": &peopleskills})
}

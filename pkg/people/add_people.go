package people

import (
	"strconv"
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AddPeopleRequestBody struct {
	Cognizantid    int    `validate:"required" gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname       string `gorm:"column:lastname" json:"last_name"`
	Firstname      string `gorm:"column:firstname" json:"first_name"`
	Middlename     string `gorm:"column:middlename" json:"middle_name"`
	Fullname       string `validate:"required" gorm:"column:fullname" json:"full_name"`
	Csvemail       string `validate:"required,email" gorm:"column:csvemail" json:"csv_email"`
	Hireddate      string `validate:"required" gorm:"column:hireddate" json:"hired_date"`
	Communityid    int    `validate:"required" gorm:"column:communityid" json:"community_id"`
	Workstateid    int    `validate:"required" gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid     int    `gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid      *int   `gorm:"column:projectid" json:"project_id"`
	Isactive       bool   `gorm:"column:isactive" json:"is_active"`
	Isprobationary bool   `gorm:"column:isprobationary" json:"is_probationary"`
	Skills         string `json:"skills"`
	Details				 string `json:"details"`
}

func (h handler) AddPeople(c *fiber.Ctx) error {
	body := AddPeopleRequestBody{
		Cognizantid:    0,
		Lastname:       "",
		Firstname:      "",
		Middlename:     "",
		Fullname:       "",
		Csvemail:       "",
		Hireddate:      "",
		Communityid:    0,
		Workstateid:    0,
		Joblevelid:     0,
		Projectid:      nil,
		Isactive:       true,
		Isprobationary: false,
	}

	// parse body, attach to AddPeopleRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var people models.Add_People

	people.Cognizantid = body.Cognizantid
	people.Lastname = body.Lastname
	people.Firstname = body.Firstname
	people.Middlename = body.Middlename
	people.Fullname = body.Fullname
	people.Hireddate = body.Hireddate
	people.Csvemail = body.Csvemail
	people.Workstateid = body.Workstateid
	people.Communityid = body.Communityid
	people.Joblevelid = body.Joblevelid
	people.Projectid = body.Projectid
	people.Isactive = body.Isactive
	people.Isprobationary = body.Isprobationary

	// create transaction for insert
	transactionErr := h.DB.Transaction(func(tx *gorm.DB) error {
		//insert people
		if createPeopleErr := tx.Create(&people).Error; createPeopleErr != nil {
			return createPeopleErr
		}

		skillSet := strings.Split(body.Skills, ",")
		
		if (len(skillSet) == 1 && skillSet[0] != "") || len(skillSet) > 1 {
			var batchSkills []*models.InsertPrimarySkill
			for _, skillId := range skillSet {
				parsedSkill, _ := strconv.Atoi(skillId)
				var skillItem models.InsertPrimarySkill
				skillItem.PeopleId = people.PeopleId
				skillItem.PeopleSkill = parsedSkill
				skillItem.IsActive = true
				batchSkills = append(batchSkills, &skillItem)
			}

			//insert skills
			if len(batchSkills) > 0 {
				if insertSkillErr := tx.Create(&batchSkills).Error; insertSkillErr != nil {
					// return any error will rollback
					return insertSkillErr
				}
			}
		}

		details := strings.Split(body.Details, ",")

		if (len(details) == 1 && details[0] != "") || len(details) > 1 {
			var batchDetails []*models.InsertPeopleDetail
			for _, detailId := range details {
				parsedDetailId, _ := strconv.Atoi(detailId)
				var detailItem models.InsertPeopleDetail
				detailItem.PeopleId = people.PeopleId
				detailItem.PeopleDetailsDescId = parsedDetailId
				detailItem.ActiveFlag = true
				batchDetails = append(batchDetails, &detailItem)
			}

			if len(batchDetails) > 0 {
				if insertDetailsErr := tx.Create(&batchDetails).Error; insertDetailsErr != nil {
					return insertDetailsErr
				}
			}
		}

		return nil
	})

	if transactionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": transactionErr.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success! Added Data!", "data": &people})
}

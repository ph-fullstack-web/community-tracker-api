package people

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdatePeopleRequestBody struct {
	//Peopleid    int    `gorm:"column:peopleid" json:"people_id"`
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

func (h handler) UpdatePeople(c *fiber.Ctx) error {
	id := c.Params("peopleid")
	body := UpdatePeopleRequestBody{
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
		Skills:         "",
		Details:        "",
	}

	trim_id := strings.TrimLeft(id, "peopleid=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		fmt.Println(validateErr)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var people models.Update_People

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

	if result := h.DB.First(&people, "peopleid = ?", trim_id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

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

		mp := make(map[string]interface{})
		mp["cognizantid"] = body.Cognizantid
		mp["lastname"] = body.Lastname
		mp["firstname"] = body.Firstname
		mp["middlename"] = body.Middlename
		mp["fullname"] = body.Fullname
		mp["csvemail"] = body.Csvemail
		mp["hireddate"] = body.Hireddate
		mp["communityid"] = body.Communityid
		mp["workstateid"] = body.Workstateid
		mp["joblevelid"] = body.Joblevelid
		mp["projectid"] = body.Projectid
		mp["isactive"] = body.Isactive
		mp["isprobationary"] = body.Isprobationary

		skillSet := strings.Split(body.Skills, ",")

		var hasSkills = false
		var batchSkills []*models.InsertPrimarySkill
		parsedMemberId, _ := strconv.Atoi(trim_id)
		if body.Skills != "" {
			hasSkills = true
			for _, skillId := range skillSet {
				parsedSkill, _ := strconv.Atoi(skillId)
				var skillItem models.InsertPrimarySkill
				skillItem.PeopleId = parsedMemberId
				skillItem.PeopleSkill = parsedSkill
				skillItem.IsActive = true
				batchSkills = append(batchSkills, &skillItem)
			}
		}

		fmt.Print(batchSkills)
		fmt.Print(body.Skills)

		details := strings.Split(body.Details, ",")

		var hasDetails = false
		var batchDetails []*models.InsertPeopleDetail
		if body.Details != "" {
			hasDetails = true
			for _, detailId := range details {
				parsedDetailId, _ := strconv.Atoi(detailId)
				var detailItem models.InsertPeopleDetail
				detailItem.PeopleId = parsedMemberId
				detailItem.PeopleDetailsDescId = parsedDetailId
				detailItem.ActiveFlag = true
				batchDetails = append(batchDetails, &detailItem)
			}
		}


		transactionErr := h.DB.Transaction(func(tx *gorm.DB) error {
			//update people
			if updatePeopleErr := tx.Model(people).Where("peopleid = ?", trim_id).Updates(mp).Save(&people).Error; updatePeopleErr != nil {
				return updatePeopleErr
			}

			//delete all skills from member
			if delSkillErr := tx.Delete(&models.PeoplePrimarySkills{}, "peopleid = ?", parsedMemberId).Error; delSkillErr != nil {
				return delSkillErr
			}
			//insert all new skills
			fmt.Print(hasSkills)
			if hasSkills {
				if insertSkillErr := tx.Create(&batchSkills).Error; insertSkillErr != nil {
					// return any error will rollback
					return insertSkillErr
				}
			}

			if delDetailsErr := tx.Delete(&models.PeopleDetails{}, "peopleid = ?", parsedMemberId).Error; delDetailsErr != nil {
				return delDetailsErr
			}

			if (hasDetails) {
				if insertDetailsErr := tx.Create(&batchDetails).Error; insertDetailsErr != nil {
					return insertDetailsErr
				}
			}
			return nil
		})
		if transactionErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": transactionErr.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &people})
	}

}

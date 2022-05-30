package people

// import (
// 	"fmt"

// 	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
// 	"github.com/gofiber/fiber/v2"
// )

// type UpdatePeopleRequestBody struct {
// 	//Peopleid    int    `gorm:"column:peopleid" json:"people_id"`
// 	Cognizantid int    `gorm:"column:cognizantid" json:"cognizantid_id"`
// 	Lastname    string `gorm:"column:lastname" json:"last_name"`
// 	Firstname   string `gorm:"column:firstname" json:"first_name"`
// 	Middlename  string `gorm:"column:middlename" json:"middle_name"`
// 	Fullname    string `gorm:"column:fullname" json:"full_name"`
// 	Csvemail    string `gorm:"column:csvemail" json:"csv_email"`
// 	Hireddate   string `gorm:"column:hireddate" json:"hired_date"`
// 	Workstateid int    `gorm:"column:workstateid" json:"workstate_id"`
// 	Joblevelid  int    `gorm:"column:joblevelid" json:"joblevel_id"`
// 	Projectid   int    `gorm:"column:projectid" json:"project_id"`
// 	Isactive    bool   `gorm:"column:isactive" json:"is_active"`
// }

// func (h handler) UpdatePeople(c *fiber.Ctx) error {
// 	id := c.Params("peopleid")
// 	body := UpdatePeopleRequestBody{
// 		Cognizantid: 0,
// 		Lastname:    "",
// 		Firstname:   "",
// 		Middlename:  "",
// 		Fullname:    "",
// 		Csvemail:    "",
// 		Hireddate:   "",
// 		Workstateid: 0,
// 		Joblevelid:  0,
// 		Projectid:   0,
// 		Isactive:    false,
// 	}

// 	// parse body, attach to UpdateCityRequestBody struct
// 	if err := c.BodyParser(&body); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	var people models.UpdatePeople

// 	people.Cognizantid = body.Cognizantid
// 	people.Lastname = body.Lastname
// 	people.Firstname = body.Firstname
// 	people.Middlename = body.Middlename
// 	people.Fullname = body.Fullname
// 	people.Hireddate = body.Hireddate
// 	people.Csvemail = body.Csvemail
// 	people.Workstateid = body.Workstateid
// 	people.Joblevelid = body.Joblevelid
// 	people.Projectid = body.Projectid
// 	people.Isactive = body.Isactive

// 	if result := h.DB.First(&people, id); result.Error != nil {

// 		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

// 	} else {

// 		people.Cognizantid = body.Cognizantid
// 		people.Lastname = body.Lastname
// 		people.Firstname = body.Firstname
// 		people.Middlename = body.Middlename
// 		people.Fullname = body.Fullname
// 		people.Hireddate = body.Hireddate
// 		people.Csvemail = body.Csvemail
// 		people.Workstateid = body.Workstateid
// 		people.Joblevelid = body.Joblevelid
// 		people.Projectid = body.Projectid
// 		people.Isactive = body.Isactive

// 		h.DB.Save(&people)
// 		fmt.Println(result)
// 		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &people})
// 	}

// }

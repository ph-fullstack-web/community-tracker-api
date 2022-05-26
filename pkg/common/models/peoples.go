package models

type Peoples struct {
	Peopleid    int    `gorm:"column:peopleid" json:"people_id"`
	Cognizantid int    `gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname    string `gorm:"column:lastname" json:"last_name"`
	Firstname   string `gorm:"column:firstname" json:"first_name"`
	Middlename  string `gorm:"column:middlename" json:"middle_name"`
	Fullname    string `gorm:"column:fullname" json:"full_name"`
	Csvemail    string `gorm:"column:csvemail" json:"csv_email"`
	Hireddate   string `gorm:"column:hireddate" json:"hired_date"`
	Workstateid int    `gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid  int    `gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid   int    `gorm:"column:projectid" json:"project_id"`
	Isactive    bool   `gorm:"column:isactive" json:"is_active"`
}

func (Peoples) TableName() string {
	return "peoples"
}
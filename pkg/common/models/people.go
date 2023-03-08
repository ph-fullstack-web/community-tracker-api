package models

type People struct {
	Peopleid       int    		`gorm:"column:peopleid" json:"people_id"`
	Cognizantid    int    		`gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname       string 		`gorm:"column:lastname" json:"last_name"`
	Firstname      string 		`gorm:"column:firstname" json:"first_name"`
	Middlename     string 		`gorm:"column:middlename" json:"middle_name"`
	Fullname       string 		`gorm:"column:fullname" json:"full_name"`
	Csvemail       string 		`gorm:"column:csvemail" json:"csv_email"`
	Hireddate      string 		`gorm:"column:hireddate" json:"hired_date"`
	Communityid    int    		`gorm:"column:communityid" json:"community_id"`
	Workstateid    int    		`gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid     int    		`gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid      int    		`gorm:"column:projectid" json:"project_id"`
	Isactive       bool   		`gorm:"column:isactive" json:"is_active"`
	Isprobationary bool   		`gorm:"column:isprobationary" json:"is_probationary"`
	Community			 Community	`gorm:"foreignKey:Communityid;references:CommunityID" json:"community"`
}

func (People) TableName() string {
	return "people"
}

type Add_People struct {
	PeopleId       int    `gorm:"primaryKey;column:peopleid" json:"people_id"`
	Cognizantid    int    `gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname       string `gorm:"column:lastname" json:"last_name"`
	Firstname      string `gorm:"column:firstname" json:"first_name"`
	Middlename     string `gorm:"column:middlename" json:"middle_name"`
	Fullname       string `gorm:"column:fullname" json:"full_name"`
	Csvemail       string `gorm:"column:csvemail" json:"csv_email"`
	Hireddate      string `gorm:"column:hireddate" json:"hired_date"`
	Communityid    int    `gorm:"column:communityid" json:"community_id"`
	Workstateid    int    `gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid     int    `gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid      *int    `gorm:"column:projectid" json:"project_id"`
	Isactive       bool   `gorm:"column:isactive" json:"is_active"`
	Isprobationary bool   `gorm:"column:isprobationary" json:"is_probationary"`
}

func (Add_People) TableName() string {
	return "people"
}

type Update_People struct {
	Cognizantid    int    `gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname       string `gorm:"column:lastname" json:"last_name"`
	Firstname      string `gorm:"column:firstname" json:"first_name"`
	Middlename     string `gorm:"column:middlename" json:"middle_name"`
	Fullname       string `gorm:"column:fullname" json:"full_name"`
	Csvemail       string `gorm:"column:csvemail" json:"csv_email"`
	Hireddate      string `gorm:"column:hireddate" json:"hired_date"`
	Communityid    int    `gorm:"column:communityid" json:"community_id"`
	Workstateid    int    `gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid     int    `gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid      *int    `gorm:"column:projectid" json:"project_id"`
	Isactive       bool   `gorm:"column:isactive" json:"is_active"`
	Isprobationary bool   `gorm:"column:isprobationary" json:"is_probationary"`
}

func (Update_People) TableName() string {
	return "people"
}

type Manager struct {
	Cognizantid int    `gorm:"column:cognizantid" json:"cognizantid_id"`
	Fullname    string `gorm:"column:fullname" json:"full_name"`
	Csvemail    string `gorm:"column:csvemail" json:"csv_email"`
}

type PeopleWithSkills struct {
	Peopleid       int        				 	`json:"people_id"`
	Cognizantid    int        				 	`json:"cognizantid_id"`
	Fullname       string     				 	`json:"full_name"`
	Csvemail       string     				 	`json:"csv_email"`
	Hireddate      string     				 	`json:"hired_date"`
	Communityid    int        				 	`json:"community_id"`
	Workstateid    int        				 	`json:"workstate_id"`
	Joblevelid     int        				 	`json:"joblevel_id"`
	Projectid      int        				 	`json:"project_id"`
	Isactive       bool       				 	`json:"is_active"`
	Isprobationary bool       				 	`json:"is_probationary"`
	Skill          []SkillSet 				 	`json:"skills"`
	Details				 []PeopleDetailsDesc 	`json:"details"`
}

type PeopleUnderCommunitySearch struct {
	Peopleid int `gorm:"column:peopleid" json:"people_id"`
	Fullname string `gorm:"column:fullname" json:"full_name"`
	// Csvemail string `gorm:"column:csvemail" json:"csv_email"`
	// Hireddate string `gorm:"column:hireddate" json:"hired_date"`
	Communityid int `gorm:"column:communityid" json:"community_id"`
	// Community Community `gorm:"foreignKey:Communityid;references:CommunityID" json:"community"`
	// Isactive bool `gorm:"column:isactive" json:"is_active"`
}

func (PeopleUnderCommunitySearch) TableName() string {
	return "people"
}
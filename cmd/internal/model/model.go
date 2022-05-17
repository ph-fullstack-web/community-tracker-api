package model

type Community struct {
	Communityid          int    `json:"communityid" gorm:"primaryKey;autoIncrement:true"`
	Communityname        string `json:"communityname" gorm:"varchar(50);not null"`
	Communitymgrpeopleid int    `json:"communitymgrpeopleid" gorm:"not null"`
	Communityimg         string `json:"communityimg"`
	Communitycolor       string `json:"communitycolor" gorm:"varchar(100)"`
	Isactive             bool   `json:"isactive" gorm:"not null"`
}

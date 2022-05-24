package models

type Community struct {
	Communityid          int    `json:"communityid" gorm:"primaryKey;autoIncrement:true"`
	Communityname        string `json:"communityname"`
	Communitymgrpeopleid int    `json:"communitymanagerpeopleid"`
	Isactive             bool   `json:"isactive"`
}

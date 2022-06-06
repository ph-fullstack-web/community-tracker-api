package models

type Peopleskills struct {
	Peopleskillsid   int    `gorm:"column:peopleskillsid" json:"peopleskills_id"`
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"peopleskills_desc"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (Peopleskills) TableName() string {
	return "peopleskills"
}

type Deleteskills struct {
	IsActive bool `gorm:"column:isactive" json:"is_active"`
}

func (Deleteskills) TableName() string {
	return "peopleskills"
}

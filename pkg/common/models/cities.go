package models

type Cities struct {
	Name    string `json:"name"`
	City_ID int    `json:"city_id"`
}

type AddCity struct {
	Name string `json:"name"`
}

func (AddCity) TableName() string {
	return "cities"
}

type UpdateCity struct {
	Name string `json:"name"`
}

func (UpdateCity) TableName() string {
	return "cities"
}

type DeleteCity struct {
	Name    string `json:"name"`
	City_ID int    `json:"city_id"`
}

func (DeleteCity) TableName() string {
	return "cities"
}

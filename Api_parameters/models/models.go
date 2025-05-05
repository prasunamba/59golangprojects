package models

type Movie struct {
	Id     int    `gorm:"primarykey"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Budget string `json:"budget"`
}

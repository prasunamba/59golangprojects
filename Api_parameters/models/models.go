package models

type Movie struct {
	Name   string `gorm:"primarykey"  json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Budget string `json:"budget"`
}

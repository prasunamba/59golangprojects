package database

import (
	"example/Api_parameters/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database_connection() {
	dsn := "prasuna:1313@tcp(127.0.0.1:3306)/myappDB"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil || DB == nil {
		log.Fatal("failed to login to DB", err, DB)
	}
	err = DB.AutoMigrate(&models.Movie{})
	if err != nil {
		log.Fatal("failed to create table")
	}
	fmt.Println("successfully create table")
}

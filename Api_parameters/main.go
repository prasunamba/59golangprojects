package main

import (
	"example/Api_parameters/database"
	"example/Api_parameters/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Set log output to the file
	log.SetOutput(logFile)
	log.Println("Application started")
	log.Println("This is a log entry to the file")
	router := gin.Default()
	router.POST("/insertrecord", handlers.Insertrecord)
	router.GET("/listmovies", handlers.Listmovies)
	router.Run()
}
func init() {
	database.Database_connection()
}

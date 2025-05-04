package handlers

import (
	"example/Api_parameters/database"
	"example/Api_parameters/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Listmovies(c *gin.Context) {
	var records []models.Movie
	result := database.DB.Find(&records)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	log.Println("listing movies")
	c.JSON(http.StatusOK, records)
}
func Insertrecord(c *gin.Context) {
	var record models.Movie
	if err := c.ShouldBindJSON(&record); err != nil {
		log.Fatal("failed to get the record   ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Insert into database
	if err := database.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("record inserted successfully")
}

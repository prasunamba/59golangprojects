package handlers

import (
	"example/Api_parameters/database"
	"example/Api_parameters/models"
	"log"
	"net/http"
	"strconv"

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
	log.Println("record inserted successfully")
}
func Deleterecord(c *gin.Context) {
	stryear := c.Param("year")
	year, _ := strconv.Atoi(stryear)
	genre := c.Param("genre")
	var movie models.Movie
	log.Println("year & genre", stryear, genre)
	result := database.DB.Where("year=? AND genre = ?", year, genre).First(&movie)
	if result.Error != nil {
		log.Fatal("record not found")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}
	if err := database.DB.Delete(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete the reccord",
		})
		return
	}
	log.Println("successfully deleted the record")
}
func Updaterecord(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "Missing 'name' query param"})
		return
	}

	updates := map[string]interface{}{}
	if yearStr := c.Query("year"); yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err == nil {
			updates["year"] = year
		}
	}
	if genre := c.Query("genre"); genre != "" {
		updates["genre"] = genre
	}
	if len(updates) == 0 {
		c.JSON(400, gin.H{"error": "No fields provided for update"})
		return
	}
	result := database.DB.Model(&models.Movie{}).Where("name = ?", name).Updates(updates)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Movie not found or nothing updated"})
		return
	}
	log.Println("record updated successfully")
	c.JSON(200, gin.H{"message": "Movie updated", "updates": updates})
}

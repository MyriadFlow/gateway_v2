package controllers

import (
	"net/http"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"

	"github.com/gin-gonic/gin"
)

func CreatePhygital(c *gin.Context) {

	var phygital models.Phygital
	if err := c.ShouldBindJSON(&phygital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&phygital).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, phygital)
}

func GetPhygital(c *gin.Context) {
	id := c.Param("id")
	var phygital models.Phygital
	if err := db.DB.First(&phygital, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phygital not found"})
		return
	}

	c.JSON(http.StatusOK, phygital)
}

// get all phygital api
func GetAllPhygital(c *gin.Context) {
	var phygitals []models.Phygital
	if err := db.DB.Find(&phygitals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, phygitals)
}

func UpdatePhygital(c *gin.Context) {
	id := c.Param("id")
	var phygital models.Phygital
	if err := db.DB.First(&phygital, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phygital not found"})
		return
	}

	if err := c.ShouldBindJSON(&phygital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&phygital).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, phygital)
}

func DeletePhygital(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Phygital{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phygital deleted successfully"})
}

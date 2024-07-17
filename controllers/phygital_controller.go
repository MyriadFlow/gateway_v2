package controllers

import (
	"encoding/json"
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
	var oldPhygital models.Phygital
	if err := db.DB.First(&oldPhygital, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phygital not found"})
		return
	}

	c.JSON(http.StatusOK, oldPhygital)
}

// get all phygital api
func GetAllPhygital(c *gin.Context) {
	var phygitals []models.Phygital
	chainTypeId := c.Param("chaintype_id")
	if err := db.DB.Find(&phygitals).Where("chaintype_id = ? " , chainTypeId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, phygitals)
}

func UpdatePhygital(c *gin.Context) {
	id := c.Param("id")
	var Phygital models.Phygital

	if err := db.DB.First(&Phygital, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phygital not found"})
		return
	}

	var input models.Phygital
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err1 error
	// Ensure the input Category is properly converted to JSON
	input.Category, err1 = json.Marshal(input.Category)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	// Update fields
	db.DB.Model(&Phygital).Updates(input)

	c.JSON(http.StatusOK, Phygital)

}

func DeletePhygital(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Phygital{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phygital deleted successfully"})
}

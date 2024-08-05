package controllers

import (
	"app.myriadflow.com/db"

	"net/http"

	"app.myriadflow.com/models"

	"github.com/gin-gonic/gin"
)

func CreateCopies(c *gin.Context) {
	var copies models.Copies
	if err := c.ShouldBindJSON(&copies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&copies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, copies)
}

func GetCopiesById(c *gin.Context) {
	id := c.Param("id")
	var copies models.Copies
	if err := db.DB.First(&copies, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Copies not found"})
		return
	}

	c.JSON(http.StatusOK, copies)
}

func GetCopiesByPhygitalID(c *gin.Context) {
	phygitalID := c.Param("phygital_id")
	var copies models.Copies
	if err := db.DB.First(&copies, "phygital_id = ?", phygitalID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Copies not found"})
		return
	}

	c.JSON(http.StatusOK, copies)
}

func GetAllCopiesByChainType(c *gin.Context) {
	chaintypeId := c.Param("chaintype_id")
	var copies []models.Copies
	if err := db.DB.Where("chaintype_id = ?" , chaintypeId).Find(&copies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, copies)
}

func GetOwnerByPhygitalAndCopyNumber(c *gin.Context) {
	phygitalID := c.Param("phygital_id")
	copyNumber := c.Param("copy_number")

	var copies models.Copies
	if err := db.DB.First(&copies, "phygital_id = ? AND copy_number = ?", phygitalID, copyNumber).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"wallet_address": copies.WalletAddress})
}


func UpdateCopies(c *gin.Context) {
	id := c.Param("id")
	var copies models.Copies
	if err := db.DB.First(&copies, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Copies not found"})
		return
	}

	if err := c.ShouldBindJSON(&copies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&copies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, copies)
}

func DeleteCopies(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Copies{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Copies deleted successfully"})
}
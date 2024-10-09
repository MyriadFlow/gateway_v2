package controllers

import (
	"app.myriadflow.com/db"
	"net/http"
	"app.myriadflow.com/models"
	"github.com/gin-gonic/gin"
)

func CreateElevate(c *gin.Context) {
	var elevate models.Elevate
	if err := c.ShouldBindJSON(&elevate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&elevate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, elevate)
}

func GetElevate(c *gin.Context) {
	id := c.Param("id")
	var elevate models.Elevate
	if err := db.DB.First(&elevate, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Elevate not found"})
		return
	}

	c.JSON(http.StatusOK, elevate)
}

func GetAllElevate(c *gin.Context) {
	var elevate []models.Elevate
	if err := db.DB.Find(&elevate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, elevate)
}

func GetAllElevateByChainType(c *gin.Context) {
	chaintypeId := c.Param("chaintype_id")
	var elevate []models.Elevate
	if err := db.DB.Where("chaintype_id = ? " , chaintypeId).Find(&elevate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, elevate)
}

func GetElevateByWalletAddress(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	var elevate models.Elevate
	if err := db.DB.Where("wallet_address = ?", walletAddress).First(&elevate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, elevate)
}

func UpdateElevate(c *gin.Context) {
	id := c.Param("id")
	var elevate models.Elevate
	if err := db.DB.First(&elevate, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Elevate not found"})
		return
	}

	if err := c.ShouldBindJSON(&elevate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&elevate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, elevate)
}

func DeleteElevate(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Elevate{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Elevate not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Elevate deleted"})
}
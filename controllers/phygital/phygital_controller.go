package phygital_controllers

import (
	"encoding/json"
	"net/http"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"

	"github.com/gin-gonic/gin"
)

func CreatePhygital(c *gin.Context) {

	var phygital models.Phygital
	var reqPhygital RequestPhygital
	if err := c.ShouldBindJSON(&reqPhygital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(reqPhygital.Images) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one image is required"})
		return
	}

	if reqPhygital.SizeOption != 0 && reqPhygital.SizeOption != 1 && reqPhygital.SizeOption != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size option"})
		return
	}

	if reqPhygital.SizeOption == 1 {

		var sizeDetails map[string]interface{}
		if err := json.Unmarshal(phygital.SizeDetails, &sizeDetails); err != nil || len(sizeDetails) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Size details are required for 'one_size_with_measurements' option"})
			return
		}
	} else if reqPhygital.SizeOption == 2 {
		if len(phygital.SizeDetails) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Size details are required for 'multiple_sizes' option"})
			return
		}

		// for _, sizeDetail := range phygital.SizeDetails {
		// 	if sizeDetail.Size == "" {
		// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Size is required for each entry in size_details"})
		// 		return
		// 	}
		// 	if sizeDetail.Quantity <= 0 {
		// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than 0"})
		// 		return
		// 	}
		// }
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

func GetAllPhygitalByRegion(c *gin.Context) {
	var phygitals []models.Phygital
	if err := db.DB.Where("elevate_region = ?", "Africa").Find(&phygitals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, phygitals)
}

func GetPhygitalByWalletAddress(c *gin.Context) {
	deployer_address := c.Param("deployer_address")
	var oldPhygital models.Phygital
	if err := db.DB.First(&oldPhygital, "deployer_address = ?", deployer_address).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phygital not found"})
		return
	}

	c.JSON(http.StatusOK, oldPhygital)
}

// get all phygital api
func GetAllPhygitalByChainType(c *gin.Context) {
	chaintypeId := c.Param("chaintype_id")
	var phygitals []models.Phygital
	if err := db.DB.Where("chaintype_id = ? ", chaintypeId).Find(&phygitals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, phygitals)
}

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

	// Ensure the input Tags is properly converted to JSON
	input.Tags, err1 = json.Marshal(input.Tags)
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

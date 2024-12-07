package phygital_controllers

import (
	"encoding/json"
	"net/http"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"gorm.io/datatypes"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size option. Must be 0, 1, or 2."})
		return
	}

	if reqPhygital.SizeOption == 0 {
		phygital.SizeDetails = datatypes.JSON([]byte("{}")) // Store empty JSON for `SizeOption == 0`
	} else {
		if len(reqPhygital.SizeDetails) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "SizeDetails is required for SizeOption 1 or 2."})
			return
		}
		phygital.SizeDetails = reqPhygital.SizeDetails // Store incoming JSON object
	}

	
	phygital.Name = reqPhygital.Name
	phygital.BrandName = reqPhygital.BrandName
	phygital.Category = reqPhygital.Category
	phygital.Tags = reqPhygital.Tags
	phygital.Description = reqPhygital.Description
	phygital.Price = &reqPhygital.Price
	phygital.Quantity = reqPhygital.Quantity
	phygital.Royality = reqPhygital.Royality
	phygital.Images = reqPhygital.Images
	phygital.ProductInfo = reqPhygital.ProductInfo
	phygital.ProductUrl = reqPhygital.ProductUrl
	phygital.Color = reqPhygital.Color
	phygital.SizeOption = reqPhygital.SizeOption
	phygital.Weight = reqPhygital.Weight
	phygital.Material = reqPhygital.Material
	phygital.Usage = reqPhygital.Usage
	phygital.Quality = reqPhygital.Quality
	phygital.Manufacturer = reqPhygital.Manufacturer
	phygital.OriginCountry = reqPhygital.OriginCountry
	phygital.MetadataURI = reqPhygital.MetadataURI
	phygital.DeployerAddress = reqPhygital.DeployerAddress
	phygital.ContractAddress = reqPhygital.ContractAddress
	phygital.GraphURL = reqPhygital.GraphURL
	phygital.ElevateRegion = reqPhygital.ElevateRegion
	phygital.CollectionID = reqPhygital.CollectionID
	phygital.ChaintypeID = reqPhygital.ChaintypeID

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

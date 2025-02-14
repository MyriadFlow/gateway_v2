package controllers

import (
	"net/http"
	"strings"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"github.com/gin-gonic/gin"
)

func CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand.SlugName = func(name string) string {
		return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	}(brand.Name)

	if err := db.DB.Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func GetBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := db.DB.First(&brand, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, brand)
}
func GetBrandsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	var brands []models.Brand

	if err := db.DB.Where("user_id = ?", userID).Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(brands) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No brands found for the user"})
		return
	}

	c.JSON(http.StatusOK, brands)
}

func GetBrandByName(c *gin.Context) {
	name := c.Param("name")
	var brand models.Brand
	if err := db.DB.First(&brand, "slug_name = ?", name).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, brand)
}

// get all brands api
func GetAllBrandsByChainType(c *gin.Context) {
	chaintypeId := c.Param("chaintype_id")
	var brands []models.Brand
	if err := db.DB.Where("chaintype_id = ?", chaintypeId).Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func GetAllBrands(c *gin.Context) {
	var brands []models.Brand
	if err := db.DB.Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func GetBrandsByManager(c *gin.Context) {
	managerID := c.Param("manager_id")
	var brands []models.Brand
	if err := db.DB.Where("manager_id = ?", managerID).Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(brands) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No brands found for the specified manager_id"})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func GetAllBrandsByRegion(c *gin.Context) {
	var brands []models.Brand
	if err := db.DB.Where("elevate_region = ?", "Africa").Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := db.DB.First(&brand, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if condition := func(name string) bool {
		return name != ""
	}; condition(brand.Name) {
		brand.SlugName = func(name string) string {
			return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
		}(brand.Name)

	}

	if err := db.DB.Save(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Brand{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}

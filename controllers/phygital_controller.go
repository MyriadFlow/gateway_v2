package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var oldPhygital Phygital
	if err := db.DB.First(&oldPhygital, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phygital not found"})
		return
	}

	var categoryMap map[string]interface{}
	if len(oldPhygital.Category) > 0 {
		err := json.Unmarshal(oldPhygital.Category, &categoryMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("error unmarshalling category for Phygital ID %s: %v", oldPhygital.ID, err)})
		}
	}
	newPhygitals := models.Phygital{
		ID:              oldPhygital.ID,
		Name:            oldPhygital.Name,
		BrandName:       oldPhygital.BrandName,
		Category:        categoryMap,
		Description:     oldPhygital.Description,
		Price:           oldPhygital.Price,
		Quantity:        oldPhygital.Quantity,
		Royality:        oldPhygital.Royality,
		Image:           oldPhygital.Image,
		ProductInfo:     oldPhygital.ProductInfo,
		Color:           oldPhygital.Color,
		Size:            oldPhygital.Size,
		Weight:          oldPhygital.Weight,
		Material:        oldPhygital.Material,
		Usage:           oldPhygital.Usage,
		Quality:         oldPhygital.Quality,
		Manufacturer:    oldPhygital.Manufacturer,
		OriginCountry:   oldPhygital.OriginCountry,
		MetadataURI:     oldPhygital.MetadataURI,
		DeployerAddress: oldPhygital.DeployerAddress,
		ContractAddress: oldPhygital.ContractAddress,
		GraphURL:        oldPhygital.GraphURL,
		CollectionID:    oldPhygital.CollectionID,
		CreatedAt:       oldPhygital.CreatedAt,
		UpdatedAt:       oldPhygital.UpdatedAt,
	}

	c.JSON(http.StatusOK, newPhygitals)
}

// get all phygital api
func GetAllPhygital(c *gin.Context) {
	var phygitals []Phygital
	if err := db.DB.Find(&phygitals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, errr := convertPhygitals(phygitals)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
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

func convertPhygitals(oldPhygitals []Phygital) ([]models.Phygital, error) {
	newPhygitals := make([]models.Phygital, len(oldPhygitals))

	for i, oldPhygital := range oldPhygitals {
		var categoryMap map[string]interface{}
		if len(oldPhygital.Category) > 0 {
			err := json.Unmarshal(oldPhygital.Category, &categoryMap)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling category for Phygital ID %s: %v", oldPhygital.ID, err)
			}
		}

		newPhygitals[i] = models.Phygital{
			ID:              oldPhygital.ID,
			Name:            oldPhygital.Name,
			BrandName:       oldPhygital.BrandName,
			Category:        categoryMap,
			Description:     oldPhygital.Description,
			Price:           oldPhygital.Price,
			Quantity:        oldPhygital.Quantity,
			Royality:        oldPhygital.Royality,
			Image:           oldPhygital.Image,
			ProductInfo:     oldPhygital.ProductInfo,
			Color:           oldPhygital.Color,
			Size:            oldPhygital.Size,
			Weight:          oldPhygital.Weight,
			Material:        oldPhygital.Material,
			Usage:           oldPhygital.Usage,
			Quality:         oldPhygital.Quality,
			Manufacturer:    oldPhygital.Manufacturer,
			OriginCountry:   oldPhygital.OriginCountry,
			MetadataURI:     oldPhygital.MetadataURI,
			DeployerAddress: oldPhygital.DeployerAddress,
			ContractAddress: oldPhygital.ContractAddress,
			GraphURL:        oldPhygital.GraphURL,
			CollectionID:    oldPhygital.CollectionID,
			CreatedAt:       oldPhygital.CreatedAt,
			UpdatedAt:       oldPhygital.UpdatedAt,
		}
	}

	return newPhygitals, nil
}

type Phygital struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name            string    `json:"name"`
	BrandName       string    `json:"brand_name"`
	Category        []byte    `gorm:"type:jsonb" json:"category"`
	Description     string    `json:"description"`
	Price           int       `json:"price"`
	Quantity        int       `json:"quantity"`
	Royality        int       `json:"royality"`
	Image           string    `json:"image"`
	ProductInfo     string    `json:"product_info"`
	Color           string    `json:"color"`
	Size            string    `json:"size"`
	Weight          int       `json:"weight"`
	Material        string    `json:"material"`
	Usage           string    `json:"usage"`
	Quality         string    `json:"quality"`
	Manufacturer    string    `json:"manufacturer"`
	OriginCountry   string    `json:"origin_country"`
	MetadataURI     string    `json:"metadata_uri"`
	DeployerAddress string    `json:"deployer_address"`
	ContractAddress string    `json:"contract_address"`
	GraphURL        string    `json:"graph_url"`
	CollectionID    uuid.UUID `json:"collection_id"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

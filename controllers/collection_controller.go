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

func CreateCollection(c *gin.Context) {
	var collection models.Collection
	if err := c.ShouldBindJSON(&collection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&collection).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, collection)
}

func GetCollection(c *gin.Context) {
	id := c.Param("id")
	var collection Collection
	if err := db.DB.First(&collection, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
		return
	}
	var categoryMap map[string]interface{}
	if len(collection.Category) > 0 {
		err := json.Unmarshal(collection.Category, &categoryMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Errorf("error unmarshalling category for collection ID %s: %v", collection.ID, err))
			// return nil, fmt.Errorf("error unmarshalling category for collection ID %s: %v", collection.ID, err)
		}
	}
	newCollections := models.Collection{
		ID:          collection.ID,
		Name:        collection.Name,
		Description: collection.Description,
		LogoImage:   collection.LogoImage,
		CoverImage:  collection.CoverImage,
		Category:    categoryMap,
		Tags:        collection.Tags,
		Status:      collection.Status,
		BrandID:     collection.BrandID,
		CreatedAt:   collection.CreatedAt,
		UpdatedAt:   collection.UpdatedAt,
	}

	c.JSON(http.StatusOK, newCollections)
}

type Collection struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LogoImage   string    `json:"logo_image"`
	CoverImage  string    `json:"cover_image"`
	Category    []byte    `gorm:"type:jsonb" json:"category"`
	Tags        string    `json:"tags"`
	Status      int       `json:"status"`
	BrandID     string    `json:"brand_id"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

// get all connection api
func GetAllCollections(c *gin.Context) {
	var collections []Collection
	if err := db.DB.Find(&collections).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Convert Category []byte to map[string]interface{}
	// for i := range collections {
	// 	var categoryMap map[string]interface{}
	// 	if err := json.Unmarshal(collections[i].Category, &categoryMap); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unmarshal category"})
	// 		return
	// 	}
	// 	collections[i].Category = categoryMap
	// }
	data, errors := convertCollections(collections)
	if errors != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors})
		return
	}

	c.JSON(http.StatusOK, data)
}

func convertCollections(oldCollections []Collection) ([]models.Collection, error) {
	newCollections := make([]models.Collection, len(oldCollections))

	for i, oldCollection := range oldCollections {
		var categoryMap map[string]interface{}
		if len(oldCollection.Category) > 0 {
			err := json.Unmarshal(oldCollection.Category, &categoryMap)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling category for collection ID %s: %v", oldCollection.ID, err)
			}
		}

		newCollections[i] = models.Collection{
			ID:          oldCollection.ID,
			Name:        oldCollection.Name,
			Description: oldCollection.Description,
			LogoImage:   oldCollection.LogoImage,
			CoverImage:  oldCollection.CoverImage,
			Category:    categoryMap,
			Tags:        oldCollection.Tags,
			Status:      oldCollection.Status,
			BrandID:     oldCollection.BrandID,
			CreatedAt:   oldCollection.CreatedAt,
			UpdatedAt:   oldCollection.UpdatedAt,
		}
	}

	return newCollections, nil
}

// func GetAllCollections(c *gin.Context) {
// 	var collections []Collection
// 	if err := db.DB.Raw("SELECT * FROM collections").Scan(&collections).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, collections)
// }

func UpdateCollection(c *gin.Context) {
	id := c.Param("id")
	var collection models.Collection
	if err := db.DB.First(&collection, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
		return
	}

	if err := c.ShouldBindJSON(&collection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&collection).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, collection)
}

func DeleteCollection(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Collection{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collection deleted successfully"})
}

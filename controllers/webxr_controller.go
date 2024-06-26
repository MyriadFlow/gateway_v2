package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"app.myriadflow.com/db"

	"app.myriadflow.com/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateWebXR(c *gin.Context) {
	var webxr models.WebXR
	if err := c.ShouldBindJSON(&webxr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&webxr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webxr)
}

func GetWebXR(c *gin.Context) {
	id := c.Param("id")
	var webxr WebXR
	if err := db.DB.First(&webxr, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var customizationsMap map[string]interface{}
	if len(webxr.Customizations) > 0 {
		err1 := json.Unmarshal(webxr.Customizations, &customizationsMap)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Errorf("error unmarshalling customizations for WebXR ID %s: %v", webxr.ID, err1.Error()),
			})
			return
		}
	}

	newWebXRs := models.WebXR{
		ID:                 webxr.ID,
		Image360:           webxr.Image360,
		Video360:           webxr.Video360,
		RewardsMetadataURI: webxr.RewardsMetadataURI,
		Customizations:     customizationsMap,
		FreeNFTImage:       webxr.FreeNFTImage,
		GoldReward:         webxr.GoldReward,
		SilverReward:       webxr.SilverReward,
		BronzeReward:       webxr.BronzeReward,
		PhygitalID:         webxr.PhygitalID,
		CreatedAt:          webxr.CreatedAt,
		UpdatedAt:          webxr.UpdatedAt,
	}

	c.JSON(http.StatusOK, newWebXRs)
}

type WebXR struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Image360           string    `json:"image360"`
	Video360           string    `json:"video360"`
	RewardsMetadataURI string    `json:"rewards_metadata_uri"`
	Customizations     []byte    `gorm:"type:jsonb" json:"customizations"`
	FreeNFTImage       string    `json:"free_nft_image"`
	GoldReward         string    `json:"gold_reward"`
	SilverReward       string    `json:"silver_reward"`
	BronzeReward       string    `json:"bronze_reward"`
	PhygitalID         string    `json:"phygital_id"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

// get all webxr
func GetAllWebXR(c *gin.Context) {
	var webxr []WebXR
	if err := db.DB.Find(&webxr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, errors := convertWebXRs(webxr)
	if errors != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func convertWebXRs(oldWebXRs []WebXR) ([]models.WebXR, error) {
	newWebXRs := make([]models.WebXR, len(oldWebXRs))

	for i, oldWebXR := range oldWebXRs {
		var customizationsMap map[string]interface{}
		if len(oldWebXR.Customizations) > 0 {
			err := json.Unmarshal(oldWebXR.Customizations, &customizationsMap)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling customizations for WebXR ID %s: %v", oldWebXR.ID, err)
			}
		}

		newWebXRs[i] = models.WebXR{
			ID:                 oldWebXR.ID,
			Image360:           oldWebXR.Image360,
			Video360:           oldWebXR.Video360,
			RewardsMetadataURI: oldWebXR.RewardsMetadataURI,
			Customizations:     customizationsMap,
			FreeNFTImage:       oldWebXR.FreeNFTImage,
			GoldReward:         oldWebXR.GoldReward,
			SilverReward:       oldWebXR.SilverReward,
			BronzeReward:       oldWebXR.BronzeReward,
			PhygitalID:         oldWebXR.PhygitalID,
			CreatedAt:          oldWebXR.CreatedAt,
			UpdatedAt:          oldWebXR.UpdatedAt,
		}
	}

	return newWebXRs, nil
}

func UpdateWebXR(c *gin.Context) {
	id := c.Param("id")
	var webxr models.WebXR
	if err := db.DB.First(&webxr, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "WebXR not found"})
		return
	}

	if err := c.ShouldBindJSON(&webxr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&webxr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webxr)
}

func DeleteWebXR(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.WebXR{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "WebXR deleted successfully"})
}

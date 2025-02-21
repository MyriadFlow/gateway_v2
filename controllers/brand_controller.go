package controllers

import (
	"net/http"
	"strings"
	"time"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	brand.ContactEmail = ""
	brand.ContactPhone = ""
	brand.ShippingAddress = ""

	c.JSON(http.StatusOK, brand)
}
func GetBrandsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	var brands []models.Brand

	if err := db.DB.Where("user_id = ?", userID).Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range brands {
		brands[i].ContactEmail = ""
		brands[i].ContactPhone = ""
		brands[i].ShippingAddress = ""
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
	var agents []models.Agent
	if err := db.DB.Where("agent_category_id = ?", brand.ID).Find(&agents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve agents"})
		return
	}
	response := brandResponse{
		ID:             brand.ID,
		Name:           brand.Name,
		SlugName:       brand.SlugName,
		AgentId:        brand.AgentId,
		AvatarId:       brand.AvatarId,
		Slogan:         brand.Slogan,
		Description:    brand.Description,
		LogoImage:      brand.LogoImage,
		CoverImage:     brand.CoverImage,
		Representative: brand.Representative,
		// ContactEmail:                brand.ContactEmail,
		// ContactPhone:                brand.ContactPhone,
		// ShippingAddress:             brand.ShippingAddress,
		Website:                     brand.Website,
		Twitter:                     brand.Twitter,
		Instagram:                   brand.Instagram,
		Facebook:                    brand.Facebook,
		Telegram:                    brand.Telegram,
		LinkedIn:                    brand.LinkedIn,
		Youtube:                     brand.Youtube,
		Discord:                     brand.Discord,
		Whatsapp:                    brand.Whatsapp,
		Google:                      brand.Google,
		Tiktok:                      brand.Tiktok,
		Snapchat:                    brand.Snapchat,
		Pinetrest:                   brand.Pinetrest,
		AdditionalLink:              brand.AdditionalLink,
		Link:                        brand.Link,
		AdditionalInfo:              brand.AdditionalInfo,
		Industry:                    brand.Industry,
		Tags:                        brand.Tags,
		Fees:                        brand.Fees,
		PayoutAddress:               brand.PayoutAddress,
		AccessMaster:                brand.AccessMaster,
		TradeHub:                    brand.TradeHub,
		Blockchain:                  brand.Blockchain,
		ChainID:                     brand.ChainID,
		ChaintypeID:                 brand.ChaintypeID,
		ManagerID:                   brand.ManagerID,
		ElevateRegion:               brand.ElevateRegion,
		WebXRExperienceWithAiAvatar: brand.WebXRExperienceWithAiAvatar,
		Image360:                    brand.Image360,
		Video360:                    brand.Video360,
		CreatedAt:                   brand.CreatedAt,
		UpdatedAt:                   brand.UpdatedAt,
		AgentDetails:                agents,
	}

	c.JSON(http.StatusOK, response)
}

type brandResponse struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name           string    `json:"name"`
	SlugName       string    `gorm:"unique" json:"slug_name"`
	AgentId        string    `json:"agent_id"`
	AvatarId       string    `json:"avatar_id"`
	Slogan         string    `json:"slogan"`
	Description    string    `json:"description"`
	LogoImage      string    `json:"logo_image"`
	CoverImage     string    `json:"cover_image"`
	Representative string    `json:"representative"`
	// ContactEmail                string    `json:"contact_email"`
	// ContactPhone                string    `json:"contact_phone"`
	// ShippingAddress             string    `json:"shipping_address"`
	Website                     string    `json:"website"`
	Twitter                     string    `json:"twitter"`
	Instagram                   string    `json:"instagram"`
	Facebook                    string    `json:"facebook"`
	Telegram                    string    `json:"telegram"`
	LinkedIn                    string    `json:"linkedin"`
	Youtube                     string    `json:"youtube"`
	Discord                     string    `json:"discord"`
	Whatsapp                    string    `json:"whatsapp"`
	Google                      string    `json:"google"`
	Tiktok                      string    `json:"tiktok"`
	Snapchat                    string    `json:"snapchat"`
	Pinetrest                   string    `json:"pinetrest"`
	AdditionalLink              string    `json:"additional_link"`
	Link                        string    `json:"link"`
	AdditionalInfo              string    `json:"additional_info"`
	Industry                    string    `json:"industry"`
	Tags                        string    `json:"tags"`
	Fees                        float64   `json:"fees" gorm:"type:decimal(20,10)"`
	PayoutAddress               string    `json:"payout_address"`
	AccessMaster                string    `json:"access_master"`
	TradeHub                    string    `json:"trade_hub"`
	Blockchain                  string    `json:"blockchain"`
	ChainID                     string    `json:"chain_id"`
	ChaintypeID                 uuid.UUID `gorm:"type:uuid" json:"chaintype_id"`
	ManagerID                   string    `json:"manager_id"` //user walletaddress
	ElevateRegion               string    `json:"elevate_region"`
	WebXRExperienceWithAiAvatar bool      `json:"webxr_experience_with_ai_avatar"`
	Image360                    string    `json:"image360"`
	Video360                    string    `json:"video360"`
	CreatedAt                   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt                   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	AgentDetails                []models.Agent
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
	for i := range brands {
		brands[i].ContactEmail = ""
		brands[i].ContactPhone = ""
		brands[i].ShippingAddress = ""
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
	for i := range brands {
		brands[i].ContactEmail = ""
		brands[i].ContactPhone = ""
		brands[i].ShippingAddress = ""
	}
	c.JSON(http.StatusOK, brands)
}

func GetAllBrandsByRegion(c *gin.Context) {
	var brands []models.Brand
	if err := db.DB.Where("elevate_region = ?", "Africa").Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range brands {
		brands[i].ContactEmail = ""
		brands[i].ContactPhone = ""
		brands[i].ShippingAddress = ""
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

	brand.ContactEmail = ""
	brand.ContactPhone = ""
	brand.ShippingAddress = ""

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

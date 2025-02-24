package controllers

import (
	"net/http"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProfile creates a new profile
func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// profile.Email = ""

	c.JSON(http.StatusOK, profile)
}

// GetProfile retrieves a profile by ID
func GetProfile(c *gin.Context) {
	id := c.Param("id")
	var profile models.Profile
	if err := db.DB.First(&profile, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// profile.Email = ""

	c.JSON(http.StatusOK, profile)
}

// GetAllProfiles retrieves all profiles
func GetAllProfiles(c *gin.Context) {
	var profiles []models.Profile
	if err := db.DB.Find(&profiles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// for i := range profiles {
	// 	profiles[i].Email = ""
	// }
	c.JSON(http.StatusOK, profiles)
}

// GetAllProfilesByChainType retrieves all profiles by ChaintypeID
func GetAllProfilesByChainType(c *gin.Context) {
	chaintypeId := c.Param("chaintype_id")
	var profiles []models.Profile
	if err := db.DB.Where("chaintype_id = ?", chaintypeId).Find(&profiles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// for i := range profiles {
	// 	profiles[i].Email = ""
	// }
	c.JSON(http.StatusOK, profiles)
}

// UpdateProfile updates a profile by walletAddress
func UpdateProfile(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	var profile models.Profile
	if err := db.DB.First(&profile, "wallet_address = ?", walletAddress).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// profile.Email = ""

	c.JSON(http.StatusOK, profile)
}

// DeleteProfile deletes a profile by ID
func DeleteProfile(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Profile{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}

func DeleteProfileByWalletAndEmail(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	email := c.Param("email")

	// Check if both walletAddress and email are provided
	if walletAddress == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet address and email are required"})
		return
	}

	// Attempt to delete the profile based on walletAddress and email
	if err := db.DB.Where("wallet_address = ? AND email = ?", walletAddress, email).Delete(&models.Profile{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}

func GetEmailByWalletAddress(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	var profile models.Profile
	if err := db.DB.Select("email").Where("wallet_address = ?", walletAddress).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": profile.Email})
}

func GetProfileByWalletAddress(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	var profile models.Profile
	if err := db.DB.Where("wallet_address = ?", walletAddress).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	// profile.Email = ""

	c.JSON(http.StatusOK, profile)
}

func GetProfileByUsername(c *gin.Context) {
	Username := c.Param("username")
	var profile models.Profile
	if err := db.DB.Where("username = ?", Username).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// profile.Email = ""

	c.JSON(http.StatusOK, profile)
}

func SaveAddresses(c *gin.Context) {
	var addresses []models.Address
	if err := c.ShouldBindJSON(&addresses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profileID := c.Param("profile_id") // Get profile ID from URL parameter
	for i := range addresses {
		addresses[i].ProfileID, _ = uuid.Parse(profileID)
	}

	if err := db.DB.Create(&addresses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func GetAddresses(c *gin.Context) {
	profileID := c.Param("profile_id")
	var addresses []models.Address
	if err := db.DB.Where("profile_id = ?", profileID).Find(&addresses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Addresses not found"})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func UpdateAddress(c *gin.Context) {
	profileID := c.Param("profile_id")
	var addresses []models.Address

	if err := c.ShouldBindJSON(&addresses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, address := range addresses {
		profileUUID, err := uuid.Parse(profileID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID format"})
			return
		}
		address.ProfileID = profileUUID

		if err := db.DB.Save(&address).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, addresses)
}

func DeleteAddress(c *gin.Context) {
	ID := c.Param("id")
	if err := db.DB.Delete(&models.Address{}, "id = ?", ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted"})
}

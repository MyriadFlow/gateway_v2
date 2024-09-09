package controllers

import (
	"net/http"
	"app.myriadflow.com/utils"
	"app.myriadflow.com/models"
	"github.com/gin-gonic/gin"
)

func SendOTPHandler(c *gin.Context) {
    var request struct {
        Email string `json:"email"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Generate OTP and expiration time
    otp, expiration := utils.GenerateOTP()

    // Send OTP via email
    if err := utils.SendOTP(request.Email, otp); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    // Save OTP and expiration in the store (model)
    models.SaveOTP(request.Email, otp, expiration)
    
    c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

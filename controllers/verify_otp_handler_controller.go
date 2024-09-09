package controllers

import (
    "net/http"
	"app.myriadflow.com/models"
    "github.com/gin-gonic/gin"
)

func VerifyOTPHandler(c *gin.Context) {
    var request struct {
        Email string `json:"email"`
        OTP   string `json:"otp"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Verify OTP from the store (model)
    if models.VerifyOTP(request.Email, request.OTP) {
        c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
    }
}

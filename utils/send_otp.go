package utils

import (
	"fmt"
	"net/smtp"
)

func SendOTP(email, otp string) error {
    from := "contact@myriadflow.com"
	// password := os.Getenv("SMTP_PASSWORD")
	password := ""
	smtpHost := "smtp.zoho.com"
	smtpPort := "587"

    msg := []byte(fmt.Sprintf("Subject: OTP Verification\n\nYour OTP is: %s ", otp))

    auth := smtp.PlainAuth("", from, password, smtpHost)
    return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, msg)
}

package utils

import (
    "crypto/rand"
    "time"
)

func GenerateOTP() (string, time.Time) {
    otp := make([]byte, 6)
    rand.Read(otp)
    for i := range otp {
        otp[i] = (otp[i] % 10) + '0' // generates numbers 0-9
    }
    // OTP valid for 5 minutes
    expiration := time.Now().Add(5 * time.Minute)
    return string(otp), expiration
}

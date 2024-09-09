package models

import (
    "sync"
    "time"
)

type OTPData struct {
    otp        string
    expiration time.Time
}

type OTPStore struct {
    data map[string]OTPData
    mu   sync.RWMutex
}

var store = OTPStore{
    data: make(map[string]OTPData),
}

func SaveOTP(email, otp string, expiration time.Time) {
    store.mu.Lock()
    defer store.mu.Unlock()
    store.data[email] = OTPData{
        otp:        otp,
        expiration: expiration,
    }
}

func VerifyOTP(email, otp string) bool {
    store.mu.RLock()
    defer store.mu.RUnlock()

    // Fetch the stored OTP data for the email
    data, exists := store.data[email]

    // Return false if the OTP doesn't exist or if it has expired
    if !exists || time.Now().After(data.expiration) {
        return false
    }

    // Return true if the OTP matches
    return data.otp == otp
}

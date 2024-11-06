package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	FullName       string    `json:"full_name"`
	StreetAddress  string    `json:"street_address"`
	StreetAddress2 string    `json:"street_address_2"`
	City           string    `json:"city"`
	Pincode        string    `json:"pincode"`
	Country        string    `json:"country"`
	ProfileID      uuid.UUID `gorm:"type:uuid" json:"profile_id"` // Foreign key
}

type Profile struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	WalletAddress      string    `json:"wallet_address"`
	CoverImage         string    `json:"cover_image"`
	ProfileImage       string    `json:"profile_image"`
	Username           string    `json:"username"`
	Bio                string    `json:"bio"`
	Website            string    `json:"website"`
	X                  string    `json:"x"`
	Instagram          string    `json:"instagram"`
	Basename           string    `json:"basename"`
	Discord            string    `json:"discord"`
	ChaintypeID        uuid.UUID `gorm:"type:uuid" json:"chaintype_id"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	Addresses          []Address `gorm:"foreignKey:ProfileID"`
	SelectedSocialLink string    `json:"selected_social_link"`
	Link               string    `json:"link"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}

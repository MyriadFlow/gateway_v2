package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Elevate struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	FullName              string    `json:"full_name"`
	EmailAddress          string    `json:"email_address"`
	BrandDescription      string    `json:"brand_description"`
	ProgramAlignment      string    `json:"program_alignment"`
	BrandVision           string    `json:"brand_vision"`
	AdditionalInformation string    `json:"additional_information"`
	Status                string    `json:"status"`
	WalletAddress         string    `json:"wallet_address"`
	ChaintypeID           uuid.UUID `gorm:"type:uuid" json:"chaintype_id"`
	CreatedAt             time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt             time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (e *Elevate) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

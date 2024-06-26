package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brand struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	LogoImage       string    `json:"logo_image"`
	CoverImage      string    `json:"cover_image"`
	Representative  string    `json:"representative"`
	ContactEmail    string    `json:"contact_email"`
	ContactPhone    string    `json:"contact_phone"`
	ShippingAddress string    `json:"shipping_address"`
	AdditionalInfo  string    `json:"additional_info"`
	Industry        string    `json:"industry"`
	Tags            string    `json:"tags"`
	Fees            int       `json:"fees"`
	PayoutAddress   string    `json:"payout_address"`
	AccessMaster    string    `json:"access_master"`
	TradeHub        string    `json:"trade_hub"`
	Blockchain      string    `json:"blockchain"`
	ChainID         string    `json:"chain_id"`
	ManagerID       int       `json:"manager_id"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (b *Brand) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

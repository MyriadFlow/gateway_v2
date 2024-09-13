package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DelegateMintFanTokenRequest struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatorWallet string    `json:"creatorWallet"`
	TokenID       string    `json:"token_id"` // Using string to store big.Int
	Amount        string    `json:"amount"` // Using string to store big.Int
	Data          string    `json:"data"`
	TxHash        string    `json:"txHash"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (a *DelegateMintFanTokenRequest) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

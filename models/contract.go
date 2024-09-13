package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MainnetFanToken struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	NFTContractAddress string    `json:"nftContractAddress"`
	Data               string    `json:"data"`
	URI                string    `json:"uri"`
	TxHash             string    `json:"txHash"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (a *MainnetFanToken) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

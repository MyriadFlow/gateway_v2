package models

import(
"time"
"github.com/google/uuid"
)

type FanToken struct {
	ID          	uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	BrandID     	string          `json:"brand_id"`
	CollectionID    uuid.UUID  		`json:"collection_id"`
	PhygitalID      string          `json:"phygital_id"`
	TokenFestID		string     		`json:"tokenFest_id"`
	CreatedAt   	time.Time       `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt   	time.Time       `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}


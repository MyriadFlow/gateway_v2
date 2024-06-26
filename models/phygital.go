package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Phygital struct {
	ID              uuid.UUID              `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name            string                 `json:"name"`
	BrandName       string                 `json:"brand_name"`
	Category        map[string]interface{} `gorm:"type:jsonb" json:"category"`
	Description     string                 `json:"description"`
	Price           int                    `json:"price"`
	Quantity        int                    `json:"quantity"`
	Royality        int                    `json:"royality"`
	Image           string                 `json:"image"`
	ProductInfo     string                 `json:"product_info"`
	Color           string                 `json:"color"`
	Size            string                 `json:"size"`
	Weight          int                    `json:"weight"`
	Material        string                 `json:"material"`
	Usage           string                 `json:"usage"`
	Quality         string                 `json:"quality"`
	Manufacturer    string                 `json:"manufacturer"`
	OriginCountry   string                 `json:"origin_country"`
	MetadataURI     string                 `json:"metadata_uri"`
	DeployerAddress string                 `json:"deployer_address"`
	ContractAddress string                 `json:"contract_address"`
	GraphURL        string                 `json:"graph_url"`
	CollectionID    uuid.UUID              `json:"collection_id"`
	CreatedAt       time.Time              `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time              `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (p *Phygital) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}

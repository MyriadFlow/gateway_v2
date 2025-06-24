package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brand struct {
	ID                          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name                        string    `json:"name"`
	SlugName                    string    `gorm:"uniqueIndex" json:"slug_name"`
	AgentId                     string    `json:"agent_id"`
	AvatarId                    string    `json:"avatar_id"`
	Slogan                      string    `json:"slogan"`
	Description                 string    `json:"description"`
	LogoImage                   string    `json:"logo_image"`
	CoverImage                  string    `json:"cover_image"`
	Representative              string    `json:"representative"`
	ContactEmail                string    `json:"contact_email"`
	ContactPhone                string    `json:"contact_phone"`
	ShippingAddress             string    `json:"shipping_address"`
	Website                     string    `json:"website"`
	Twitter                     string    `json:"twitter"`
	Instagram                   string    `json:"instagram"`
	Facebook                    string    `json:"facebook"`
	Telegram                    string    `json:"telegram"`
	LinkedIn                    string    `json:"linkedin"`
	Youtube                     string    `json:"youtube"`
	Discord                     string    `json:"discord"`
	Whatsapp                    string    `json:"whatsapp"`
	Google                      string    `json:"google"`
	Tiktok                      string    `json:"tiktok"`
	Snapchat                    string    `json:"snapchat"`
	Pinetrest                   string    `json:"pinetrest"`
	AdditionalLink              string    `json:"additional_link"`
	Link                        string    `json:"link"`
	AdditionalInfo              string    `json:"additional_info"`
	Industry                    string    `json:"industry"`
	Tags                        string    `json:"tags"`
	Fees                        float64   `json:"fees" gorm:"type:decimal(20,10)"`
	PayoutAddress               string    `json:"payout_address"`
	AccessMaster                string    `json:"access_master"`
	TradeHub                    string    `json:"trade_hub"`
	Blockchain                  string    `json:"blockchain"`
	ChainID                     string    `json:"chain_id"`
	ChaintypeID                 uuid.UUID `gorm:"type:uuid" json:"chaintype_id"`
	ManagerID                   string    `json:"manager_id"` //user walletaddress
	ElevateRegion               string    `json:"elevate_region"`
	WebXRExperienceWithAiAvatar bool      `json:"webxr_experience_with_ai_avatar"`
	Image360                    string    `json:"image360"`
	Video360                    string    `json:"video360"`
	CreatedAt                   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt                   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (b *Brand) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

package models

import "github.com/google/uuid"

type Agent struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	AgentID         string    `json:"agentId"`
	Name            string    `json:"agentName"`
	Clients         []string  `json:"agentClients" gorm:"type:jsonb"`
	Port            int       `json:"agentPort"`
	Domain          string    `json:"agentDomain"`
	Status          string    `json:"agentStatus"`
	Image           string    `json:"agentImage"`
	Voice           string    `json:"agentVoice"`
	Avatar          string    `json:"avatar"`
	AgentCategoryID string    `json:"agentCategoryId"`
	CategoryType    string    `json:"categoryType"`
}

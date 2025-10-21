package entity

import (
	"time"

	"gorm.io/gorm"
)

type ContentSetting struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	UserID       string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID   string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	TopHeader    []string       `gorm:"type:text[]" json:"top_header"`
	ButtonHeader []string       `gorm:"type:text[]" json:"button_header"`
	Feature      []string       `gorm:"type:text[]" json:"feature"`
	Footer       []string       `gorm:"type:text[]" json:"footer"`
	CreatedBy    string         `json:"created_by"`
	UpdatedBy    string         `json:"updated_by"`
	DeletedBy    string         `json:"deleted_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

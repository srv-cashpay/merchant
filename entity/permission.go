package entity

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	Label      string         `gorm:"label,omitempty" json:"label"`
	Icon       string         `gorm:"icon" json:"icon"`
	To         string         `gorm:"to" json:"to"`
	CreatedBy  string         `gorm:"created_by" json:"created_by"`
	UpdatedBy  string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy  string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

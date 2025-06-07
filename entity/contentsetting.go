package entity

import (
	"time"

	"gorm.io/gorm"
)

type ContentSetting struct {
	ID          string         `gorm:"primary_key,omitempty" json:"id"`
	UserID      string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	Logo        string         `gorm:"logo,omitempty" json:"logo"`
	Title       string         `gorm:"title,omitempty" json:"title"`
	Description string         `gorm:"description" json:"description"`
	LinkAndroid string         `gorm:"link_android" json:"link_android"`
	LinkApple   string         `gorm:"link_apple" json:"link_apple"`
	CreatedBy   string         `gorm:"created_by" json:"created_by"`
	UpdatedBy   string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy   string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

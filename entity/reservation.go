package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID         string         `gorm:"primary_key,omitempty" json:"id"`
	UserID     string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	Table      string         `gorm:"order_id,omitempty" json:"order_id"`
	Status     string         `gorm:"status" json:"status"`
	CreatedBy  string         `gorm:"created_by" json:"created_by"`
	DeletedBy  string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID              string            `gorm:"primary_key,omitempty" json:"id"`
	UserID          string            `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID      string            `gorm:"type:varchar(36);index" json:"merchant_id"`
	Nomor           int               `gorm:"nomor,omitempty" json:"nomor"`
	VoucherGenerate []VoucherGenerate `gorm:"foreignKey:VoucherID;constraint:OnDelete:CASCADE;" json:"voucher_generate"`
	CreatedBy       string            `gorm:"created_by" json:"created_by"`
	UpdatedBy       string            `gorm:"updated_by" json:"updated_ by"`
	DeletedBy       string            `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}

type VoucherGenerate struct {
	ID          string `gorm:"primaryKey" json:"id"`
	VoucherID   string `gorm:"type:varchar(36);index" json:"voucher_id"`
	MerchantID  string `gorm:"type:varchar(36);index" json:"merchant_id"`
	VoucherName string `gorm:"voucher_name,omitempty" json:"voucher_name"`
	VoucherLink string `gorm:"voucher_link,omitempty" json:"voucher_link"`
	StartDate   string `gorm:"start_date,omitempty" json:"start_date"`
	EndDate     string `gorm:"end_date,omitempty" json:"end_date"`
	Status      bool   `gorm:"status,omitempty" json:"status"`
}

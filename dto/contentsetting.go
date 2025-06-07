package dto

import "time"

type GetByContentSettingIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type ContentSettingRequest struct {
	ID          string `gorm:"primary_key,omitempty" json:"id"`
	UserID      string `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string `gorm:"type:varchar(36);index" json:"merchant_id"`
	Logo        string `gorm:"logo,omitempty" json:"logo"`
	Title       string `gorm:"title,omitempty" json:"title"`
	Description string `gorm:"description" json:"description"`
	LinkAndroid string `gorm:"link_android" json:"link_android"`
	LinkApple   string `gorm:"link_apple" json:"link_apple"`
}

type ContentSettingResponse struct {
	ID          string `gorm:"primary_key,omitempty" json:"id"`
	UserID      string `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string `gorm:"type:varchar(36);index" json:"merchant_id"`
	Logo        string `gorm:"logo,omitempty" json:"logo"`
	Title       string `gorm:"title,omitempty" json:"title"`
	Description string `gorm:"description" json:"description"`
	LinkAndroid string `gorm:"link_android" json:"link_android"`
	LinkApple   string `gorm:"link_apple" json:"link_apple"`
}

type UpdateContentSettingRequest struct {
	ID          string    `gorm:"primary_key,omitempty" json:"id"`
	UserID      string    `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string    `gorm:"type:varchar(36);index" json:"merchant_id"`
	Logo        string    `gorm:"logo,omitempty" json:"logo"`
	Title       string    `gorm:"title,omitempty" json:"title"`
	Description string    `gorm:"description" json:"description"`
	LinkAndroid string    `gorm:"link_android" json:"link_android"`
	LinkApple   string    `gorm:"link_apple" json:"link_apple"`
	UpdatedBy   string    `json:"update_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateContentSettingResponse struct {
	ID          string    `gorm:"primary_key,omitempty" json:"id"`
	UserID      string    `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string    `gorm:"type:varchar(36);index" json:"merchant_id"`
	Logo        string    `gorm:"logo,omitempty" json:"logo"`
	Title       string    `gorm:"title,omitempty" json:"title"`
	Description string    `gorm:"description" json:"description"`
	LinkAndroid string    `gorm:"link_android" json:"link_android"`
	LinkApple   string    `gorm:"link_apple" json:"link_apple"`
	UpdatedBy   string    `json:"update_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

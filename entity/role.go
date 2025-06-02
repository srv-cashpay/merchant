package entity

type Role struct {
	ID         string `gorm:"primary_key,omitempty" json:"id"`
	Role       string `gorm:"type:varchar(36)" json:"role"`
	UserID     string `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string `gorm:"type:varchar(36);index" json:"merchant_id"`
}

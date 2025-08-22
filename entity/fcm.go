package entity

type FCMToken struct {
	UserID string `gorm:"primaryKey"`
	Token  string `gorm:"type:varchar(256);index" json:"token"`
}

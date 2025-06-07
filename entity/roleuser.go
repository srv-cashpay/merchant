package entity

import "time"

type RoleUser struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID       string    `gorm:"type:varchar(36);index,omitempty" json:"role_id"`
	UserID       string    `gorm:"type:varchar(36);index,omitempty" json:"user_id"`
	PermissionID uint      `gorm:"permission_id;index,omitempty" json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}

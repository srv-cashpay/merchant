package entity

type RoleUserPermission struct {
	ID           uint `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleUserID   uint `gorm:"role_id;index,omitempty" json:"role_id"`
	PermissionID uint `gorm:"permission_id;index,omitempty" json:"permission_id"`
}

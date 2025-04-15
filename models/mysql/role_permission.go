package mysql

type RolePermission struct {
	BaseModel      `json:",inline"`
	RoleId         uint64 `json:"name,omitempty" gorm:"column:name"`
	PermissionId   uint64 `json:"permission_id,omitempty" gorm:"column:permission_id"`
	PermissionType uint64 `json:"permission_type,omitempty" gorm:"column:permission_type"`
}

func (m *RolePermission) TableName() string {
	return "role_permissions"
}

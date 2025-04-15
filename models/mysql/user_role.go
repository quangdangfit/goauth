package mysql

type UserRole struct {
	BaseModel `json:",inline"`
	UserId    int64 `json:"user_id,omitempty" gorm:"column:user_id"`
	RoleId    int64 `json:"role_id,omitempty" gorm:"column:role_id"`
}

func (m *UserRole) TableName() string {
	return "user_roles"
}

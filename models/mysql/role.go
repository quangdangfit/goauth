package mysql

type Role struct {
	BaseModel   `json:",inline"`
	Name        string `json:"name,omitempty" gorm:"column:name"`
	Description string `json:"phone,omitempty" gorm:"column:phone"`
}

func (m *Role) TableName() string {
	return "roles"
}

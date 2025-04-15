package mysql

type Group struct {
	BaseModel   `json:",inline"`
	Name        string  `gorm:"unique;not null" json:"name"`
	Description string  `json:"description"`
	ParentId    *uint64 `json:"parent_id,omitempty"`
}

func (m *Group) TableName() string {
	return "groups"
}

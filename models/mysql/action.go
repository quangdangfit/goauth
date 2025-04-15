package mysql

type Action struct {
	BaseModel   `json:",inline"`
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `json:"description"`
	API         string `json:"api,omitempty"`
	Method      string `json:"method,omitempty"`
	GroupId     uint64 `json:"group_id,omitempty"`
}

func (m *Action) TableName() string {
	return "actions"
}

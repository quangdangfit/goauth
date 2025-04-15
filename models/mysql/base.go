package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBaseModel interface {
	TableName() string
	BeforeInsert() error
}

type BaseModel struct {
	Id        uint64         `json:"id" gorm:"column:id"`
	UUID      string         `json:"uuid" gorm:"column:uuid;primaryKey;type:varchar(100)"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (m *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	m.UUID = uuid.NewString()
	return nil
}

func (m *BaseModel) IsDeleted() bool {
	return !m.DeletedAt.Time.IsZero() && m.DeletedAt.Valid
}

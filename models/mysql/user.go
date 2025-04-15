package mysql

import (
	"github.com/rinard84/auth-service/models/types"
	api "github.com/rinard84/proto/api/auth"
	"gorm.io/gorm"
)

type User struct {
	BaseModel     `json:",inline"`
	Username      string           `json:"user_name,omitempty" gorm:"column:user_name"`
	Password      string           `json:"password,omitempty" gorm:"column:password"`
	Name          string           `json:"name,omitempty" gorm:"column:name"`
	Phone         string           `json:"phone,omitempty" gorm:"column:phone"`
	Email         string           `json:"email,omitempty" gorm:"column:email"`
	VerifiedPhone bool             `json:"verified_phone,omitempty" gorm:"column:verified_phone"`
	VerifiedEmail bool             `json:"verified_email,omitempty" gorm:"column:verified_email"`
	Status        types.UserStatus `json:"status,omitempty" gorm:"status;index;type:varchar(30)"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	err = m.BaseModel.BeforeCreate(tx)
	if err != nil {
		return err
	}

	if m.Status == "" {
		m.Status = types.UserStatusActive
	}
	return nil
}

func (m *User) Proto() *api.User {
	return &api.User{
		Id:            int64(m.Id),
		Username:      m.Username,
		Name:          m.Name,
		Phone:         m.Phone,
		Email:         m.Email,
		VerifiedPhone: m.VerifiedPhone,
		VerifiedEmail: m.VerifiedEmail,
	}
}

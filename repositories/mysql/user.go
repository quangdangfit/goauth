package mysql

import (
	"github.com/go-logr/logr"
	"github.com/quangdangfit/goauth/config"
	models "github.com/quangdangfit/goauth/models/mysql"
	"github.com/quangdangfit/goauth/repositories"
	"gorm.io/gorm"
)

type UserRepository interface {
	repositories.Repository[models.User]
}

type userRepository struct {
	repo[models.User]
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(logger logr.Logger, cfg *config.Config, db *gorm.DB) UserRepository {
	return &userRepository{
		repo[models.User]{
			logger: logger,
			cfg:    cfg,
			db:     db,
		},
	}
}

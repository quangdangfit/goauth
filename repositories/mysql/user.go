package mysql

import (
	"github.com/go-logr/logr"
	"github.com/rinard84/auth-service/config"
	models "github.com/rinard84/auth-service/models/mysql"
	"github.com/rinard84/auth-service/repositories"
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

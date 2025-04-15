package mysql

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/rinard84/auth-service/config"
	"github.com/rinard84/gokit/library/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type repo[Model any] struct {
	logger logr.Logger
	cfg    *config.Config
	db     *gorm.DB
}

func (r repo[Model]) WithTransaction(ctx context.Context, f func() error) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	txFn := func(tx *gorm.DB) error {
		return f()
	}
	return r.db.WithContext(ctx).Transaction(txFn)
}

func (r repo[Model]) Create(ctx context.Context, doc *Model) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx).Model(new(Model))
	if err := db.Create(&doc).Error; err != nil {
		return err
	}
	return nil
}

func (r repo[Model]) Update(ctx context.Context, uuid string, doc *Model) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx).Model(new(Model))
	if err := db.Where("uuid = ?", uuid).Updates(&doc).Error; err != nil {
		return err
	}
	return nil
}

func (r repo[Model]) Delete(ctx context.Context, uuid string) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx)
	if err := db.Delete(new(Model), "uuid = ?", uuid).Error; err != nil {
		return err
	}
	return nil
}

func (r repo[Model]) GetById(ctx context.Context, uuid string) (*Model, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx).Model(new(Model))
	var res Model
	if err := db.First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r repo[Model]) GetOne(ctx context.Context, filter map[string]interface{}, sort []string) (*Model, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx).Model(new(Model))
	var res Model
	if err := db.Where(filter).Order(sort).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r repo[Model]) GetList(ctx context.Context, filter map[string]interface{}, limit, offset int, sort []string) ([]Model, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx).Model(new(Model))
	var res []Model
	if err := db.Where(filter).Limit(limit).Offset(offset).Order(sort).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r repo[Model]) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	db := r.db.WithContext(ctx).Model(new(Model))
	var total int64
	if err := db.Where(filter).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func newDB(cfg database.MySQLConfig, env string) (*gorm.DB, error) {
	// force a connection and test that it worked
	gormConfig := &gorm.Config{} //nolint
	if env != config.ProductionEnvironment {
		gormConfig.Logger = gormlog.Default.LogMode(gormlog.Info)
	}
	db, err := gorm.Open(mysql.Open(cfg.DSN()), gormConfig)
	if err != nil {
		panic(err)
	}

	if env != config.ProductionEnvironment {
		db = db.Debug()
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// force a connection and test that it worked
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

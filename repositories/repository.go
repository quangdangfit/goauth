package repositories

import (
	"context"
)

type Repository[Model any] interface {
	WithTransaction(ctx context.Context, f func() error) error
	Create(ctx context.Context, doc *Model) error
	Update(ctx context.Context, uuid string, doc *Model) error
	Delete(ctx context.Context, uuid string) error
	GetById(ctx context.Context, uuid string) (*Model, error)
	GetOne(ctx context.Context, filter map[string]interface{}, sort []string) (*Model, error)
	GetList(ctx context.Context, filter map[string]interface{}, limit, offset int, sort []string) ([]Model, error)
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}

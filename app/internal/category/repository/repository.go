package repository

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/category"
)

type Repository interface {
	GetAll() ([]category.Category, error)
	GetAllWithContext(ctx context.Context) ([]category.Category, error)
	GetByName(name string) (*category.Category, error)
	GetByNameWithContext(ctx context.Context, name string) (*category.Category, error)
}

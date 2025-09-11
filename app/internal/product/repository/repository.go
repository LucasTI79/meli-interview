package repository

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/product"
)

type Repository interface {
	GetAll() ([]product.Product, error)
	GetByID(productId string) (*product.Product, error)
	GetAllWithContext(ctx context.Context) ([]product.Product, error)
	GetByIDWithContext(ctx context.Context, productId string) (*product.Product, error)
}

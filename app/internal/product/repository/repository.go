package repository

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/product"
)

type Repository interface {
	GetAll(filters product.ProductFilter) ([]product.Product, int, error)
	GetByID(productId string) (*product.Product, error)
	GetAllWithContext(ctx context.Context, filters product.ProductFilter) ([]product.Product, int, error)
	GetByIDWithContext(ctx context.Context, productId string) (*product.Product, error)
}

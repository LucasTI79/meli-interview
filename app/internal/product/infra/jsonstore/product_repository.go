package jsonstore

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/infra/jsonstore"
	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/repository"
)

type productRepository struct {
	repo *jsonstore.JSONRepository[product.Product]
}

func NewProductRepository(fileName string) (repository.Repository, error) {
	getID := func(entity product.Product) string {
		return entity.Id
	}
	repo, err := jsonstore.NewJSONRepository(fileName, getID)
	if err != nil {
		return nil, err
	}
	return &productRepository{repo: repo}, nil
}

func (r *productRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	err := r.repo.FindAll(func(entity product.Product) error {
		products = append(products, entity)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetByID(productId string) (*product.Product, error) {
	product, err := r.repo.FindByID(productId)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAllWithContext(ctx context.Context) ([]product.Product, error) {
	var products []product.Product
	err := r.repo.FindAll(func(entity product.Product) error {
		products = append(products, entity)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetByIDWithContext(ctx context.Context, productId string) (*product.Product, error) {
	product, err := r.repo.FindByID(productId)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

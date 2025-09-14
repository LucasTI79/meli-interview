package jsonstore

import (
	"context"
	"strings"

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

func (r *productRepository) GetAll(filters product.ProductFilter) ([]product.Product, int, error) {
	var result []product.Product

	total, err := r.repo.FindAllWherePaginated(
		func(p product.Product) bool {
			return matchProduct(p, filters)
		},
		filters.Page,
		filters.PageSize,
		func(p product.Product) error {
			result = append(result, p)
			return nil
		},
	)

	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (r *productRepository) GetByID(productId string) (*product.Product, error) {
	product, err := r.repo.FindByID(productId)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAllWithContext(ctx context.Context, filters product.ProductFilter) ([]product.Product, int, error) {
	var result []product.Product

	total, err := r.repo.FindAllWherePaginated(
		func(p product.Product) bool {
			return matchProduct(p, filters)
		},
		filters.Page,
		filters.PageSize,
		func(p product.Product) error {
			result = append(result, p)
			return nil
		},
	)

	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (r *productRepository) GetByIDWithContext(ctx context.Context, productId string) (*product.Product, error) {
	product, err := r.repo.FindByID(productId)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func matchProduct(p product.Product, f product.ProductFilter) bool {
	if f.Name != "" && !strings.Contains(strings.ToLower(p.Name), strings.ToLower(f.Name)) {
		return false
	}

	if len(f.Categories) > 0 {
		found := false
		for _, cat := range f.Categories {
			if strings.EqualFold(cat, p.Category) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if f.MinPrice > 0 && p.Price < f.MinPrice {
		return false
	}
	if f.MaxPrice > 0 && p.Price > f.MaxPrice {
		return false
	}
	return true
}

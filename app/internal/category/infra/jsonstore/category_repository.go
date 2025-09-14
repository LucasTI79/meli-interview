package jsonstore

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/category"
	"github.com/lucasti79/meli-interview/internal/category/repository"
	"github.com/lucasti79/meli-interview/internal/infra/jsonstore"
	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
)

type categoryRepository struct {
	repo *jsonstore.JSONRepository[product.Product]
}

func NewCategoryRepository(fileName string) (repository.Repository, error) {
	getID := func(entity product.Product) string {
		return entity.Category
	}
	repo, err := jsonstore.NewJSONRepository(fileName, getID)
	if err != nil {
		return nil, err
	}
	return &categoryRepository{repo: repo}, nil
}

func (r *categoryRepository) GetAll() ([]category.Category, error) {
	catMap := make(map[string]struct{})

	err := r.repo.FindAll(func(p product.Product) error {
		catMap[p.Category] = struct{}{}
		return nil
	})
	if err != nil {
		return nil, err
	}

	categories := make([]category.Category, 0, len(catMap))
	for c := range catMap {
		categories = append(categories, category.Category{Name: c})
	}

	return categories, nil
}

func (r *categoryRepository) GetAllWithContext(ctx context.Context) ([]category.Category, error) {
	catMap := make(map[string]struct{})

	err := r.repo.FindAll(func(p product.Product) error {
		catMap[p.Category] = struct{}{}
		return nil
	})
	if err != nil {
		return nil, err
	}

	categories := make([]category.Category, 0, len(catMap))
	for c := range catMap {
		categories = append(categories, category.Category{Name: c})
	}

	return categories, nil
}

func (r *categoryRepository) GetByName(name string) (*category.Category, error) {
	var found bool
	err := r.repo.FindAllWhere(
		func(p product.Product) bool {
			if p.Category == name {
				found = true
				return true
			}
			return false
		},
		func(p product.Product) error {
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	if found {
		return &category.Category{Name: name}, nil
	}
	return nil, apperrors.ErrResourceNotExists
}

func (r *categoryRepository) GetByNameWithContext(ctx context.Context, name string) (*category.Category, error) {
	var found bool
	err := r.repo.FindAllWhere(
		func(p product.Product) bool {
			if p.Category == name {
				found = true
				return true
			}
			return false
		},
		func(p product.Product) error {
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	if found {
		return &category.Category{Name: name}, nil
	}
	return nil, apperrors.ErrResourceNotExists
}

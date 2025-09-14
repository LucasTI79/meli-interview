package service

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/category"
	"github.com/lucasti79/meli-interview/internal/category/repository"
)

type service struct {
	repo repository.Repository
}

type Service interface {
	GetAll() ([]category.Category, error)
	GetAllWithContext(ctx context.Context) ([]category.Category, error)
	GetByName(name string) (*category.Category, error)
	GetByNameWithContext(ctx context.Context, name string) (*category.Category, error)
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() ([]category.Category, error) {
	return s.repo.GetAll()
}

func (s *service) GetAllWithContext(ctx context.Context) ([]category.Category, error) {
	return s.repo.GetAllWithContext(ctx)
}

func (s *service) GetByName(name string) (*category.Category, error) {
	categories, err := s.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *service) GetByNameWithContext(ctx context.Context, name string) (*category.Category, error) {
	categories, err := s.repo.GetByNameWithContext(ctx, name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

package service

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/category"
	"github.com/lucasti79/meli-interview/internal/category/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s *Service) GetAll() ([]category.Category, error) {
	return s.repo.GetAll()
}

func (s *Service) GetAllWithContext(ctx context.Context) ([]category.Category, error) {
	return s.repo.GetAllWithContext(ctx)
}

func (s *Service) GetByName(name string) (*category.Category, error) {
	categories, err := s.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *Service) GetByNameWithContext(ctx context.Context, name string) (*category.Category, error) {
	categories, err := s.repo.GetByNameWithContext(ctx, name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

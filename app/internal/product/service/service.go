package service

import (
	"context"

	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s *Service) GetAll() ([]product.Product, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(productId string) (*product.Product, error) {
	return s.repo.GetByID(productId)
}

func (s *Service) GetAllWithContext(ctx context.Context) ([]product.Product, error) {
	return s.repo.GetAllWithContext(ctx)
}

func (s *Service) GetByIDWithContext(ctx context.Context, productId string) (*product.Product, error) {
	pr, err := s.repo.GetByIDWithContext(ctx, productId)

	if err != nil {
		return nil, err
	}

	return pr, nil
}

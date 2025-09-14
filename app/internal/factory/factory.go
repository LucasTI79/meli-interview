package factory

import (
	CategoryApi "github.com/lucasti79/meli-interview/internal/category/api"
	CategoryJsonRepository "github.com/lucasti79/meli-interview/internal/category/infra/jsonstore"
	CategoryRepository "github.com/lucasti79/meli-interview/internal/category/repository"
	CategoryService "github.com/lucasti79/meli-interview/internal/category/service"
	ProductApi "github.com/lucasti79/meli-interview/internal/product/api"
	ProductJsonRepository "github.com/lucasti79/meli-interview/internal/product/infra/jsonstore"
	ProductRepository "github.com/lucasti79/meli-interview/internal/product/repository"
	ProductService "github.com/lucasti79/meli-interview/internal/product/service"
)

type AppFactory struct {
	ProductHandler  *ProductApi.Handler
	CategoryHandler *CategoryApi.Handler
}

func NewProductHandler(repo ProductRepository.Repository) (*ProductApi.Handler, error) {
	if repo == nil {
		var err error
		repo, err = ProductJsonRepository.NewProductRepository("products.jsonl")
		if err != nil {
			return nil, err
		}
	}

	service := ProductService.NewService(repo)
	handler := ProductApi.NewHandler(service)
	return handler, nil
}

func NewCategoryHandler(repo CategoryRepository.Repository) (*CategoryApi.Handler, error) {
	if repo == nil {
		var err error
		repo, err = CategoryJsonRepository.NewCategoryRepository("products.jsonl")
		if err != nil {
			return nil, err
		}
	}

	service := CategoryService.NewService(repo)
	handler := CategoryApi.NewHandler(service)
	return handler, nil
}

func NewAppFactory() (*AppFactory, error) {
	productHandler, err := NewProductHandler(nil)

	if err != nil {
		return nil, err
	}

	categoryHandler, err := NewCategoryHandler(nil)
	if err != nil {
		return nil, err
	}

	return &AppFactory{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
	}, nil
}

var appFactory *AppFactory

func InitFactory() error {
	var err error
	appFactory, err = NewAppFactory()
	return err
}

func GetFactory() *AppFactory {
	if appFactory == nil {
		panic("AppFactory not initialized")
	}
	return appFactory
}

package factory

import (
	ProductApi "github.com/lucasti79/meli-interview/internal/product/api"
	ProductJsonRepository "github.com/lucasti79/meli-interview/internal/product/infra/jsonstore"
	ProductRepository "github.com/lucasti79/meli-interview/internal/product/repository"
	ProductService "github.com/lucasti79/meli-interview/internal/product/service"
)

type AppFactory struct {
	ProductHandler *ProductApi.Handler
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

func NewAppFactory() (*AppFactory, error) {
	productHandler, err := NewProductHandler(nil)
	if err != nil {
		return nil, err
	}

	return &AppFactory{
		ProductHandler: productHandler,
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

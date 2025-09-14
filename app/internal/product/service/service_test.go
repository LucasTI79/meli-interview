package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/infra/mocks"
	"github.com/lucasti79/meli-interview/internal/product/service"
	"github.com/stretchr/testify/assert"
)

func TestService_GetAll(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	filters := product.ProductFilter{}
	expectedProducts := []product.Product{
		{Id: "1", Name: "Product 1"},
		{Id: "2", Name: "Product 2"},
	}

	// Configura o retorno do mock
	mockRepo.On("GetAll", filters).Return(expectedProducts, len(expectedProducts), nil).Once()

	products, total, err := svc.GetAll(filters)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	assert.Equal(t, 2, total)

	mockRepo.AssertExpectations(t)
}

func TestService_GetByID_Success(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	expectedProduct := &product.Product{Id: "1", Name: "Product 1"}

	mockRepo.On("GetByID", "1").Return(expectedProduct, nil).Once()

	pr, err := svc.GetByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, pr)

	mockRepo.AssertExpectations(t)
}

func TestService_GetByID_Error(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	mockRepo.On("GetByID", "not-found").Return(nil, errors.New("not found")).Once()

	pr, err := svc.GetByID("not-found")

	assert.Error(t, err)
	assert.Nil(t, pr)

	mockRepo.AssertExpectations(t)
}

func TestService_GetAllWithContext(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	ctx := context.Background()
	filters := product.ProductFilter{}
	expectedProducts := []product.Product{
		{Id: "1", Name: "Product 1"},
	}

	mockRepo.On("GetAllWithContext", ctx, filters).Return(expectedProducts, len(expectedProducts), nil).Once()

	products, total, err := svc.GetAllWithContext(ctx, filters)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	assert.Equal(t, 1, total)

	mockRepo.AssertExpectations(t)
}

func TestService_GetByIDWithContext(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	ctx := context.Background()
	expectedProduct := &product.Product{Id: "1", Name: "Product 1"}

	mockRepo.On("GetByIDWithContext", ctx, "1").Return(expectedProduct, nil).Once()

	pr, err := svc.GetByIDWithContext(ctx, "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, pr)

	mockRepo.AssertExpectations(t)
}

func TestService_GetByIDWithContext_Error(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	ctx := context.Background()

	// Configura o mock para retornar erro
	mockRepo.On("GetByIDWithContext", ctx, "not-found").
		Return(nil, errors.New("not found")).
		Once()

	pr, err := svc.GetByIDWithContext(ctx, "not-found")

	assert.Error(t, err)
	assert.Nil(t, pr)

	mockRepo.AssertExpectations(t)
}

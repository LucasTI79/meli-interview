package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lucasti79/meli-interview/internal/category"
	"github.com/lucasti79/meli-interview/internal/category/infra/mocks"
	"github.com/lucasti79/meli-interview/internal/category/service"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_GetAll(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	expected := []category.Category{
		{Name: "Electronics"},
		{Name: "Books"},
	}

	mockRepo.On("GetAll").Return(expected, nil).Once()

	result, err := svc.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetAllWithContext(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	ctx := context.Background()
	expected := []category.Category{
		{Name: "Electronics"},
	}

	mockRepo.On("GetAllWithContext", ctx).Return(expected, nil).Once()

	result, err := svc.GetAllWithContext(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetByName_Success(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	expected := &category.Category{Name: "Electronics"}

	mockRepo.On("GetByName", "Electronics").Return(expected, nil).Once()

	result, err := svc.GetByName("Electronics")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetByName_Error(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	mockRepo.On("GetByName", "NotFound").Return(nil, errors.New("not found")).Once()

	result, err := svc.GetByName("NotFound")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetByNameWithContext_Success(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	ctx := context.Background()
	expected := &category.Category{Name: "Books"}

	mockRepo.On("GetByNameWithContext", ctx, "Books").Return(expected, nil).Once()

	result, err := svc.GetByNameWithContext(ctx, "Books")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetByNameWithContext_Error(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	svc := service.NewService(mockRepo)

	ctx := context.Background()

	mockRepo.On("GetByNameWithContext", ctx, "NotFound").Return(nil, errors.New("not found")).Once()

	result, err := svc.GetByNameWithContext(ctx, "NotFound")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

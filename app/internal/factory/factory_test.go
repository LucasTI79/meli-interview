package factory_test

import (
	"testing"

	categoryMocks "github.com/lucasti79/meli-interview/internal/category/infra/mocks"
	"github.com/lucasti79/meli-interview/internal/factory"
	productMocks "github.com/lucasti79/meli-interview/internal/product/infra/mocks"
	"github.com/stretchr/testify/require"
)

func TestNewProductHandler_WithNilRepo(t *testing.T) {
	handler, err := factory.NewProductHandler(nil)
	require.NoError(t, err)
	require.NotNil(t, handler)
}

func TestNewProductHandler_WithMockRepo(t *testing.T) {
	mockRepo := new(productMocks.RepositoryMock)
	handler, err := factory.NewProductHandler(mockRepo)
	require.NoError(t, err)
	require.NotNil(t, handler)
}

func TestNewCategoryHandler_WithNilRepo(t *testing.T) {
	handler, err := factory.NewCategoryHandler(nil)
	require.NoError(t, err)
	require.NotNil(t, handler)
}

func TestNewCategoryHandler_WithMockRepo(t *testing.T) {
	mockRepo := new(categoryMocks.RepositoryMock)
	handler, err := factory.NewCategoryHandler(mockRepo)
	require.NoError(t, err)
	require.NotNil(t, handler)
}

func TestNewAppFactory(t *testing.T) {
	appFactory, err := factory.NewAppFactory()
	require.NoError(t, err)
	require.NotNil(t, appFactory)
	require.NotNil(t, appFactory.ProductHandler)
	require.NotNil(t, appFactory.CategoryHandler)
}

func TestInitFactoryAndGetFactory(t *testing.T) {
	// Ensure panic if not initialized
	require.Panics(t, func() {
		factory.GetFactory()
	})

	// Initialize
	err := factory.InitFactory()
	require.NoError(t, err)

	appFactory := factory.GetFactory()
	require.NotNil(t, appFactory)
	require.NotNil(t, appFactory.ProductHandler)
	require.NotNil(t, appFactory.CategoryHandler)
}

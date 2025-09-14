package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/api"
	"github.com/lucasti79/meli-interview/internal/product/infra/mocks"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
	"github.com/lucasti79/meli-interview/pkg/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetAll_Success(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	mockService.On("GetAllWithContext", mock.Anything, mock.AnythingOfType("product.ProductFilter")).
		Return([]product.Product{
			{Id: "1", Name: "Prod1", Category: "Cat1", Price: 10},
		}, 1, nil)

	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products?name=Prod&categories=Cat1&minPrice=5&maxPrice=50&page=1&pageSize=10", nil)
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetAll_NoContent(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	mockService.On("GetAllWithContext", mock.Anything, mock.AnythingOfType("product.ProductFilter")).
		Return([]product.Product{}, 0, nil)

	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetAll_ValidationError(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products?page=-1", nil)
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetAll_ServiceError(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	mockService.On("GetAllWithContext", mock.Anything, mock.AnythingOfType("product.ProductFilter")).
		Return(nil, 0, errors.New("internal error"))

	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	mockService.On("GetByIDWithContext", mock.Anything, "123").
		Return(&product.Product{Id: "123", Name: "Prod123"}, nil)

	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/123", nil)
	req = testutil.WithUrlParam(t, req, "productId", "123")
	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	mockService.On("GetByIDWithContext", mock.Anything, "123").
		Return(nil, apperrors.ErrResourceNotExists)

	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/123", nil)
	req = testutil.WithUrlParam(t, req, "productId", "123")
	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetByID_InternalError(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	mockService.On("GetByIDWithContext", mock.Anything, "123").
		Return(nil, errors.New("internal"))

	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/123", nil)
	req = testutil.WithUrlParam(t, req, "productId", "123")
	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetByID_MissingID(t *testing.T) {
	mockService := new(mocks.ServiceMock)
	h := api.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/", nil)
	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

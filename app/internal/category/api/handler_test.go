package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasti79/meli-interview/internal/category"
	api "github.com/lucasti79/meli-interview/internal/category/api"
	"github.com/lucasti79/meli-interview/internal/category/infra/mocks"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
	"github.com/lucasti79/meli-interview/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetByName_Success(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	expected := &category.Category{Name: "Books"}
	mockSvc.On("GetByNameWithContext", mock.Anything, "Books").Return(expected, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories/Books", nil)
	req = testutil.WithUrlParam(t, req, "categoryName", "Books")
	w := httptest.NewRecorder()

	h.GetByName(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Books")
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetByName_BadRequest(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories/", nil)
	// Não precisa adicionar param → simula ausência
	w := httptest.NewRecorder()

	h.GetByName(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), category.ErrCategoryInvalidID)
}

func TestHandler_GetByName_NotFound(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	mockSvc.On("GetByNameWithContext", mock.Anything, "NotFound").Return(nil, apperrors.ErrResourceNotExists).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories/NotFound", nil)
	req = testutil.WithUrlParam(t, req, "categoryName", "NotFound")
	w := httptest.NewRecorder()

	h.GetByName(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), category.ErrCategoryNotFound)
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetByName_InternalError(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	mockSvc.On("GetByNameWithContext", mock.Anything, "Books").Return(nil, errors.New("some error")).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories/Books", nil)
	req = testutil.WithUrlParam(t, req, "categoryName", "Books")
	w := httptest.NewRecorder()

	h.GetByName(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), apperrors.ErrInternalError.Error())
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetAll_Success(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	categories := []category.Category{{Name: "Electronics"}, {Name: "Books"}}
	mockSvc.On("GetAllWithContext", mock.Anything).Return(categories, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories", nil)
	w := httptest.NewRecorder()

	h.GetAll(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Electronics")
	assert.Contains(t, w.Body.String(), "Books")
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetAll_NoContent(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	mockSvc.On("GetAllWithContext", mock.Anything).Return([]category.Category{}, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories", nil)
	w := httptest.NewRecorder()

	h.GetAll(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetAll_InternalError(t *testing.T) {
	mockSvc := new(mocks.ServiceMock)
	h := api.NewHandler(mockSvc)

	mockSvc.On("GetAllWithContext", mock.Anything).Return(nil, errors.New("some error")).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories", nil)
	w := httptest.NewRecorder()

	h.GetAll(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), apperrors.ErrInternalError.Error())
	mockSvc.AssertExpectations(t)
}

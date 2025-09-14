package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/lucasti79/meli-interview/internal/category"
	"github.com/lucasti79/meli-interview/internal/category/service"
	httpdto "github.com/lucasti79/meli-interview/internal/infra/http"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
	"github.com/lucasti79/meli-interview/pkg/web/response"
)

type Handler struct {
	service   service.Service
	validator *validator.Validate
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
	}
}

// GetAll godoc
// @Summary      Get all categories
// @Description  Get all categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  CategoriesResult
// @Failure      400  {object}  httpdto.ErrorResponse
// @Failure      500  {object}  httpdto.ErrorResponse
// @Router       /api/v1/categories [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllWithContext(r.Context())

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, httpdto.ErrorResponse{
			Code:    apperrors.ErrInternalError.Error(),
			Message: "internal server error",
			Status:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if len(categories) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	result := httpdto.Result[[]category.Category]{
		Data: categories,
	}

	response.JSON(w, http.StatusOK, result)
}

// GetByName godoc
// @Summary      Get a category by name
// @Description  Get a category by name
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        categoryName   path      string  true  "Category Name"
// @Success      200  {object}  CategoryResult
// @Failure      400  {object}  httpdto.ErrorResponse
// @Failure      404  {object}  httpdto.ErrorResponse
// @Failure      500  {object}  httpdto.ErrorResponse
// @Router       /api/v1/categories/{categoryName} [get]
func (h *Handler) GetByName(w http.ResponseWriter, r *http.Request) {
	categoryName := chi.URLParam(r, "categoryName")

	if categoryName == "" {
		response.JSON(w, http.StatusBadRequest, httpdto.ErrorResponse{
			Code:    category.ErrCategoryInvalidID,
			Message: "category name is required",
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	cat, err := h.service.GetByNameWithContext(r.Context(), categoryName)

	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrResourceNotExists):
			response.JSON(w, http.StatusNotFound, httpdto.ErrorResponse{
				Code:    category.ErrCategoryNotFound,
				Message: err.Error(),
				Status:  http.StatusText(http.StatusNotFound),
			})
		default:
			response.JSON(w, http.StatusInternalServerError, httpdto.ErrorResponse{
				Code:    apperrors.ErrInternalError.Error(),
				Message: "internal server error",
				Status:  http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	result := httpdto.Result[category.Category]{
		Data: *cat,
	}
	response.JSON(w, http.StatusOK, result)
}

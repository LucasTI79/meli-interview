package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	httpdto "github.com/lucasti79/meli-interview/internal/infra/http"
	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/service"
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
// @Summary List all products
// @Description Get a list of all available products with optional filters
// @Tags products
// @Accept  json
// @Produce json
// @Param filters query product.ProductFilter false "Product filters"
// @Success 200 {object} ProductPaginatedResult
// @Success 204 "No content"
// @Failure 400 {object} httpdto.ErrorResponse
// @Failure 500 {object} httpdto.ErrorResponse
// @Router  /api/v1/products [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	filters := product.ProductFilter{
		Name: r.URL.Query().Get("name"),
	}

	if cats := r.URL.Query().Get("categories"); cats != "" {
		filters.Categories = strings.Split(cats, ",")
	}

	if min := r.URL.Query().Get("minPrice"); min != "" {
		filters.MinPrice, _ = strconv.ParseFloat(min, 64)
	}
	if max := r.URL.Query().Get("maxPrice"); max != "" {
		filters.MaxPrice, _ = strconv.ParseFloat(max, 64)
	}

	// paginação default
	filters.Page = 1
	filters.PageSize = 10

	if page := r.URL.Query().Get("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			filters.Page = p
		}
	}
	if size := r.URL.Query().Get("pageSize"); size != "" {
		if s, err := strconv.Atoi(size); err == nil {
			filters.PageSize = s
		}
	}

	if err := h.validator.Struct(filters); err != nil {
		response.JSON(w, http.StatusBadRequest, httpdto.ErrorResponse{
			Code:    apperrors.ErrValidation.Error(),
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	products, total, err := h.service.GetAllWithContext(r.Context(), filters)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, httpdto.ErrorResponse{
			Code:    apperrors.ErrInternalError.Error(),
			Message: "internal server error",
			Status:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	result := httpdto.PaginatedResult[product.Product]{
		Data:       products,
		TotalCount: total,
		Page:       filters.Page,
		PageSize:   filters.PageSize,
	}

	response.JSON(w, http.StatusOK, result)
}

// GetByID godoc
// @Summary Get a product by ID
// @Description Retrieve details of a product by its ID
// @Tags products
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} ProductResult
// @Failure 400 {object} httpdto.ErrorResponse
// @Failure 404 {object} httpdto.ErrorResponse
// @Failure 500 {object} httpdto.ErrorResponse
// @Router /api/v1/products/{productId} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "productId")

	if productId == "" {
		response.JSON(w, http.StatusBadRequest, httpdto.ErrorResponse{
			Code:    product.ErrProductInvalidID,
			Message: "product ID is required",
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	pr, err := h.service.GetByIDWithContext(r.Context(), productId)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrResourceNotExists):
			response.JSON(w, http.StatusNotFound, httpdto.ErrorResponse{
				Code:    product.ErrProductNotFound,
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

	response.JSON(w, http.StatusOK, httpdto.Result[*product.Product]{Data: pr})
}

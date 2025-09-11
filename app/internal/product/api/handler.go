package api

import (
	"errors"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/lucasti79/meli-interview/internal/product"
	_ "github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/service"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
	"github.com/lucasti79/meli-interview/pkg/web/response"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

// GetAll godoc
// @Summary List all products
// @Description Get a list of all available products
// @Tags products
// @Produce json
// @Success 200 {array} product.Product "List of products"
// @Router  /api/v1/products [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllWithContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response.JSON(w, http.StatusOK, products)
}

// GetByID godoc
// @Summary Get a product by ID
// @Description Retrieve details of a product by its ID
// @Tags products
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} product.Product "Product found"
// @Router /api/v1/products/{productId} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "productId")

	if productId == "" {
		response.Error(w, http.StatusBadRequest, "product ID is required", product.ProductInvalidID)
		return
	}

	pr, err := h.service.GetByIDWithContext(r.Context(), productId)
	if err != nil {

		switch {
		case errors.Is(err, apperrors.ErrResourceNotExists):
			response.Error(w, http.StatusNotFound, "product not found", product.ProductNotFound)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error", "")
		}
		return
	}
	response.JSON(w, http.StatusOK, pr)
}

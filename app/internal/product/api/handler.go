package api

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	_ "github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/service"
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
		http.Error(w, "No products found", http.StatusNotFound)
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
	product, err := h.service.GetByIDWithContext(r.Context(), productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	response.JSON(w, http.StatusOK, product)
}

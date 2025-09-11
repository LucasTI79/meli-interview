package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/meli-interview/internal/product/api"
)

func buildProductsRoutes(productHandler *api.Handler) http.Handler {

	r := chi.NewRouter()

	r.Get("/", productHandler.GetAll)
	r.Get("/{productId}", productHandler.GetByID)

	return r
}

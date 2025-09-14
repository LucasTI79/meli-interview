package router

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/lucasti79/meli-interview/internal/category/api"
)

func buildCategoriesRoutes(categoryHandler *api.Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", categoryHandler.GetAll) // GET /api/v1/categories
	r.Get("/{categoryName}", categoryHandler.GetByName)
	return r
}

package router

import (
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lucasti79/meli-interview/internal/factory"
)

type router struct {
}

func (router *router) MapRoutes(appFactory *factory.AppFactory) http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),
	)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // origens permitidas
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // valor em segundos (cache do preflight request)
	}))

	r.Mount("/", buildDocsRoutes())

	r.Route("/api/v1", func(rp chi.Router) {
		rp.Route("/products", func(rp chi.Router) {
			rp.Mount("/", buildProductsRoutes(appFactory.ProductHandler))
		})

		rp.Route("/categories", func(rp chi.Router) {
			rp.Mount("/", buildCategoriesRoutes(appFactory.CategoryHandler))
		})
	})

	return r
}

func NewRouter() *router {
	return &router{}
}

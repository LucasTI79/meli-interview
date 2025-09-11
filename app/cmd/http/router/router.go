package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Mount("/", buildDocsRoutes())

	r.Route("/api/v1", func(rp chi.Router) {
		rp.Route("/products", func(rp chi.Router) {
			rp.Mount("/", buildProductsRoutes(appFactory.ProductHandler))
		})
	})

	return r
}

func NewRouter() *router {
	return &router{}
}

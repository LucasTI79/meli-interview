package router

import (
	"fmt"
	"net/http"

	scalargo "github.com/bdpiprava/scalar-go"
	chi "github.com/go-chi/chi/v5"
	_ "github.com/lucasti79/meli-interview/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func buildDocsRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.json"),
	))

	r.Get("/docs", func(w http.ResponseWriter, req *http.Request) {
		html, err := scalargo.NewV2(
			scalargo.WithSpecURL("/swagger/doc.json"),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	return r
}

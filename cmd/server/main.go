package main

import (
	"github.com/alexandre-dos-reis/go-templ-web/templates/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.Logger)

	// routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home().Render(r.Context(), w)
	})

	router.Get("/blog", func(w http.ResponseWriter, r *http.Request) {
		pages.Blog().Render(r.Context(), w)
	})

	router.Get("/parcours", func(w http.ResponseWriter, r *http.Request) {
		pages.Parcours().Render(r.Context(), w)
	})

	router.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		pages.Contact().Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", router)
}

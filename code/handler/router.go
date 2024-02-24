package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	pages := PageController(db)
	r.Get("/", pages.Index)

	return r
}

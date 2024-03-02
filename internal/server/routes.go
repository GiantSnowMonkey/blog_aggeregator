package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(customCors())
	router.Use(middleware.Logger)
	routerV1 := chi.NewRouter()
	router.Mount("/v1", routerV1)
	routerV1.Get("/readiness", handlerReadiness)

	return router
}

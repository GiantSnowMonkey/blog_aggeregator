package server

import (
	"net/http"

	databse "github.com/GiantSnowMonkey/blog_aggeregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiConfig struct {
	DB *databse.Queries
}

func (s Server) RegisterRoutes() http.Handler {
	dbQueries := databse.LoadDB()
	apiCfg := &apiConfig{
		DB: dbQueries,
	}
	router := chi.NewRouter()
	router.Use(customCors())
	router.Use(middleware.Logger)
	routerV1 := chi.NewRouter()
	router.Mount("/v1", routerV1)
	routerV1.Get("/readiness", handlerReadiness)

	return router
}

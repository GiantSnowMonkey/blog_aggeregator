package server

import (
	"net/http"

	database "github.com/GiantSnowMonkey/blog_aggeregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiConfig struct {
	DB *database.Queries
}

func (s Server) RegisterRoutes() http.Handler {
	dbQueries := database.LoadDB()
	apiCfg := &apiConfig{
		DB: dbQueries,
	}
	router := chi.NewRouter()
	routerV1 := chi.NewRouter()

	router.Use(customCors(), middleware.Logger)

	router.Mount("/v1", routerV1)

	routerV1.Get("/readiness", handlerReadiness)
	routerV1.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
	routerV1.Post("/users", apiCfg.handlerUsersCreate)
	routerV1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedsCreate))

	return router
}

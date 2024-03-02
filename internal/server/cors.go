package server

import (
	"github.com/go-chi/cors"
	"net/http"
)

func customCors() func(http.Handler) http.Handler {
	corsOptions :=
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			MaxAge:           300,
		}
	return cors.Handler(corsOptions)
}

package server

import (
	"context"
	"log"
	"net/http"

	auth "github.com/GiantSnowMonkey/blog_aggeregator/internal/auth"
	database "github.com/GiantSnowMonkey/blog_aggeregator/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := cfg.GetUserFromApiKey(r.Header)
		if err != nil {
			return
		}
		handler(w, r, user)
	})
}

func (cfg *apiConfig) GetUserFromApiKey(header http.Header) (database.User, error) {
	apiKey, err := auth.GetApiKeyFromHeader(header)
	if err != nil {
		return database.User{}, err
	}
	user, err := cfg.DB.GetUserByApiKey(context.Background(), pgtype.Text{String: apiKey, Valid: true})
	if err != nil {
		log.Println("Failed to get user with ApiKey:", err)
		return database.User{}, err
	}
	return user, nil
}

package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	database "github.com/GiantSnowMonkey/blog_aggeregator/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Couldn't parse the params:", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	timeNowTimeStamp := pgtype.Timestamp{Time: time.Now(), Valid: true}
	user, err := cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        pgtype.UUID{Bytes: [16]byte(uuid.New()), Valid: true},
		CreatedAt: timeNowTimeStamp,
		UpdatedAt: timeNowTimeStamp,
		Name:      params.Name,
	})
	if err != nil {
		log.Println("Couldn't create user:", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, database.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	})
}

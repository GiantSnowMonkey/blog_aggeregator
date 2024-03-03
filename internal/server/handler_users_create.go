package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	databse "github.com/GiantSnowMonkey/blog_aggeregator/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RespondUser struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"update_at"`
	Name      string           `json:"name"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Couldn't parse the params:", err)
	}
	timeNowTimeStamp := pgtype.Timestamp{Time: time.Now(), Valid: true}
	userID := uuid.New()
	println("UUID:", userID.String())
	user, err := cfg.DB.CreateUser(context.Background(), databse.CreateUserParams{
		ID:        pgtype.UUID{Bytes: [16]byte(userID), Valid: true},
		CreatedAt: timeNowTimeStamp,
		UpdatedAt: timeNowTimeStamp,
		Name:      params.Name,
	})
	if err != nil {
		log.Println(err)
	}
	respondWithJSON(w, http.StatusOK, RespondUser{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
	})
}

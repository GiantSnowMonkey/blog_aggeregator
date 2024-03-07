package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	database "github.com/GiantSnowMonkey/blog_aggeregator/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Failed to get params from request body")
		respondWithError(w, http.StatusBadRequest, "bad request params")
		return
	}
	timeNowTimeStamp := pgtype.Timestamp{Time: time.Now(), Valid: true}
	response, err := c.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		CreatedAt: timeNowTimeStamp,
		UpdatedAt: timeNowTimeStamp,
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		log.Println("Failed to create feed")
		respondWithError(w, http.StatusInternalServerError, "failed to create feed")
		return
	}
	respondWithJSON(w, http.StatusOK, database.Feed{
		ID:        response.ID,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
		UserID:    response.UserID,
		Url:       response.Url,
		Name:      response.Name,
	})
}

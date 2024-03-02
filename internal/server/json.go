package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responding  with 5XX error:", msg)
	}
	type Response struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, statusCode, Response{Error: msg})
}

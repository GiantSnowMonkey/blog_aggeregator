package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

// ErrNoAuthHeaderIncluded -
var ErrNoAuthHeaderIncluded = errors.New("not auth header included in request")

func GetApiKeyFromHeader(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		log.Println("Authorization header is empty")
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		log.Println("Malformed authorization header")
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

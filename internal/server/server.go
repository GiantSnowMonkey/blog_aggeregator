package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	newServer := &Server{
		port: port,
	}

	server := &http.Server{
		Addr:        fmt.Sprint(":", newServer.port),
		IdleTimeout: time.Minute,
		Handler:     http.NotFoundHandler(),
	}
	return server, err
}

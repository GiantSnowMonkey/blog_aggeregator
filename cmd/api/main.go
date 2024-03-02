package main

import (
	"log"

	"github.com/GiantSnowMonkey/blog_aggeregator/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Running server on port:", server.Addr[1:])
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

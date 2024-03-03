package databse

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func LoadDB() *Queries {
	dbProtocol := os.Getenv("DATABASE_URL")
	println("dbProto:", dbProtocol)
	if dbProtocol == "" {
		log.Fatalln("Could not load db protocol from env file")
	}
	conn, err := pgx.Connect(context.Background(), dbProtocol)
	// defer conn.Close(context.Background())
	if err != nil {
		log.Fatalln("Could not open db:", conn)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalln("Could not open db:", conn)
	}
	return New(conn)
}

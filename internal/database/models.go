// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Feed struct {
	ID        int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Name      string
	Url       string
	UserID    pgtype.UUID
}

type User struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Name      string
	ApiKey    pgtype.Text
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id, name, createdAt, updatedAt)
VALUES ($1, $2, $3, $4)
RETURNING id, name, createdat, updatedat
`

type CreateUserParams struct {
	ID        uuid.UUID
	Name      string
	Createdat time.Time
	Updatedat time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Createdat,
		arg.Updatedat,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

package users_postgres_repository

import "time"

type UserModel struct {
	ID      int
	Version int64

	Name         string
	Email        string
	PasswordHash string
	Role         string
	Plan         *string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

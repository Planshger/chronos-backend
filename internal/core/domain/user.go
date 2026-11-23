package domain

import (
	"fmt"
	"net/mail"
	"time"

	core_errors "github.com/Planshger/chronos-backend/internal/core/errors"
)

type User struct {
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

func NewUser(
	id int,
	version int64,
	name string,
	email string,
	passwordHash string,
	role string,
	plan *string,
	status string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:           id,
		Version:      version,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Plan:         plan,
		Status:       status,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func NewUserUninitialized(
	name string,
	email string,
	passwordHash string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		name,
		email,
		passwordHash,
		"user",
		nil,
		"pending-payment",
		time.Time{},
		time.Time{},
	)
}

func (u *User) Validate() error {
	nameLength := len([]rune(u.Name))

	if nameLength < 3 || nameLength > 100 {
		return fmt.Errorf(
			"invalid Name len: %d: %w",
			nameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	emailLength := len([]rune(u.Email))

	if emailLength == 0 || emailLength > 255 {
		return fmt.Errorf("invalid Email length: %w", core_errors.ErrInvalidArgument)
	}

	email, err := mail.ParseAddress(u.Email)
	if err != nil || email.Address != u.Email {
		return fmt.Errorf("invalid Email format: %w", core_errors.ErrInvalidArgument)
	}

	passwordHashLength := len([]rune(u.PasswordHash))

	if passwordHashLength == 0 || passwordHashLength > 255 {
		return fmt.Errorf("invalid PasswordHash: %w", core_errors.ErrInvalidArgument)
	}

	if u.Role != "" && u.Role != "user" && u.Role != "admin" {
		return fmt.Errorf("invalid Role: %w", core_errors.ErrInvalidArgument)
	}

	if u.Plan != nil {
		planLength := len([]rune(*u.Plan))

		if planLength > 30 {
			return fmt.Errorf("invalid Plan: %w", core_errors.ErrInvalidArgument)
		}
	}

	if u.Status != "active" && u.Status != "blocked" && u.Status != "pending-payment" {
		return fmt.Errorf("invalid Status: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

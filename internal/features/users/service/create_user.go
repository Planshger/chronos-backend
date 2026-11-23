package users_services

import (
	"context"
	"fmt"

	"github.com/Planshger/chronos-backend/internal/core/domain"
	core_errors "github.com/Planshger/chronos-backend/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	command CreateUserCommand,
) (domain.User, error) {
	if command.Password == "" {
		return domain.User{}, fmt.Errorf("password is required: %w", core_errors.ErrInvalidArgument)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user := domain.NewUserUninitialized(
		command.Name,
		command.Email,
		string(passwordHash),
	)

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	user, err = s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

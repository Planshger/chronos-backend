package users_services

import (
	"context"

	"github.com/Planshger/chronos-backend/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type CreateUserCommand struct {
	Name     string
	Email    string
	Password string
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

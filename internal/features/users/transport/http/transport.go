package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Planshger/chronos-backend/internal/core/domain"
	core_http_server "github.com/Planshger/chronos-backend/internal/core/transport/http/server"
	users_services "github.com/Planshger/chronos-backend/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		command users_services.CreateUserCommand,
	) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}

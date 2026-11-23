package users_transport_http

import (
	"net/http"

	"github.com/Planshger/chronos-backend/internal/core/domain"
	core_logger "github.com/Planshger/chronos-backend/internal/core/logger"
	core_http_request "github.com/Planshger/chronos-backend/internal/core/transport/http/request"
	core_http_response "github.com/Planshger/chronos-backend/internal/core/transport/http/response"
	users_services "github.com/Planshger/chronos-backend/internal/features/users/service"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	ID      int     `json:"id"`
	Version int64   `json:"version"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Role    string  `json:"role"`
	Plan    *string `json:"plan"`
	Status  string  `json:"status"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain, err := h.usersService.CreateUser(ctx, domainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) users_services.CreateUserCommand {
	return users_services.CreateUserCommand{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:      user.ID,
		Version: user.Version,
		Name:    user.Name,
		Email:   user.Email,
		Role:    user.Role,
		Plan:    user.Plan,
		Status:  user.Status,
	}
}

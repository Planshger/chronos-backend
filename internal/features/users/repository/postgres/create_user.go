package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Planshger/chronos-backend/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO chronos.users (name, email, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, now(), now())
	RETURNING id, version, name, email, password_hash, role, plan, status, created_at, updated_at
	`

	row := r.pool.QueryRow(ctx, query, user.Name, user.Email, user.PasswordHash)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Name,
		&userModel.Email,
		&userModel.PasswordHash,
		&userModel.Role,
		&userModel.Plan,
		&userModel.Status,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Name,
		userModel.Email,
		userModel.PasswordHash,
		userModel.Role,
		userModel.Plan,
		userModel.Status,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	)

	return userDomain, nil
}

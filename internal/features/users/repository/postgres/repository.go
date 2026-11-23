package users_postgres_repository

import core_postgres_pool "github.com/Planshger/chronos-backend/internal/core/repository/postgres/pool"

type UsersRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersPepository(
	pool core_postgres_pool.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}

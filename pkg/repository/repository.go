package repository

import (
	"github.com/jmoiron/sqlx"
	"go-demo/pkg/repository/postgres"
	"go-demo/types"
)

type AuthorizationRepository interface {
	CreateUser(user types.User) (int, error)
	GetUserByUsername(username string) (types.User, error)
}

type TodoListRepository interface {
}

type TodoItemRepository interface {
}

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	TodoItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: postgres.NewAuthPostgres(
			db),
	}
}

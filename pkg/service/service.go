package service

import (
	"go-demo/pkg/repository"
	"go-demo/types"
)

type AuthorizationService interface {
	CreateUser(user types.User) (int, error)
	SignIn(username string, password string) (string, error)
}

type TodoListService interface {
}

type TodoItemService interface {
}

type Service struct {
	AuthorizationService
	TodoListService
	TodoItemService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repo),
	}
}

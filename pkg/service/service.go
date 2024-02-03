package service

import (
	todo "github.com/delapaska/university-storage"
	"github.com/delapaska/university-storage/pkg/repository"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

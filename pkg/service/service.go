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
type ProjectList interface {
	Create(userId int, project todo.ProjectList) (int, error)
}
type Service struct {
	Authorization
	ProjectList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ProjectList:   NewProjectListService(repos.ProjectList),
	}
}

package repository

import (
	todo "github.com/delapaska/university-storage"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type ProjectList interface {
	Create(userId int, project todo.ProjectList) (int, error)
	GetAll(userId int) ([]todo.ProjectList, error)
	GetById(userId int, listId int) (todo.ProjectList, error)
}
type Repository struct {
	Authorization
	ProjectList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ProjectList:   NewProjectListPostgres(db),
	}
}

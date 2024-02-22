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
	CreateFolder(projectId int, folderName string) (int, error)
	GenerateProjectToken(projectId int) (string, error)
	GetAllFolders(projectId int) ([]string, error)
	ConnectUserToProject(userId, projectId int, token string) error
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

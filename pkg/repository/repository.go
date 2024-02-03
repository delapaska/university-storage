package repository

import (
	todo "github.com/delapaska/university-storage"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

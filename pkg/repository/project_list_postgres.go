package repository

import (
	"fmt"

	todo "github.com/delapaska/university-storage"
	"github.com/jmoiron/sqlx"
)

type ProjectListPostgres struct {
	db *sqlx.DB
}

func NewProjectListPostgres(db *sqlx.DB) *ProjectListPostgres {
	return &ProjectListPostgres{db: db}
}

func (r *ProjectListPostgres) Create(userId int, project todo.ProjectList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, directory) VALUES ($1,$2) RETURNING id", projectListTable)
	fmt.Println(project.Title)
	fmt.Println(project.Directory)
	row := tx.QueryRow(createListQuery, project.Title, project.Directory)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, project_id) VALUES ($1,$2)", usersProjectsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ProjectListPostgres) GetAll(userId int) ([]todo.ProjectList, error) {
	var projects []todo.ProjectList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.directory FROM %s tl INNER JOIN %s ul on tl.id = ul.project_id WHERE ul.user_id = $1", projectListTable, usersProjectsTable)
	err := r.db.Select(&projects, query, userId)
	return projects, err
}

func (r *ProjectListPostgres) GetById(userId int, listId int) (todo.ProjectList, error) {
	var project todo.ProjectList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.directory FROM %s tl INNER JOIN %s ul on tl.id = ul.project_id WHERE ul.user_id = $1 AND ul.project_id = $2", projectListTable, usersProjectsTable)
	err := r.db.Get(&project, query, userId, listId)
	return project, err
}

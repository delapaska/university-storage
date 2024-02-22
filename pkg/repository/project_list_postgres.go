package repository

import (
	"encoding/hex"
	"fmt"
	"math/rand"

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
	createListQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", projectListTable)

	row := tx.QueryRow(createListQuery, project.Title)
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
	query := fmt.Sprintf("SELECT tl.id, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.project_id WHERE ul.user_id = $1", projectListTable, usersProjectsTable)
	err := r.db.Select(&projects, query, userId)
	return projects, err
}

func (r *ProjectListPostgres) GetById(userId int, listId int) (todo.ProjectList, error) {
	var project todo.ProjectList
	query := fmt.Sprintf("SELECT tl.id, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.project_id WHERE ul.user_id = $1 AND ul.project_id = $2", projectListTable, usersProjectsTable)
	err := r.db.Get(&project, query, userId, listId)
	return project, err
}

func (r *ProjectListPostgres) CreateFolder(projectId int, folderName string) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var folderId int
	createFolderQuery := fmt.Sprintf("INSERT INTO %s (project_id, folder_name) VALUES ($1, $2)", projectFolders)
	fmt.Println("Creating...", createFolderQuery)

	_, err = tx.Exec(createFolderQuery, projectId, folderName)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return folderId, tx.Commit()
}

func (r *ProjectListPostgres) GetAllFolders(projectId int) ([]string, error) {
	var folders []string
	rows, err := r.db.Query("SELECT folder_name FROM project_folders WHERE project_id = $1", projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var folder string
		if err := rows.Scan(&folder); err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}
	return folders, nil
}
func (r *ProjectListPostgres) GenerateProjectToken(projectId int) (string, error) {
	// Проверяем наличие истекших токенов и удаляем их
	_, err := r.db.Exec("DELETE FROM project_tokens WHERE created_at < NOW() - INTERVAL '5 minutes'")
	if err != nil {
		return "", err
	}

	// Генерируем новый токен
	token, err := GenerateToken()
	if err != nil {
		return "", err
	}

	// Вставляем новый токен в базу данных
	_, err = r.db.Exec("INSERT INTO project_tokens (project_id, token) VALUES ($1, $2)", projectId, token)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (r *ProjectListPostgres) ConnectUserToProject(userId, projectId int, token string) error {
	// Проверяем существование токена в базе данных
	fmt.Println("ssocjsacam")
	//var projectId int
	//	var createdAt time.Time

	res, _ := r.db.Exec("SELECT project_id, created_at FROM project_tokens WHERE token = $1", token)
	ans, _ := res.RowsAffected()

	// Проверяем, что токен создан не более 15 минут назад
	//	if time.Since(createdAt) > 15*time.Minute {
	//		return errors.New("Токен устарел")
	//	}
	fmt.Println(ans)
	if ans != 0 {
		_, err := r.db.Exec("INSERT INTO users_lists (user_id, project_id) VALUES ($1, $2)", userId, projectId)
		fmt.Println("Info:", userId, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateToken() (string, error) {
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(tokenBytes), nil
}

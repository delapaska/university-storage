package service

import (
	todo "github.com/delapaska/university-storage"
	"github.com/delapaska/university-storage/pkg/repository"
)

type ProjectListService struct {
	repo repository.ProjectList
}

func NewProjectListService(repo repository.ProjectList) *ProjectListService {
	return &ProjectListService{repo: repo}
}

func (s *ProjectListService) Create(userId int, project todo.ProjectList) (int, error) {
	return s.repo.Create(userId, project)
}
func (s *ProjectListService) GetAll(userId int) ([]todo.ProjectList, error) {
	return s.repo.GetAll(userId)
}
func (s *ProjectListService) GetById(userId int, listId int) (todo.ProjectList, error) {
	return s.repo.GetById(userId, listId)
}
func (s *ProjectListService) CreateFolder(projectId int, folderName string) (int, error) {
	return s.repo.CreateFolder(projectId, folderName)
}

func (s *ProjectListService) GenerateProjectToken(projectId int) (string, error) {
	return s.repo.GenerateProjectToken(projectId)
}

func (s *ProjectListService) GetAllFolders(projectId int) ([]string, error) {
	return s.repo.GetAllFolders(projectId)
}
func (s *ProjectListService) ConnectUserToProject(userId, projectId int, token string) error {
	return s.repo.ConnectUserToProject(userId, projectId, token)
}

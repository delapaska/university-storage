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

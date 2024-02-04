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

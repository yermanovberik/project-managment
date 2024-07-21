package service

import (
	"project-managment/internal/app/models"
	"project-managment/internal/app/repository"
)

type ProjectService interface {
	CreateProject(project models.Project) error
	GetProjectById(id int) (models.Project, error)
	GetAllProjects() ([]models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProjectById(id int) error
	GetTasksByProjectId(projectId int) ([]models.Task, error)
	SearchProjectsByTitle(title string) ([]models.Project, error)
	SearchProjectsByManager(userId int) ([]models.Project, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) CreateProject(project models.Project) error {
	return s.repo.CreateProject(project)
}

func (s *projectService) GetProjectById(id int) (models.Project, error) {
	return s.repo.GetProjectById(id)
}

func (s *projectService) GetAllProjects() ([]models.Project, error) {
	return s.repo.GetAllProjects()
}

func (s *projectService) UpdateProject(project models.Project) error {
	return s.repo.UpdateProject(project)
}

func (s *projectService) DeleteProjectById(id int) error {
	return s.repo.DeleteProjectById(id)
}

func (s *projectService) GetTasksByProjectId(projectId int) ([]models.Task, error) {
	return s.repo.GetTasksByProjectId(projectId)
}

func (s *projectService) SearchProjectsByTitle(title string) ([]models.Project, error) {
	return s.repo.SearchProjectsByTitle(title)
}

func (s *projectService) SearchProjectsByManager(userId int) ([]models.Project, error) {
	return s.repo.SearchProjectsByManager(userId)
}

package service

import (
	"project-managment/internal/app/models"
	"project-managment/internal/app/repository"
)

type TaskService interface {
	CreateTask(task models.Task) error
	GetTaskById(id int) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTaskById(id int) error
	SearchTasksByTitle(title string) ([]models.Task, error)
	SearchTasksByStatus(status string) ([]models.Task, error)
	SearchTasksByPriority(priority string) ([]models.Task, error)
	SearchTasksByAssignee(userId int) ([]models.Task, error)
	SearchTasksByProject(projectId int) ([]models.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task models.Task) error {
	return s.repo.CreateTask(task)
}

func (s *taskService) GetTaskById(id int) (models.Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) UpdateTask(task models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *taskService) DeleteTaskById(id int) error {
	return s.repo.DeleteTaskById(id)
}

func (s *taskService) SearchTasksByTitle(title string) ([]models.Task, error) {
	return s.repo.SearchTasksByTitle(title)
}

func (s *taskService) SearchTasksByStatus(status string) ([]models.Task, error) {
	return s.repo.SearchTasksByStatus(status)
}

func (s *taskService) SearchTasksByPriority(priority string) ([]models.Task, error) {
	return s.repo.SearchTasksByPriority(priority)
}

func (s *taskService) SearchTasksByAssignee(userId int) ([]models.Task, error) {
	return s.repo.SearchTasksByAssignee(userId)
}

func (s *taskService) SearchTasksByProject(projectId int) ([]models.Task, error) {
	return s.repo.SearchTasksByProject(projectId)
}

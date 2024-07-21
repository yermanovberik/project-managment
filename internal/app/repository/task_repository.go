package repository

import (
	"errors"
	"gorm.io/gorm"
	"project-managment/internal/app/models"
)

var ErrTaskNotFound = errors.New("Task not found")

type TaskRepository interface {
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

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task models.Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetTaskById(id int) (models.Task, error) {
	var task models.Task
	err := r.db.Find(&task, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, ErrTaskNotFound
		}
		return models.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTask(task models.Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTaskById(id int) error {
	if err := r.db.Delete(&models.Task{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTaskNotFound
		}
		return err
	}
	return nil
}

func (r *taskRepository) SearchTasksByTitle(title string) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) SearchTasksByStatus(status string) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("status = ?", status).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) SearchTasksByPriority(priority string) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("priority = ?", priority).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) SearchTasksByAssignee(userId int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("assignee_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) SearchTasksByProject(projectId int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("project_id = ?", projectId).Find(&tasks).Error
	return tasks, err
}

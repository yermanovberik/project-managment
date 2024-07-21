package repository

import (
	"errors"
	"gorm.io/gorm"
	"project-managment/internal/app/models"
)

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	CreateProject(project models.Project) error
	GetProjectById(id int) (models.Project, error)
	GetAllProjects() ([]models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProjectById(id int) error
	GetTasksByProjectId(projectId int) ([]models.Task, error)
	SearchProjectsByTitle(title string) ([]models.Project, error)
	SearchProjectsByManager(userId int) ([]models.Project, error)
}

type projectRepository struct {
	Db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{Db: db}
}

func (r *projectRepository) CreateProject(project models.Project) error {
	return r.Db.Create(&project).Error
}

func (r *projectRepository) GetProjectById(id int) (models.Project, error) {
	var project models.Project
	err := r.Db.First(&project, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Project{}, ErrProjectNotFound
		}
		return models.Project{}, err
	}
	return project, nil
}

func (r *projectRepository) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	err := r.Db.Find(&projects).Error
	return projects, err
}

func (r *projectRepository) UpdateProject(project models.Project) error {
	return r.Db.Save(&project).Error
}

func (r *projectRepository) DeleteProjectById(id int) error {
	if err := r.Db.Delete(&models.Project{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProjectNotFound
		}
		return err
	}
	return nil
}

func (r *projectRepository) GetTasksByProjectId(projectId int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.Db.Where("project_id = ?", projectId).Find(&tasks).Error
	return tasks, err
}

func (r *projectRepository) SearchProjectsByTitle(title string) ([]models.Project, error) {
	var projects []models.Project
	err := r.Db.Where("title LIKE ?", "%"+title+"%").Find(&projects).Error
	return projects, err
}

func (r *projectRepository) SearchProjectsByManager(userId int) ([]models.Project, error) {
	var projects []models.Project
	err := r.Db.Where("manager_id = ?", userId).Find(&projects).Error
	return projects, err
}

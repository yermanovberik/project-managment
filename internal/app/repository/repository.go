package repository

import (
	"gorm.io/gorm"
	"project-managment/internal/app/dto"
	"project-managment/internal/app/models"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	err := r.Db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) CreateUser(user models.User) error {

	err := r.Db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetById(id int) (models.User, error) {
	var user models.User
	err := r.Db.Find(&user, id).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *Repository) DeleteUserById(id string) error {
	err := r.Db.Delete(&models.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByName(name string) (models.User, error) {
	var user models.User
	err := r.Db.Find(&user, "name = ?", name).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.Db.Find(&user, "email = ?", email).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) GetAllProjects() (models.Project, error) {
	var project models.Project

	err := r.Db.Find(&project).Error

	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

func (r *Repository) CreateProject(project dto.ProjectRequestDTO) error {
	err := r.Db.Create(&project).Error

	if err != nil {
		return err
	}

	return nil
}

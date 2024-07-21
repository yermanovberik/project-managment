package repository

import (
	"errors"
	"gorm.io/gorm"
	"project-managment/internal/app/models"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserById(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetById(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByName(name string) (models.User, error)
	DeleteUserById(id string) error
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{Db: db}
}
func (r *userRepository) CreateUser(user models.User) error {
	return r.Db.Create(&user).Error
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.Db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.Db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, ErrUserNotFound
	}
	return user, err
}

func (r *userRepository) UpdateUser(user models.User) error {
	return r.Db.Save(&user).Error
}

func (r *userRepository) DeleteUserById(id int) error {
	err := r.Db.Delete(&models.User{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}
	return err
}

func (r *userRepository) GetUserTasks(userId int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.Db.Where("assignee_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (r *userRepository) SearchUsersByName(name string) ([]models.User, error) {
	var users []models.User
	err := r.Db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	return users, err
}

func (r *userRepository) SearchUsersByEmail(email string) ([]models.User, error) {
	var users []models.User
	err := r.Db.Where("email LIKE ?", "%"+email+"%").Find(&users).Error
	return users, err
}

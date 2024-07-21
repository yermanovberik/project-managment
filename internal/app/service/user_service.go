package service

import (
	"project-managment/internal/app/models"
	"project-managment/internal/app/repository"
)

type UserService interface {
	CreateUser(user models.User) error
	GetUserById(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetById(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByName(name string) (models.User, error)
	DeleteUserById(id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	err := s.repo.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById(id int) (models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) DeleteUserById(id int) error {
	return s.repo.DeleteUserById(id)
}

func (s *userService) GetUserTasks(userId int) ([]models.Task, error) {
	return s.repo.GetUserTasks(userId)
}

func (s *userService) SearchUsersByName(name string) ([]models.User, error) {
	return s.repo.SearchUsersByName(name)
}

func (s *userService) SearchUsersByEmail(email string) ([]models.User, error) {
	return s.repo.SearchUsersByEmail(email)
}

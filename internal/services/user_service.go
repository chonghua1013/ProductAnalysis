package services

import (
	"project-name/internal/models"
	"project-name/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id string, user *models.User) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepo.Create(user)
}

func (s *userService) UpdateUser(id string, user *models.User) (*models.User, error) {
	return s.userRepo.Update(id, user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

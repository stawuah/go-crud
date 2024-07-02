package service

import (
	"github/stawuah/go-crud/internal/modules/users/models"
	"github/stawuah/go-crud/internal/modules/users/repository"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
	RegisterUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

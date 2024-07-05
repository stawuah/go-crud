package service

import (
	"github/stawuah/go-crud/internal/modules/users/models"
	"github/stawuah/go-crud/internal/modules/users/repository"
)

// UserService defines the interface for user-related operations
type UserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUser(id int) (*models.User, error)
	RegisterUser(user *models.User) error
}

// userService is the concrete implementation of UserService
type userService struct {
	repo repository.UserRepository1 // Ensure this matches the repository interface name
}

// NewUserService creates a new instance of userService
func NewUserService(r repository.UserRepository1) UserService {
	return &userService{
		repo: r,
	}
}

// GetUser retrieves a user by their ID
func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

// RegisterUser registers a new user
func (s *userService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAllUsers()
}

package repository

import (
	"github/stawuah/go-crud/internal/modules/users/models"
	"github/stawuah/go-crud/internal/modules/users/queries"
)

type UserRepository1 interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
}

type userRepository struct {
	queries *queries.UserQueries
}

func NewUserRepository(q *queries.UserQueries) UserRepository1 {
	return &userRepository{
		queries: q,
	}
}
func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	return r.queries.GetUserByID(id)
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.queries.CreateUser(user)
}

func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	return r.queries.GetAllUsers()
}

package repository

import (
	"github/stawuah/go-crud/internal/modules/users/models"
	"github/stawuah/go-crud/internal/modules/users/queries"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	queries *queries.UserQueries
}

func NewUserRepository(q *queries.UserQueries) UserRepository {
	return &userRepository{q}
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	return r.queries.GetUserByID(id)
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.queries.CreateUser(user)
}

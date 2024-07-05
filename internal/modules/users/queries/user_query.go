package queries

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github/stawuah/go-crud/internal/modules/users/models"
)

type UserQueries struct {
	db *sqlx.DB
}

func NewUserQueries(db *sqlx.DB) *UserQueries {
	return &UserQueries{db}
}

func (q *UserQueries) CreateUser(user *models.User) error {
	_, err := q.db.NamedExec(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`, user)
	return err
}

func (q *UserQueries) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := q.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (q *UserQueries) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := q.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

package repository

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"user_service/internal/domain/user"
	"user_service/internal/repository/interfaces"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) CreateUser() (int, error) {
	return http.StatusOK, nil
}

func (u *UserRepository) UpdateUser() (int, error) {
	return http.StatusNotImplemented, nil
}

func (u *UserRepository) DeleteUser(id string) (int, error) {
	return http.StatusOK, nil
}

func (u *UserRepository) SearchUser(queryType, query string) ([]user.Entity, error) {
	return make([]user.Entity, 0), nil
}

func (u *UserRepository) GetUsers() ([]user.Entity, error) {
	res, err := u.db.Query("SELECT * FROM ")
}
func (u *UserRepository) GetUserById(id string) (user.Entity, error) {
	return user.Entity{}, nil
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}

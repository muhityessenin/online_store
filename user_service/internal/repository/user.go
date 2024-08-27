package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"user_service/internal/domain/user"
	"user_service/internal/repository/interfaces"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) CreateUser(input *user.InputResponse) (int, error) {
	_, err := u.db.Exec("INSERT INTO users (name, email, address, registration_date, role) VALUES ($1, $2, $3, $4, $5)",
		input.Name, input.Email, input.Address, input.RegistrationDate, input.Role)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (u *UserRepository) UpdateUser(id string, input *user.InputResponse) (int, error) {
	_, err := u.db.Exec("UPDATE users SET name = $1, email = $2, address = $3, registration_date = $4, role = $5 WHERE id = $5",
		input.Name, input.Email, input.Address, input.RegistrationDate, input.Role, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (u *UserRepository) DeleteUser(id string) (int, error) {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (u *UserRepository) SearchUser(queryType, query string) ([]user.Entity, error) {
	var users []user.Entity
	var err error
	var q string
	if queryType == "name" {
		q = fmt.Sprintf("SELECT * FROM users WHERE name = $1", query)
	} else if queryType == "email" {
		q = fmt.Sprintf("SELECT * FROM users WHERE email = $1", query)
	}
	err = u.db.Select(&users, q)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) GetUsers() ([]user.Entity, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []user.Entity
	for rows.Next() {
		var user user.Entity
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Address, &user.RegistrationDate, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (u *UserRepository) GetUserById(id string) (user.Entity, error) {
	var user user.Entity
	row, err := u.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return user, err
	}
	defer row.Close()
	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Address, &user.RegistrationDate, &user.Role); err != nil {
			return user, err
		}
	}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}

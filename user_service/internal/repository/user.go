package repository

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"user_service/internal/domain/user"
	"user_service/internal/repository/interfaces"
)

const (
	usersTable    = "users"
	productsTable = "products"
	ordersTable   = "orders"
	paymentsTable = "payments"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) CreateUser() (int, error) {
	_, err := u.db.Exec("INSERT INTO %s (name, email, address, registration_date, role) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
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
	rows, err := u.db.Query("SELECT * FROM %s", usersTable)
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
	return user.Entity{}, nil
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}

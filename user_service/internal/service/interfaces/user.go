package interfaces

import "user_service/internal/domain/user"

type UserService interface {
	GetUsers() ([]user.Entity, error)
	CreateUser() (int, error)
	GetUserById(id string) (user.Entity, error)
	UpdateUser() (int, error)
	DeleteUser(id string) (int, error)
	SearchUser(queryType, query string) ([]user.Entity, error)
}

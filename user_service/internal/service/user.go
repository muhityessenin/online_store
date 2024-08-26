package service

import (
	"user_service/internal/domain/user"
	interf "user_service/internal/repository/interfaces"
	services "user_service/internal/service/interfaces"
)

type UserService struct {
	UserRepository interf.UserRepository
}

func NewUserService(userRepository interf.UserRepository) services.UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) GetUsers() ([]user.Entity, error) {
	return u.UserRepository.GetUsers()
}

func (u *UserService) CreateUser() (int, error) {
	return u.UserRepository.CreateUser()
}

func (u *UserService) GetUserById(id string) (user.Entity, error) {
	return u.UserRepository.GetUserById(id)
}

func (u *UserService) UpdateUser() (int, error) {
	return u.UserRepository.UpdateUser()
}

func (u *UserService) DeleteUser(id string) (int, error) {
	return u.UserRepository.DeleteUser(id)
}

func (u *UserService) SearchUser(queryType, query string) ([]user.Entity, error) {
	return u.UserRepository.SearchUser(queryType, query)
}

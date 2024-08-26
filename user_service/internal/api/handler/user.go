package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_service/internal/service/interfaces"
)

type UserHandler struct {
	UserService interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	res, _ := u.UserService.GetUsers()
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) CreateUser(c *gin.Context)  {}
func (u *UserHandler) GetUserById(c *gin.Context) {}
func (u *UserHandler) UpdateUser(c *gin.Context)  {}
func (u *UserHandler) DeleteUser(c *gin.Context)  {}
func (u *UserHandler) SearchUser(c *gin.Context)  {}

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_service/internal/domain/user"
	"user_service/internal/service/interfaces"
	"user_service/pkg"
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
	var response pkg.Response
	res, err := u.UserService.GetUsers()
	if err != nil {
		response = pkg.Response{
			Status:  http.StatusServiceUnavailable,
			Message: err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response = pkg.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) CreateUser(c *gin.Context) {
	var input user.InputResponse
	var res pkg.Response
	if err := c.ShouldBindJSON(&input); err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := u.UserService.CreateUser(&input)
	if err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input2"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = pkg.Response{Status: http.StatusOK, Data: "", Message: "success"}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) GetUserById(c *gin.Context) {
	var response pkg.Response
	res, err := u.UserService.GetUserById(c.Param("id"))
	if err != nil {
		response = pkg.Response{Status: http.StatusServiceUnavailable, Data: "", Message: "invalid input"}
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response = pkg.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	c.JSON(http.StatusOK, response)
}
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var res pkg.Response
	var input user.InputResponse
	if err := c.ShouldBindJSON(&input); err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := u.UserService.UpdateUser(c.Param("id"), &input)
	if err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = pkg.Response{Status: http.StatusOK, Data: "", Message: "success"}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) DeleteUser(c *gin.Context) {
	var res pkg.Response
	_, err := u.UserService.DeleteUser(c.Param("id"))
	if err != nil {
		res = pkg.Response{Status: http.StatusServiceUnavailable, Data: "", Message: "invalid input"}
		c.JSON(http.StatusServiceUnavailable, res)
	}
	res = pkg.Response{Status: http.StatusOK, Data: "", Message: "success"}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) SearchUser(c *gin.Context) {
	var response pkg.Response
	var res []user.Entity
	var err error
	name := c.Query("name")
	email := c.Query("email")
	if name != "" {
		res, err = u.UserService.SearchUser("name", c.Query("name"))
	} else if email != "" {
		res, err = u.UserService.SearchUser("email", c.Query("email"))
	}
	if err != nil {
		response = pkg.Response{Status: http.StatusServiceUnavailable, Data: "", Message: "invalid input"}
		c.JSON(http.StatusServiceUnavailable, response)
	}
	response = pkg.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	c.JSON(http.StatusOK, response)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	router.GET("/", userHandler.GetUsers)
	router.POST("/", userHandler.CreateUser)
	router.PUT("/:id", userHandler.UpdateUser)
	router.DELETE("/:id", userHandler.DeleteUser)
	router.GET("/:id", userHandler.GetUserById)
	router.GET("/search", userHandler.SearchUser)
}

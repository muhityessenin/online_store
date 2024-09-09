package routes

import (
	"github.com/gin-gonic/gin"
	"product_service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, productHandler *handler.ProductHandler) {
	router.GET("/", productHandler.GetProducts)
	router.POST("/", productHandler.CreateProduct)
	router.PUT("/:id", productHandler.UpdateProduct)
	router.DELETE("/:id", productHandler.DeleteProduct)
	router.GET("/:id", productHandler.GetProductById)
	router.GET("/search", productHandler.SearchUser)
}

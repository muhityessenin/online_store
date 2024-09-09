package handler

import (
	_ "database/sql"
	_ "errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "net/http"
	"product_service/internal/service/interfaces"
)

type ProductHandler struct {
	ProductService interfaces.ProductService
}

func NewProductHandler(service interfaces.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}
func (handler *ProductHandler) GetProducts(ctx *gin.Context)    {}
func (handler *ProductHandler) GetProductById(ctx *gin.Context) {}
func (handler *ProductHandler) CreateProduct(ctx *gin.Context)  {}
func (handler *ProductHandler) UpdateProduct(ctx *gin.Context)  {}
func (handler *ProductHandler) DeleteProduct(ctx *gin.Context)  {}
func (handler *ProductHandler) SearchUser(context *gin.Context) {}

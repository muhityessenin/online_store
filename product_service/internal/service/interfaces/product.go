package interfaces

import "product_service/internal/domain/product"

type ProductService interface {
	GetProducts() ([]product.Entity, error)
	CreateProduct(input *product.InputResponse) (int, error)
	GetProductById(id string) (product.Entity, error)
	UpdateProduct(id string, input *product.InputResponse) (int, error)
	DeleteProduct(id string) (int, error)
	SearchProduct(queryType, query string) ([]product.Entity, error)
}

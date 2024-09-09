package service

import (
	"product_service/internal/domain/product"
	interf "product_service/internal/repository/interfaces"
	services "product_service/internal/service/interfaces"
)

type ProductService struct {
	ProductRepository interf.ProductRepository
}

func (p ProductService) GetProducts() ([]product.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) CreateProduct(input *product.InputResponse) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) GetProductById(id string) (product.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) UpdateProduct(id string, input *product.InputResponse) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) DeleteProduct(id string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) SearchProduct(queryType, query string) ([]product.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductService(productRepository interf.ProductRepository) services.ProductService {
	return &ProductService{ProductRepository: productRepository}
}

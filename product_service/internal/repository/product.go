package repository

import (
	"github.com/jmoiron/sqlx"
	"product_service/internal/domain/product"
	"product_service/internal/repository/interfaces"
)

type ProductRepository struct {
	db *sqlx.DB
}

func (p ProductRepository) GetProducts() ([]product.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) CreateProduct(input *product.InputResponse) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetProductById(id string) (product.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateProduct(id string, input *product.InputResponse) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) DeleteProduct(id string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) SearchProduct(queryType, query string) ([]product.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db *sqlx.DB) interfaces.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

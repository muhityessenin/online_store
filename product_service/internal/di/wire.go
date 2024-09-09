//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "product_service/internal/api"
	"product_service/internal/api/handler"
	"product_service/internal/config"
	"product_service/internal/db"
	"product_service/internal/repository"
	"product_service/internal/service"
)

func Initialize(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewProductHandler,
		repository.NewProductRepository,
		service.NewProductService,
		http.NewServer,
	)
	return &http.Server{}, nil
}

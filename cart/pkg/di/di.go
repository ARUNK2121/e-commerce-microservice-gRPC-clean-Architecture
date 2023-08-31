package di

import (
	server "cart/pkg/api"
	"cart/pkg/api/service"
	client "cart/pkg/client"
	"cart/pkg/config"
	"cart/pkg/db"
	"cart/pkg/repository"
	"cart/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	adminUseCase := usecase.NewCartUseCase(cartRepository, productClient)

	adminServiceServer := service.NewCartServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}

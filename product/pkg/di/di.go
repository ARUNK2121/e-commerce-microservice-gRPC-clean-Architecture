package di

import (
	server "product/pkg/api"
	"product/pkg/api/service"
	"product/pkg/config"
	"product/pkg/db"
	"product/pkg/repository"
	"product/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewProductRepository(gormDB)
	adminUseCase := usecase.NewProductUseCase(adminRepository)

	adminServiceServer := service.NewProductServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}

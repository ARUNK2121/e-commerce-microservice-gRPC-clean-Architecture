package di

import (
	server "users/pkg/api"
	"users/pkg/api/service"
	"users/pkg/config"
	"users/pkg/db"
	"users/pkg/repository"
	"users/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewUserUseCase(adminRepository)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}

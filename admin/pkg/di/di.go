package di

import (
	server "admin/pkg/api"
	"admin/pkg/api/service"
	"admin/pkg/config"
	"admin/pkg/db"
	"admin/pkg/repository"
	"admin/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)

	adminServiceServer := service.NewAdminServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}

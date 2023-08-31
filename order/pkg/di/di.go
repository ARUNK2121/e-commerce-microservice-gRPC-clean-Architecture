package di

import (
	server "order/pkg/api"
	"order/pkg/api/service"
	"order/pkg/client"
	"order/pkg/config"
	"order/pkg/db"
	"order/pkg/repository"
	"order/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient := client.NewCartClient(cfg)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient)

	orderServiceServer := service.NewOrderServer(orderUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, orderServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}

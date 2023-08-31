package di

import (
	server "api_gateway/pkg/api"
	"api_gateway/pkg/api/handler"
	"api_gateway/pkg/client"
	"api_gateway/pkg/config"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	productClient := client.NewProductClient(cfg)
	productHandler := handler.NewProductHandler(productClient)

	userClient := client.NewUserClient(cfg)
	userHandler := handler.NewUserHandler(userClient)

	cartClient := client.NewCartClient(cfg)
	cartHandler := handler.NewCartHandler(cartClient)

	orderClient := client.NewOrderClient(cfg)
	orderHandler := handler.NewOrderHandler(orderClient)

	serverHTTP := server.NewServerHTTP(adminHandler, productHandler, userHandler, cartHandler, orderHandler)

	return serverHTTP, nil
}

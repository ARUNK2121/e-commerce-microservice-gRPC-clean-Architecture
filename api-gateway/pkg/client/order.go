package client

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/config"
	pb "api_gateway/pkg/proto/order"
	"api_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type orderClient struct {
	Client pb.OrderClient
}

func NewOrderClient(cfg config.Config) interfaces.OrderClient {

	grpcConnection, err := grpc.Dial(cfg.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewOrderClient(grpcConnection)

	return &orderClient{
		Client: grpcClient,
	}

}

func (a *orderClient) MakeOrder(cart_id, address_id int) (int, error) {

	res, err := a.Client.MakeOrder(context.Background(), &pb.MakeOrderRequest{
		UserId:    int64(cart_id),
		AddressId: int64(address_id),
	})
	if err != nil {
		return 500, err
	}

	return int(res.Status), nil

}

func (a *orderClient) GetOrders(user_id int) ([]models.GetOrders, error) {

	response, err := a.Client.GetOrders(context.Background(), &pb.GetOrderRequest{
		UserId: int64(user_id),
	})
	if err != nil {
		return []models.GetOrders{}, err
	}

	var result []models.GetOrders
	for _, v := range response.Details {
		result = append(result, models.GetOrders{
			ID:          int(v.OrderId),
			FinalPrice:  float64(v.Amount),
			OrderStatus: v.OrderStatus,
		})
	}

	return result, nil

}

func (a *orderClient) EditOrderStatus(req models.EditStatus) (string, error) {

	response, err := a.Client.EditOrderStatus(context.Background(), &pb.EditStatusRequest{
		OrderId: int64(req.OrderID),
		Status:  req.Status,
	})
	if err != nil {
		return "error", err
	}

	return response.Response, nil

}

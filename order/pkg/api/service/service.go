package service

import (
	"context"
	"fmt"
	pb "order/pkg/proto/order"
	interfaces "order/pkg/usecase/interface"
)

type OrderServer struct {
	UseCase interfaces.OrderUseCase
	pb.UnimplementedOrderServer
}

func NewOrderServer(useCase interfaces.OrderUseCase) pb.OrderServer {

	return &OrderServer{
		UseCase: useCase,
	}

}

func (o *OrderServer) MakeOrder(ctx context.Context, req *pb.MakeOrderRequest) (*pb.MakeOrderResponse, error) {

	fmt.Println("reaches service")

	cart_id := int(req.UserId)
	address_id := int(req.AddressId)
	err := o.UseCase.MakeOrder(cart_id, address_id)
	if err != nil {
		return &pb.MakeOrderResponse{
			Status: 500,
		}, err
	}

	return &pb.MakeOrderResponse{
		Status: 201,
	}, nil

}

func (o *OrderServer) GetOrders(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {

	fmt.Println("reaches service")

	user_id := int(req.UserId)
	details, err := o.UseCase.GetOrders(user_id)
	if err != nil {
		return &pb.GetOrderResponse{}, err
	}

	var result pb.GetOrderResponse

	for _, v := range details {
		result.Details = append(result.Details, &pb.OrderDetail{
			OrderId:     int64(v.ID),
			Amount:      float32(v.FinalPrice),
			OrderStatus: v.OrderStatus,
		})
	}

	return &result, nil

}

func (o *OrderServer) EditOrderStatus(ctx context.Context, req *pb.EditStatusRequest) (*pb.EditStatusResponse, error) {

	fmt.Println("reaches service")

	order_id := int(req.OrderId)
	status := req.Status
	err := o.UseCase.EditOrderStatus(order_id, status)
	if err != nil {
		return &pb.EditStatusResponse{}, err
	}

	return &pb.EditStatusResponse{
		Response: "successfull edited the order status",
	}, nil

}

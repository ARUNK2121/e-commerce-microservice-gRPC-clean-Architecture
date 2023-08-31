package service

import (
	pb "cart/pkg/proto/cart"
	interfaces "cart/pkg/usecase/interface"
	"cart/pkg/utils/models"
	"context"
	"fmt"
)

type CartServer struct {
	CartUseCase interfaces.CartUseCase
	pb.UnimplementedCartServer
}

func NewCartServer(useCase interfaces.CartUseCase) pb.CartServer {

	return &CartServer{
		CartUseCase: useCase,
	}

}

func (a *CartServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {

	fmt.Println("reaches service")

	details := models.AddToCart{
		UserID:      int(req.UserId),
		InventoryID: int(req.ProductId),
	}

	err := a.CartUseCase.AddToCart(details)
	if err != nil {
		return &pb.AddToCartResponse{
			Status: 500,
			Error:  err.Error(),
		}, err
	}

	return &pb.AddToCartResponse{
		Status: 201,
		Error:  "",
	}, nil

}

func (a *CartServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {

	user_id := int(req.Id)

	cart, err := a.CartUseCase.GetCart(user_id)
	if err != nil {
		return &pb.GetCartResponse{
			Items: []*pb.GetCart{},
		}, err
	}

	var result pb.GetCartResponse
	for _, v := range cart {
		result.Items = append(result.Items, &pb.GetCart{
			ProductName: v.ProductName,
			CategoryId:  int64(v.Category_id),
			Quantity:    int64(v.Quantity),
			Total:       float32(v.Total),
		})
	}

	return &result, nil

}

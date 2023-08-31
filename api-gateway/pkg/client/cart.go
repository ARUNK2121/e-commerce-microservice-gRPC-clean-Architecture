package client

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/config"
	pb "api_gateway/pkg/proto/cart"
	"api_gateway/pkg/utils/models"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type cartClient struct {
	Client pb.CartClient
}

func NewCartClient(cfg config.Config) interfaces.CartClient {

	grpcConnection, err := grpc.Dial(cfg.CartSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCartClient(grpcConnection)

	return &cartClient{
		Client: grpcClient,
	}

}

func (a *cartClient) AddToCart(add models.AddToCart) error {

	fmt.Println("reaches client")

	res, err := a.Client.AddToCart(context.Background(), &pb.AddToCartRequest{
		UserId:    int64(add.UserID),
		ProductId: int64(add.InventoryID),
	})
	if err != nil {
		return err
	}

	if res.Error != "" {
		return errors.New(res.Error)
	}

	return nil

}

func (p *cartClient) GetCart(id int) ([]models.GetCart, error) {
	res, err := p.Client.GetCart(context.Background(), &pb.GetCartRequest{
		Id: int64(id),
	})

	if err != nil {
		return []models.GetCart{}, err
	}

	var products []models.GetCart
	for _, p := range res.Items {
		products = append(products, models.GetCart{
			ProductName:     p.ProductName,
			Category_id:     int(p.CategoryId),
			Quantity:        int(p.Quantity),
			Total:           float64(p.Total),
			DiscountedPrice: float64(p.Total),
		})
	}

	return products, nil
}

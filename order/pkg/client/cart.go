package client

import (
	"context"
	"fmt"
	clienterface "order/pkg/client/interface"
	"order/pkg/config"
	pb "order/pkg/proto/cart"
	"order/pkg/utils/models"

	"google.golang.org/grpc"
)

type cartClient struct {
	Client pb.CartClient
}

func NewCartClient(cfg config.Config) clienterface.CartClient {

	grpcConnection, err := grpc.Dial(cfg.CartSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCartClient(grpcConnection)

	return &cartClient{
		Client: grpcClient,
	}

}

func (c *cartClient) GetCart(user_id int) ([]models.GetCart, error) {
	fmt.Println("reached just before rpc call")
	res, err := c.Client.GetCart(context.Background(), &pb.GetCartRequest{
		Id: int64(user_id),
	})
	if err != nil {
		return []models.GetCart{}, err
	}

	var result []models.GetCart
	for _, v := range res.Items {
		result = append(result, models.GetCart{
			ProductName:     v.ProductName,
			Category_id:     int(v.CategoryId),
			Quantity:        int(v.Quantity),
			Total:           float64(v.Total),
			DiscountedPrice: float64(v.Total),
		})
	}

	fmt.Println(result)

	return result, nil
}

package client

import (
	"cart/pkg/config"
	pb "cart/pkg/proto/product"
	"cart/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type client struct {
	Client pb.ProductClient
}

func NewProductClient(c *config.Config) *client {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewProductClient(cc)

	return &client{
		Client: pbClient,
	}
}

func (c *client) GetStockOfProducts(id int) (int, error) {
	res, err := c.Client.GetStockOfProducts(context.Background(), &pb.GetStockRequest{ProductID: int64(id)})
	if err != nil {
		return 0, err
	}

	stock := int(res.Stock)

	return stock, nil
}

func (c *client) GetProductsDetails(products []int) ([]models.GetCart, error) {
	var req pb.GetDetailsRequest
	for _, v := range products {
		req.Products = append(req.Products, int64(v))
	}
	res, err := c.Client.GetProductsDetails(context.Background(), &req)
	if err != nil {
		return []models.GetCart{}, err
	}

	var getcart []models.GetCart
	for _, v := range res.Details {
		getcart = append(getcart, models.GetCart{
			ProductName:     v.ProductName,
			Category_id:     int(v.CategoryId),
			Quantity:        int(v.Quantity),
			Total:           float64(v.Total),
			DiscountedPrice: float64(v.Total),
		})
	}

	return getcart, nil
}

package client

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/config"
	pb "api_gateway/pkg/proto/product"
	"api_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type productClient struct {
	Client pb.ProductClient
}

func NewProductClient(cfg config.Config) interfaces.ProductClient {

	grpcConnection, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductClient(grpcConnection)

	return &productClient{
		Client: grpcClient,
	}

}

func (p *productClient) AddProduct(model models.Inventories) (int, error) {
	res, err := p.Client.AddProduct(context.Background(), &pb.AddProductRequest{
		ID:          int64(model.ID),
		CategoryID:  int64(model.CategoryID),
		ProductName: model.ProductName,
		Stock:       int64(model.Stock),
		Size:        model.Size,
		Price:       float32(model.Price),
	})
	if err != nil {
		return 500, nil
	}

	return int(res.Status), nil
}

func (p *productClient) ListProducts(page int) ([]models.Inventories, error) {
	res, err := p.Client.ListProducts(context.Background(), &pb.ListProductRequest{
		Page: int64(page),
	})

	if err != nil {
		return []models.Inventories{}, err
	}

	var products []models.Inventories
	for _, p := range res.LisrtProducts {
		products = append(products, models.Inventories{
			ID:          uint(p.ID),
			CategoryID:  int(p.CategoryID),
			ProductName: p.ProductName,
			Stock:       int(p.Stock),
			Size:        p.Size,
			Price:       float64(p.Price),
		})
	}

	return products, nil
}

package service

import (
	"context"
	"fmt"
	pb "product/pkg/proto"
	interfaces "product/pkg/usecase/interface"
	"product/pkg/utils/models"
)

type ProductServer struct {
	productUseCase interfaces.ProductUseCase
	pb.UnimplementedProductServer
}

func NewProductServer(useCase interfaces.ProductUseCase) pb.ProductServer {

	return &ProductServer{
		productUseCase: useCase,
	}

}

func (p *ProductServer) AddProduct(ctx context.Context, details *pb.AddProductRequest) (*pb.AddProductResponse, error) {

	fmt.Println("reached service")

	model := models.Inventories{
		ID:          uint(details.ID),
		CategoryID:  int(details.CategoryID),
		ProductName: details.ProductName,
		Size:        details.Size,
		Stock:       int(details.Stock),
		Price:       float64(details.Price),
	}

	status, err := p.productUseCase.AddProduct(model)
	if err != nil {
		return &pb.AddProductResponse{
			Status: int64(status),
		}, err
	}

	return &pb.AddProductResponse{
		Status: int64(status),
	}, nil

}

func (p *ProductServer) ListProducts(ctx context.Context, details *pb.ListProductRequest) (*pb.ListProductResponse, error) {

	fmt.Println("reached service")

	page := int(details.Page)

	products, err := p.productUseCase.ListProducts(page)
	if err != nil {
		return &pb.ListProductResponse{
			LisrtProducts: []*pb.ProductDetails{},
		}, err
	}

	var result pb.ListProductResponse
	for _, v := range products {
		result.LisrtProducts = append(result.LisrtProducts, &pb.ProductDetails{
			ID:          int64(v.ID),
			CategoryID:  int64(v.CategoryID),
			ProductName: v.ProductName,
			Size:        v.Size,
			Stock:       int64(v.Stock),
			Price:       float32(v.Price),
		})
	}

	return &result, nil

}

func (p *ProductServer) GetStockOfProducts(ctx context.Context, details *pb.GetStockRequest) (*pb.GetStockResponse, error) {

	stock, err := p.productUseCase.GetStockOfProducts(int(details.ProductID))
	if err != nil {
		return &pb.GetStockResponse{
			Stock: int64(stock),
		}, err
	}

	return &pb.GetStockResponse{
		Stock: int64(stock),
	}, err

}

func (p *ProductServer) GetProductsDetails(ctx context.Context, req *pb.GetDetailsRequest) (*pb.GetDetailsResponse, error) {

	fmt.Println("reached service")

	products, err := p.productUseCase.GetProductsDetails(req.Products)
	if err != nil {
		return &pb.GetDetailsResponse{
			Details: []*pb.CartProducts{},
		}, err
	}

	var result pb.GetDetailsResponse
	for _, v := range products {
		result.Details = append(result.Details, &pb.CartProducts{
			ProductName: v.ProductName,
			CategoryId:  int64(v.Category_id),
			Quantity:    1,
			Total:       float32(v.Total),
		})
	}

	return &result, nil

}

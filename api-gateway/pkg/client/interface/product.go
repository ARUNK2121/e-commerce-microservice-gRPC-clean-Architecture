package interfaces

import "api_gateway/pkg/utils/models"

type ProductClient interface {
	AddProduct(add models.Inventories) (int, error)
	ListProducts(int) ([]models.Inventories, error)
}

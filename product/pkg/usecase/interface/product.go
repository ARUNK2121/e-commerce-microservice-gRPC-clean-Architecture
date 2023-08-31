package interfaces

import "product/pkg/utils/models"

type ProductUseCase interface {
	AddProduct(inventory models.Inventories) (int, error)
	ListProducts(page int) ([]models.Inventories, error)
	GetStockOfProducts(id int) (int, error)
	GetProductsDetails([]int64) ([]models.GetCart, error)
}

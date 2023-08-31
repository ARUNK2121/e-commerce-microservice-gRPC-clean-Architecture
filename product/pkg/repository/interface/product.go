package interfaces

import "product/pkg/utils/models"

type ProductRepository interface {
	AddProduct(inventory models.Inventories) (int, error)
	ListProducts(page int) ([]models.Inventories, error)
	GetStockOfProducts(id int) (int, error)
	FindProductNameFromId(int) (string, error)
	FindCategoryFromProductID(int) (int, error)
	FindPriceFromID(int) (float64, error)
}

package interfaces

import "cart/pkg/utils/models"

type ProductClient interface {
	GetStockOfProducts(int) (int, error)
	GetProductsDetails(products []int) ([]models.GetCart, error)
}

package usecase

import (
	"fmt"
	interfaces "product/pkg/repository/interface"
	services "product/pkg/usecase/interface"
	"product/pkg/utils/models"
)

type productUseCase struct {
	Repository interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		Repository: repo,
	}
}

func (p *productUseCase) AddProduct(inventory models.Inventories) (int, error) {

	fmt.Println("reached usecase")

	InventoryResponse, err := p.Repository.AddProduct(inventory)
	if err != nil {
		return InventoryResponse, err
	}

	return InventoryResponse, nil

}

func (i *productUseCase) ListProducts(page int) ([]models.Inventories, error) {

	productDetails, err := i.Repository.ListProducts(page)
	if err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil

}

func (i *productUseCase) GetStockOfProducts(id int) (int, error) {

	stock, err := i.Repository.GetStockOfProducts(id)
	if err != nil {
		return stock, err
	}

	return stock, nil

}

func (i *productUseCase) GetProductsDetails(products []int64) ([]models.GetCart, error) {

	var result []models.GetCart
	for _, v := range products {
		//find product name using v
		name, err := i.Repository.FindProductNameFromId(int(v))
		if err != nil {
			return []models.GetCart{}, err
		}
		//find category using v
		category, err := i.Repository.FindCategoryFromProductID(int(v))
		if err != nil {
			return []models.GetCart{}, err
		}

		//get price in total
		price, err := i.Repository.FindPriceFromID(int(v))
		if err != nil {
			return []models.GetCart{}, err
		}

		result = append(result, models.GetCart{
			ProductName: name,
			Category_id: category,
			Total:       price,
		})
	}

	return result, nil

}

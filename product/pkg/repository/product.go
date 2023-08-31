package repository

import (
	"fmt"
	interfaces "product/pkg/repository/interface"
	"product/pkg/utils/models"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) AddProduct(inventory models.Inventories) (int, error) {

	fmt.Println("reached repo")

	query := `
    INSERT INTO inventories (category_id, product_name, size, stock, price)
    VALUES (?, ?, ?, ?, ?);
    `
	err := p.DB.Exec(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price).Error
	if err != nil {
		return 500, err
	}

	return 200, nil

}

func (ad *productRepository) ListProducts(page int) ([]models.Inventories, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 5
	var productDetails []models.Inventories

	if err := ad.DB.Raw("select id,category_id,product_name,size,stock,price from inventories limit $1 offset $2", 5, offset).Scan(&productDetails).Error; err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil

}

func (ad *productRepository) GetStockOfProducts(id int) (int, error) {

	var stock int

	if err := ad.DB.Raw("select stock from inventories where id = $1", id).Scan(&stock).Error; err != nil {
		return stock, err
	}

	return stock, nil

}

func (ad *productRepository) FindProductNameFromId(id int) (string, error) {

	var name string

	if err := ad.DB.Raw("select product_name from inventories where id = $1", id).Scan(&name).Error; err != nil {
		return name, err
	}

	return name, nil

}

func (ad *productRepository) FindCategoryFromProductID(id int) (int, error) {

	var category int

	if err := ad.DB.Raw("select category_id from inventories where id = $1", id).Scan(&category).Error; err != nil {
		return category, err
	}

	return category, nil

}

func (ad *productRepository) FindPriceFromID(id int) (float64, error) {

	var price float64

	if err := ad.DB.Raw("select price from inventories where id = $1", id).Scan(&price).Error; err != nil {
		return price, err
	}

	return price, nil

}

package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (c *cartRepository) FindCart(id int) (int, error) {
	var cart int

	if err := c.DB.Raw("SELECT id FROM carts WHERE user_id=?", id).Scan(&cart).Error; err != nil {
		return 0, err
	}

	return cart, nil
}

func (c *cartRepository) CreateCart(id int) (int, error) {
	fmt.Println("call is triggered")
	var cart int
	err := c.DB.Exec(`
		INSERT INTO carts (user_id)
		VALUES ( ? )`, id).Error
	if err != nil {
		return 0, err
	}

	if err := c.DB.Raw("select id from carts where user_id=?", id).Scan(&cart).Error; err != nil {
		return 0, err
	}
	fmt.Println(cart)
	return cart, nil
}

func (c *cartRepository) CheckIfAlreadyExists(cart int, inv int) (bool, error) {
	var s int
	if err := c.DB.Raw("select count(*) from line_items where cart_id=$1 and inventory_id = $2", cart, inv).Scan(&s).Error; err != nil {
		return true, err
	}

	return s > 0, nil
}

func (c *cartRepository) AddToLineItems(cart_id int, inv_id int) error {
	err := c.DB.Exec(`
	INSERT INTO line_items (cart_id,inventory_id)
	VALUES ($1,$2)`, cart_id, inv_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *cartRepository) GetProductIDsFromCart(cart int) ([]int, error) {

	var products []int
	if err := c.DB.Raw("SELECT inventory_id FROM line_items WHERE cart_id=? AND ordered=false", cart).Scan(&products).Error; err != nil {
		return []int{}, err
	}

	return products, nil
}

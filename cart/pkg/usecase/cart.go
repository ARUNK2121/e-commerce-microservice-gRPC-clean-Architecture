package usecase

import (
	clienterface "cart/pkg/client/interface"
	interfaces "cart/pkg/repository/interface"
	services "cart/pkg/usecase/interface"
	"cart/pkg/utils/models"
	"errors"
	"fmt"
)

type cartUseCase struct {
	Repo          interfaces.CartRepository
	ProductClient clienterface.ProductClient
}

func NewCartUseCase(repo interfaces.CartRepository, client clienterface.ProductClient) services.CartUseCase {
	return &cartUseCase{
		Repo:          repo,
		ProductClient: client,
	}
}

func (c *cartUseCase) AddToCart(add models.AddToCart) error {
	//find cart of user
	cart, err := c.Repo.FindCart(add.UserID)
	if err != nil {
		return err
	}
	//if no cart then create new cart
	if cart == 0 {
		cart, err = c.Repo.CreateCart(add.UserID)
		if err != nil {
			return err
		}
	}
	//check if item already exists then return error
	exists, err := c.Repo.CheckIfAlreadyExists(cart, add.InventoryID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("already exists")
	}
	//check if product have stock available
	stock, err := c.ProductClient.GetStockOfProducts(add.InventoryID)
	if err != nil {
		return err
	}

	fmt.Println(stock)
	//add to line items with 1 quantity

	err = c.Repo.AddToLineItems(cart, add.InventoryID)
	if err != nil {
		return err
	}

	fmt.Println("i was right")

	return nil
}

func (c *cartUseCase) GetCart(id int) ([]models.GetCart, error) {
	//find cart id
	cart, err := c.Repo.FindCart(id)
	if err != nil {
		return []models.GetCart{}, err
	}
	// find line items id in cart
	products, err := c.Repo.GetProductIDsFromCart(cart)
	if err != nil {
		return []models.GetCart{}, err
	}

	fmt.Println(products)
	//find product details from products table
	details, err := c.ProductClient.GetProductsDetails(products)
	if err != nil {
		return []models.GetCart{}, err
	}
	//store product details into getcart model

	return details, nil
}

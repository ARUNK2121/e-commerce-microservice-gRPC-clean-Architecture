package interfaces

import "api_gateway/pkg/utils/models"

type CartClient interface {
	AddToCart(add models.AddToCart) error
	GetCart(id int) ([]models.GetCart, error)
}

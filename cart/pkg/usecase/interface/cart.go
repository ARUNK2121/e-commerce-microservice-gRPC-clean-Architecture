package interfaces

import "cart/pkg/utils/models"

type CartUseCase interface {
	AddToCart(models.AddToCart) error
	GetCart(id int) ([]models.GetCart, error)
}

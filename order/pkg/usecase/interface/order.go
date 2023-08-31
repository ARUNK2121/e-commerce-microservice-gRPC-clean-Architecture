package interfaces

import "order/pkg/utils/models"

type OrderUseCase interface {
	MakeOrder(cart_id, address_id int) error
	GetOrders(user_id int) ([]models.GetOrders, error)
	EditOrderStatus(int, string) error
}

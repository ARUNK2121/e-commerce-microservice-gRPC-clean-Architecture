package interfaces

import "api_gateway/pkg/utils/models"

type OrderClient interface {
	MakeOrder(cart_id, address_id int) (int, error)
	GetOrders(user_id int) ([]models.GetOrders, error)
	EditOrderStatus(models.EditStatus) (string, error)
}

package interfaces

import "order/pkg/utils/models"

type OrderRepository interface {
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	GetOrders(int) ([]models.GetOrders, error)
	EditOrderStatus(int, string) error
}

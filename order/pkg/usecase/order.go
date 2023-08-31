package usecase

import (
	"fmt"
	clienterface "order/pkg/client/interface"
	interfaces "order/pkg/repository/interface"
	services "order/pkg/usecase/interface"
	"order/pkg/utils/models"
)

type orderUseCase struct {
	Repository interfaces.OrderRepository
	CartClient clienterface.CartClient
}

func NewOrderUseCase(repo interfaces.OrderRepository, cart clienterface.CartClient) services.OrderUseCase {
	return &orderUseCase{
		Repository: repo,
		CartClient: cart,
	}
}

func (o *orderUseCase) MakeOrder(user_id, address_id int) error {
	cart, err := o.CartClient.GetCart(user_id)
	if err != nil {
		return err
	}

	var total float64
	for _, v := range cart {
		total = total + v.DiscountedPrice
	}

	payment_id := 1

	order_id, err := o.Repository.OrderItems(user_id, address_id, payment_id, total)
	if err != nil {
		return err
	}

	if err := o.Repository.AddOrderProducts(order_id, cart); err != nil {
		return err
	}

	fmt.Println(cart)
	return nil
}

func (o *orderUseCase) GetOrders(user_id int) ([]models.GetOrders, error) {

	details, err := o.Repository.GetOrders(user_id)
	if err != nil {
		return []models.GetOrders{}, nil
	}

	return details, err
}

func (o *orderUseCase) EditOrderStatus(order_id int, status string) error {

	err := o.Repository.EditOrderStatus(order_id, status)
	if err != nil {
		return err
	}

	return nil
}

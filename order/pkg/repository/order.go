package repository

import (
	"fmt"
	interfaces "order/pkg/repository/interface"
	"order/pkg/utils/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (i *orderRepository) OrderItems(userid, addressid, paymentid int, total float64) (int, error) {

	var id int
	query := `
    INSERT INTO orders (user_id,address_id, payment_method_id, final_price)
    VALUES (?, ?, ?, ?)
    RETURNING id
    `
	i.DB.Raw(query, userid, addressid, paymentid, total).Scan(&id)

	return id, nil

}

func (i *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

	query := `
    INSERT INTO order_items (order_id,product_name,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `

	for _, v := range cart {

		if err := i.DB.Exec(query, order_id, v.ProductName, v.Quantity, v.Total).Error; err != nil {
			return err
		}
	}

	return nil

}

func (i *orderRepository) GetOrders(id int) ([]models.GetOrders, error) {

	var model []models.GetOrders

	if err := i.DB.Raw("SELECT id,final_price,order_status FROM orders WHERE user_id = $1", id).Scan(&model).Error; err != nil {
		return []models.GetOrders{}, err
	}

	fmt.Println(model)

	return model, nil

}

func (i *orderRepository) EditOrderStatus(id int, status string) error {

	if err := i.DB.Exec("UPDATE orders SET order_status = $1 where id = $2", status, id).Error; err != nil {
		return err
	}

	return nil

}

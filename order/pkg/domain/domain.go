package domain

import "gorm.io/gorm"

type PaymentMethod struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}

type Order struct {
	gorm.Model
	UserID          uint    `json:"user_id" gorm:"not null"`
	AddressID       uint    `json:"address_id" gorm:"not null"`
	PaymentMethodID uint    `json:"paymentmethod_id"`
	FinalPrice      float64 `json:"price"`
	OrderStatus     string  `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string  `json:"payment_status" gorm:"payment_status:2;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID')"`
}

type OrderItem struct {
	ID          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     uint    `json:"order_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

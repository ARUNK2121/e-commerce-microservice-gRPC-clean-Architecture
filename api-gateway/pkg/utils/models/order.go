package models

type GetOrders struct {
	ID          int     `json:"id"`
	FinalPrice  float64 `json:"final_price"`
	OrderStatus string  `json:"order_status"`
}

type EditStatus struct {
	OrderID int    `json:"order_id"`
	Status  string `json:"status"`
}

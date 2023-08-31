package models

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}

type GetCart struct {
	ProductName     string  `json:"product_name"`
	Category_id     int     `json:"category_id"`
	Quantity        int     `json:"quantity"`
	Total           float64 `json:"total_price"`
	DiscountedPrice float64 `json:"discounted_price"`
}

type GetOrders struct {
	ID          int     `json:"id"`
	FinalPrice  float64 `json:"final_price"`
	OrderStatus string  `json:"order_status"`
}

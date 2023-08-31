package models

type Inventories struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type GetCart struct {
	ProductName     string  `json:"product_name"`
	Category_id     int     `json:"category_id"`
	Quantity        int     `json:"quantity"`
	Total           float64 `json:"total_price"`
	DiscountedPrice float64 `json:"discounted_price"`
}

package domain

type Inventories struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size" gorm:"size:5;default:'M';check:size IN ('S', 'M', 'L', 'XL', 'XXL')"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

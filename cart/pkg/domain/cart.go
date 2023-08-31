package domain

type Cart struct {
	ID     uint `json:"id" gorm:"primarykey"`
	UserID uint `json:"user_id" gorm:"not null"`
}

type LineItems struct {
	ID          uint `json:"id" gorm:"primarykey"`
	CartID      uint `json:"cart_id" gorm:"not null"`
	InventoryID uint `json:"inventory_id" gorm:"not null"`
	Quantity    int  `json:"quantity" gorm:"default:1"`
	Ordered     bool `json:"ordered" gorm:"default:false"`
}

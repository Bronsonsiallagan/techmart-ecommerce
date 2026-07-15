package models

type Cart struct {
	ID     uint       `json:"id" gorm:"primaryKey"`
	UserID uint       `json:"user_id" gorm:"unique"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
}
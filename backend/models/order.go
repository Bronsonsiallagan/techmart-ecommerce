package models

import "time"

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status" gorm:"default:pending"`
	PaymentProofURL  string      `json:"payment_proof_url"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
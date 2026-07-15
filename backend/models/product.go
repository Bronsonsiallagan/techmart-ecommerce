package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `json:"created_at"`
}
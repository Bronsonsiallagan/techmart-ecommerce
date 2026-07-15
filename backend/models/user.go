package models

import "time"

type User struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name"`
	Email            string    `json:"email" gorm:"unique"`
	Password         string    `json:"-"`
	Role             string    `json:"role" gorm:"default:customer"`
	Gender           string    `json:"gender"`
	AvatarURL        string    `json:"avatar_url"`
	SecurityQuestion string    `json:"security_question"`
	SecurityAnswer   string    `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
}
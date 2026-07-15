package utils

import (
	"techmart-backend/config"
	"techmart-backend/models"
)

func CreateNotification(userID uint, title string, message string) {
	notif := models.Notification{
		UserID:  userID,
		Title:   title,
		Message: message,
	}
	config.DB.Create(&notif)
}

// kirim notifikasi ke semua admin
func NotifyAllAdmins(title string, message string) {
	var admins []models.User
	config.DB.Where("role = ?", "admin").Find(&admins)

	for _, admin := range admins {
		CreateNotification(admin.ID, title, message)
	}
}
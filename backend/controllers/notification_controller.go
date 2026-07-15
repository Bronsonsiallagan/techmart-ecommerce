package controllers

import (
	"net/http"

	"techmart-backend/config"
	"techmart-backend/models"

	"github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var notifications []models.Notification
	config.DB.Where("user_id = ?", userID).Order("created_at desc").Limit(20).Find(&notifications)

	c.JSON(http.StatusOK, notifications)
}

func GetUnreadCount(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var count int64
	config.DB.Model(&models.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&count)

	c.JSON(http.StatusOK, gin.H{"unread_count": count})
}

func MarkAsRead(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))
	notifID := c.Param("id")

	var notif models.Notification
	if err := config.DB.Where("id = ? AND user_id = ?", notifID, userID).First(&notif).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notifikasi tidak ditemukan"})
		return
	}

	notif.IsRead = true
	config.DB.Save(&notif)

	c.JSON(http.StatusOK, notif)
}

func MarkAllAsRead(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	config.DB.Model(&models.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Update("is_read", true)

	c.JSON(http.StatusOK, gin.H{"message": "Semua notifikasi ditandai sudah dibaca"})
}
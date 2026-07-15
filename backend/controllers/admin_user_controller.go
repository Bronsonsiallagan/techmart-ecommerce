package controllers

import (
	"net/http"

	"techmart-backend/config"
	"techmart-backend/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers - admin only, lihat semua customer
func GetAllUsers(c *gin.Context) {
	var users []models.User
	query := config.DB.Where("role = ?", "customer")

	if search := c.Query("search"); search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUserDetail - admin only, lihat detail 1 customer + order-nya
func GetUserDetail(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	var orders []models.Order
	config.DB.Where("user_id = ?", userID).Preload("Items.Product").Order("created_at desc").Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"orders": orders,
	})
}
package controllers

import (
	"net/http"

	"techmart-backend/config"
	"techmart-backend/models"

	"github.com/gin-gonic/gin"
)

type MonthlyRevenue struct {
	Month string  `json:"month"`
	Total float64 `json:"total"`
}

type TopProduct struct {
	Name      string `json:"name"`
	TotalSold int    `json:"total_sold"`
}

type StatusCount struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

func GetAdminStats(c *gin.Context) {
	var totalCustomers int64
	config.DB.Model(&models.User{}).Where("role = ?", "customer").Count(&totalCustomers)

	var totalProducts int64
	config.DB.Model(&models.Product{}).Count(&totalProducts)

	var totalOrders int64
	config.DB.Model(&models.Order{}).Count(&totalOrders)

	var totalRevenue float64
	config.DB.Model(&models.Order{}).Where("status IN ?", []string{"paid", "shipped", "completed"}).
		Select("COALESCE(SUM(total_price), 0)").Scan(&totalRevenue)

	// pendapatan per bulan (6 bulan terakhir)
	var monthlyRevenue []MonthlyRevenue
	config.DB.Model(&models.Order{}).
		Select("DATE_FORMAT(created_at, '%Y-%m') as month, COALESCE(SUM(total_price), 0) as total").
		Where("status IN ? AND created_at >= DATE_SUB(NOW(), INTERVAL 6 MONTH)", []string{"paid", "shipped", "completed"}).
		Group("month").
		Order("month asc").
		Scan(&monthlyRevenue)

	// produk terlaris (top 5)
	var topProducts []TopProduct
	config.DB.Table("order_items").
		Select("products.name as name, SUM(order_items.quantity) as total_sold").
		Joins("JOIN products ON products.id = order_items.product_id").
		Group("products.name").
		Order("total_sold desc").
		Limit(5).
		Scan(&topProducts)

	// jumlah order per status
	var statusCounts []StatusCount
	config.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&statusCounts)

	c.JSON(http.StatusOK, gin.H{
		"total_customers": totalCustomers,
		"total_products":  totalProducts,
		"total_orders":    totalOrders,
		"total_revenue":   totalRevenue,
		"monthly_revenue": monthlyRevenue,
		"top_products":    topProducts,
		"status_counts":   statusCounts,
	})
}
package routes

import (
	"techmart-backend/controllers"
	"techmart-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// ==========================
	// Authentication
	// ==========================
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/forgot-password/question", controllers.GetSecurityQuestion)
		auth.POST("/forgot-password/reset", controllers.VerifyAndResetPassword)
	}

	// ==========================
	// Categories
	// ==========================
	categories := api.Group("/categories")
	{
		categories.GET("", controllers.GetCategories)
		categories.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.CreateCategory)
		categories.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.UpdateCategory)
		categories.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteCategory)
	}

	// ==========================
	// Products
	// ==========================
	products := api.Group("/products")
	{
		products.GET("", controllers.GetProducts)
		products.GET("/:id", controllers.GetProductByID)
		products.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.CreateProduct)
		products.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.UpdateProduct)
		products.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteProduct)
	}

	// ==========================
	// Cart - hanya customer
	// ==========================
	cart := api.Group("/cart")
	cart.Use(
		middleware.AuthMiddleware(),
		middleware.CustomerOnlyMiddleware(),
	)
	{
		cart.GET("", controllers.GetCart)
		cart.POST("", controllers.AddCartItem)
		cart.PUT("/:itemId", controllers.UpdateCartItem)
		cart.DELETE("/:itemId", controllers.DeleteCartItem)
	}

	// ==========================
	// Orders - hanya customer
	// ==========================
	orders := api.Group("/orders")
	orders.Use(
		middleware.AuthMiddleware(),
		middleware.CustomerOnlyMiddleware(),
	)
	{
		orders.POST("", controllers.Checkout)
		orders.GET("", controllers.GetMyOrders)
		orders.GET("/:id", controllers.GetOrderByID)
		orders.POST("/:id/payment-proof", controllers.UploadPaymentProof)
	}

	// ==========================
	// Admin Orders
	// ==========================
	adminOrders := api.Group("/admin/orders")
	adminOrders.Use(
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
	)
	{
		adminOrders.GET("", controllers.GetAllOrders)
		adminOrders.PUT("/:id/status", controllers.UpdateOrderStatus)
	}

	// ==========================
	// Upload Produk
	// ==========================
	upload := api.Group("/upload")
	upload.Use(
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
	)
	{
		upload.POST("/image", controllers.UploadImage)
	}

	// ==========================
	// Profile
	// ==========================
	profile := api.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	{
		profile.GET("", controllers.GetProfile)
		profile.PUT("", controllers.UpdateProfile)
		profile.PUT("/change-password", controllers.ChangePassword)
	}

	// ==========================
	// Upload Profile
	// ==========================
	uploadProfile := api.Group("/upload-profile")
	uploadProfile.Use(middleware.AuthMiddleware())
	{
		uploadProfile.POST("/image", controllers.UploadImage)
	}

	// ==========================
	// Admin Users
	// ==========================
	adminUsers := api.Group("/admin/users")
	adminUsers.Use(
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
	)
	{
		adminUsers.GET("", controllers.GetAllUsers)
		adminUsers.GET("/:id", controllers.GetUserDetail)
	}

	// ==========================
	// Notifications
	// ==========================
	notifications := api.Group("/notifications")
	notifications.Use(middleware.AuthMiddleware())
	{
		notifications.GET("", controllers.GetNotifications)
		notifications.GET("/unread-count", controllers.GetUnreadCount)
		notifications.PUT("/:id/read", controllers.MarkAsRead)
		notifications.PUT("/read-all", controllers.MarkAllAsRead)
	}

	// ==========================
	// Admin Statistics
	// ==========================
	adminStats := api.Group("/admin/stats")
	adminStats.Use(
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
	)
	{
		adminStats.GET("", controllers.GetAdminStats)
	}
}
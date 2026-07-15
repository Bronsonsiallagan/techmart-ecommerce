package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"techmart-backend/config"
	"techmart-backend/models"
	"techmart-backend/utils"

	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var cart models.Cart
	if err := config.DB.Where("user_id = ?", userID).Preload("Items.Product").First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Keranjang tidak ditemukan"})
		return
	}

	if len(cart.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keranjang kosong"})
		return
	}

	// hitung total & cek stok
	var totalPrice float64
	for _, item := range cart.Items {
		if item.Product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stok " + item.Product.Name + " tidak cukup"})
			return
		}
		totalPrice += item.Product.Price * float64(item.Quantity)
	}

	// buat order
	order := models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		Status:     "pending",
	}
	config.DB.Create(&order)

	// buat order items & kurangi stok
	for _, item := range cart.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}
		config.DB.Create(&orderItem)

		config.DB.Model(&models.Product{}).Where("id = ?", item.ProductID).
			Update("stock", item.Product.Stock-item.Quantity)
	}

	// kosongkan cart
	config.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	config.DB.Preload("Items.Product").First(&order, order.ID)

	// kirim notifikasi ke semua admin
	utils.NotifyAllAdmins("Pesanan Baru", fmt.Sprintf("Ada pesanan baru #%d senilai Rp%.0f", order.ID, order.TotalPrice))

	c.JSON(http.StatusCreated, order)
}

func GetMyOrders(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var orders []models.Order
	config.DB.Where("user_id = ?", userID).Preload("Items.Product").Order("created_at desc").Find(&orders)

	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))
	orderID := c.Param("id")

	var order models.Order
	if err := config.DB.Where("id = ? AND user_id = ?", orderID, userID).Preload("Items.Product").First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Admin: lihat semua order
func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Preload("Items.Product").Order("created_at desc").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

type UpdateOrderStatusInput struct {
	Status string `json:"status" binding:"required"`
}

// Admin: update status order
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}

	var input UpdateOrderStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.Status = input.Status
	config.DB.Save(&order)

	// kirim notifikasi ke customer pemilik order
	statusText := map[string]string{
		"paid":      "sudah dikonfirmasi dan dibayar",
		"shipped":   "sedang dikirim",
		"completed": "telah selesai",
	}
	if text, ok := statusText[input.Status]; ok {
		utils.CreateNotification(order.UserID, "Update Pesanan", fmt.Sprintf("Pesanan #%d Anda %s", order.ID, text))
	}

	c.JSON(http.StatusOK, order)
}

func UploadPaymentProof(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))
	orderID := c.Param("id")

	var order models.Order
	if err := config.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}

	if order.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order ini sudah tidak bisa diupload bukti pembayaran"})
		return
	}

	file, err := c.FormFile("proof")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File bukti transfer tidak ditemukan"})
		return
	}

	ext := filepath.Ext(file.Filename)
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".pdf": true}
	if !allowedExt[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file harus jpg, jpeg, png, webp, atau pdf"})
		return
	}

	newFileName := fmt.Sprintf("proof_%d_%d%s", order.ID, time.Now().UnixNano(), ext)
	savePath := "uploads/" + newFileName

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	proofURL := fmt.Sprintf("http://localhost:8080/uploads/%s", newFileName)

	order.PaymentProofURL = proofURL
	order.Status = "waiting_confirmation"
	config.DB.Save(&order)

	// kirim notifikasi ke semua admin
	utils.NotifyAllAdmins("Bukti Transfer Baru", fmt.Sprintf("Order #%d mengupload bukti transfer, perlu diverifikasi", order.ID))

	c.JSON(http.StatusOK, order)
}
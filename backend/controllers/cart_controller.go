package controllers

import (
	"net/http"

	"techmart-backend/config"
	"techmart-backend/models"

	"github.com/gin-gonic/gin"
)

// helper: ambil atau buat cart untuk user tertentu
func getOrCreateCart(userID uint) (models.Cart, error) {
	var cart models.Cart
	err := config.DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		cart = models.Cart{UserID: userID}
		if err := config.DB.Create(&cart).Error; err != nil {
			return cart, err
		}
	}
	return cart, nil
}

func GetCart(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	cart, err := getOrCreateCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil keranjang"})
		return
	}

	config.DB.Preload("Items.Product").First(&cart, cart.ID)
	c.JSON(http.StatusOK, cart)
}

type AddCartItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

func AddCartItem(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var input AddCartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek stok produk
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	cart, err := getOrCreateCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil keranjang"})
		return
	}

	// cek apakah item sudah ada di cart, kalau ada tinggal tambah quantity
	var existingItem models.CartItem
	err = config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).First(&existingItem).Error

	if err == nil {
		existingItem.Quantity += input.Quantity
		config.DB.Save(&existingItem)
	} else {
		newItem := models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		config.DB.Create(&newItem)
	}

	config.DB.Preload("Items.Product").First(&cart, cart.ID)
	c.JSON(http.StatusOK, cart)
}

type UpdateCartItemInput struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func UpdateCartItem(c *gin.Context) {
	itemID := c.Param("itemId")

	var input UpdateCartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var item models.CartItem
	if err := config.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
		return
	}

	item.Quantity = input.Quantity
	config.DB.Save(&item)

	c.JSON(http.StatusOK, item)
}

func DeleteCartItem(c *gin.Context) {
	itemID := c.Param("itemId")

	if err := config.DB.Delete(&models.CartItem{}, itemID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil dihapus dari keranjang"})
}
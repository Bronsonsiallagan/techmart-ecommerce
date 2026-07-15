package controllers

import (
	"net/http"
	"strconv"

	"techmart-backend/config"
	"techmart-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	query := config.DB.Preload("Category")

	if search := c.Query("search"); search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// hitung total data (sebelum pagination)
	var total int64
	query.Model(&models.Product{}).Count(&total)

	// pagination
	page := 1
	limit := 8
	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}
	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}
	offset := (page - 1) * limit

	query.Limit(limit).Offset(offset).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"total_page": (total + int64(limit) - 1) / int64(limit),
	})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, product)
}

type ProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageURL:    input.ImageURL,
		CategoryID:  input.CategoryID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat produk"})
		return
	}

	config.DB.Preload("Category").First(&product, product.ID)
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.ImageURL = input.ImageURL
	product.CategoryID = input.CategoryID

	config.DB.Save(&product)
	config.DB.Preload("Category").First(&product, product.ID)

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idUint, _ := strconv.Atoi(id)

	// cek apakah produk pernah muncul di order_items
	var orderItemCount int64
	config.DB.Model(&models.OrderItem{}).Where("product_id = ?", idUint).Count(&orderItemCount)

	if orderItemCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Produk tidak dapat dihapus karena sudah pernah dipesan. Set stok menjadi 0 jika ingin menghentikan penjualan."})
		return
	}

	// hapus dulu dari cart_items kalau ada (aman dihapus karena belum jadi order)
	config.DB.Where("product_id = ?", idUint).Delete(&models.CartItem{})

	if err := config.DB.Delete(&models.Product{}, idUint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}
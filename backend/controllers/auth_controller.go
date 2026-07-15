package controllers

import (
	"net/http"
	"strings"

	"techmart-backend/config"
	"techmart-backend/models"
	"techmart-backend/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=6"`
	SecurityQuestion string `json:"security_question" binding:"required"`
	SecurityAnswer   string `json:"security_answer" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah email sudah terdaftar
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password"})
		return
	}

	hashedAnswer, err := utils.HashPassword(strings.ToLower(strings.TrimSpace(input.SecurityAnswer)))
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data"})
	return
	}

	user := models.User{
	Name:             input.Name,
	Email:            input.Email,
	Password:         hashedPassword,
	Role:             "customer",
	SecurityQuestion: input.SecurityQuestion,
	SecurityAnswer:   hashedAnswer,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
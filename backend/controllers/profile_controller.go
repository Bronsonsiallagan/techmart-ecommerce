package controllers

import (
	"net/http"

	"techmart-backend/config"
	"techmart-backend/models"
	"techmart-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type UpdateProfileInput struct {
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender"`
	AvatarURL string `json:"avatar_url"`
}

func UpdateProfile(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	var input UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Gender = input.Gender

	if input.AvatarURL != "" {
		user.AvatarURL = input.AvatarURL
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui profil"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type ChangePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

func ChangePassword(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	var input ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek password lama
	if !utils.CheckPasswordHash(input.OldPassword, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password lama salah"})
		return
	}

	// Hash password baru
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password baru"})
		return
	}

	user.Password = hashedPassword

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengubah password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah",
	})
}
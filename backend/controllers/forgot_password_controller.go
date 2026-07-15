package controllers

import (
	"net/http"
	"strings"

	"techmart-backend/config"
	"techmart-backend/models"
	"techmart-backend/utils"

	"github.com/gin-gonic/gin"
)

type GetSecurityQuestionInput struct {
	Email string `json:"email" binding:"required,email"`
}

// Langkah 1: user masukkan email, sistem kembalikan pertanyaan keamanannya
func GetSecurityQuestion(c *gin.Context) {
	var input GetSecurityQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if user.SecurityQuestion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Akun ini belum memiliki pertanyaan keamanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"security_question": user.SecurityQuestion})
}

type VerifyAndResetInput struct {
	Email          string `json:"email" binding:"required,email"`
	SecurityAnswer string `json:"security_answer" binding:"required"`
	NewPassword    string `json:"new_password" binding:"required,min=6"`
}

// Langkah 2: verifikasi jawaban + set password baru sekaligus
func VerifyAndResetPassword(c *gin.Context) {
	var input VerifyAndResetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	normalizedAnswer := strings.ToLower(strings.TrimSpace(input.SecurityAnswer))
	if !utils.CheckPasswordHash(normalizedAnswer, user.SecurityAnswer) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Jawaban keamanan salah"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password baru"})
		return
	}

	user.Password = hashedPassword
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Password berhasil direset. Silakan login dengan password baru."})
}
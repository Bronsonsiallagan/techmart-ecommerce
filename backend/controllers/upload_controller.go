package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	// validasi ekstensi file
	ext := filepath.Ext(file.Filename)
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowedExt[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file harus jpg, jpeg, png, atau webp"})
		return
	}

	// buat nama file unik supaya tidak bentrok
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := "uploads/" + newFileName

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	imageURL := fmt.Sprintf("http://localhost:8080/uploads/%s", newFileName)

	c.JSON(http.StatusOK, gin.H{
		"image_url": imageURL,
	})
}
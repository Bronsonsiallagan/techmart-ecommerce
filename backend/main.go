package main

import (
	"log"
	"time"

	"techmart-backend/config"
	"techmart-backend/models"
	"techmart-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Tidak menemukan file .env, menggunakan environment variable sistem")
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Notification{},
	)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// serve folder uploads sebagai file statis, bisa diakses lewat /uploads/namafile.jpg
	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	r.Run(":8080")
}
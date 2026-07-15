package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role == "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin tidak dapat mengakses fitur belanja"})
			c.Abort()
			return
		}
		c.Next()
	}
}
package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc { // Middleware untuk menangani server agar tidak crash saat panic (terjadi hal tidak terduga)
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log panic
				log.Printf("PANIC RECOVERED: %v", err)

				// Berikan error message
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "An unexpected error occurred. Please try again later.",
				})

				c.Abort()
			}
		}()

		// Lanjut ke handler selanjutnya
		c.Next()
	}
}

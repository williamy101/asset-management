package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc { // Middleware untuk logging; waktu eksekusi request dan response, status code, IP, dll
	return func(c *gin.Context) {
		// mulai timer
		startTime := time.Now()

		// lanjut prosesnya
		c.Next()

		// hitung durasi eksekusi
		duration := time.Since(startTime)

		// Log detail request
		log.Printf("HTTP METHOD: %s | ROUTE: %s | CODE: %d | DURATION: %v | IP: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
			c.ClientIP(),
		)
	}
}

package middleware

import (
	"go-asset-management/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware untuk memproses dan memvalidasi token
func AuthMiddleware(requiredRole ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Validasi token
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Cek role
		roleAllowed := false
		for _, role := range requiredRole {
			if claims.Role == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)

		// lanjut ke proses lain
		log.Println("[DEBUG] Middleware passed. Proceeding to next handler.")
		c.Next()
	}
}

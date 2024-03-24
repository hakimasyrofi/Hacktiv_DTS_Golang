package middleware

import (
	"MyGramProject/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Authorization token is required")
			c.Abort()
			return
		}

		// Parse dan validasi token
		userID, err := helpers.ParseToken(tokenString)
		if err != nil {
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Simpan userID dalam context untuk digunakan di handler
		c.Set("userID", userID)

		c.Next()
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/ANSHSINGH050404/movie_streaming/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		// Handle "Bearer <token>" format
		extractedToken := strings.Split(clientToken, " ")
		if len(extractedToken) == 2 {
			clientToken = extractedToken[1]
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", (*claims)["email"])
		c.Set("user_id", (*claims)["user_id"])
		c.Next()
	}
}

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-tasks-app-practice/internal/handlers"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("here is your token", tokenStr)

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return handlers.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token" + err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Claims"})
			c.Abort()
			return
		}
		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

		c.Next()

	}
}

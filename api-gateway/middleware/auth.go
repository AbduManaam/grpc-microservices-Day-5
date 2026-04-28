package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("super-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")

		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return
		}

		tokenstr := strings.TrimPrefix(auth, "Bearer ")

		_, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secret, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}
		c.Next()
	}
}
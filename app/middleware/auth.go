package middleware

import (
	"fmt"
	"github.com/fthyilmz/workshop-go.git/app/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		tokenString = strings.ReplaceAll(tokenString, " ", "")

		token, err := utils.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		c.Next()
	}
}

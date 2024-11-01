package authentication

import (
	"net/http"

	"github.com/CrudOperationUsingAuthentication/pkg/models"
	"github.com/dgriJalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Jwtkey = []byte("your_secret_key_my")

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token Provided"})
			c.Abort()
			return
		}
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return Jwtkey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("username", claims.UserName)
		c.Next()
	}
}

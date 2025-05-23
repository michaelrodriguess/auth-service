package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/michaelrodriguess/auth_service/config"
	"github.com/michaelrodriguess/auth_service/internal/repository"
)

type CustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware(repo *repository.UserAuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		isBlocked, err := repo.IsTokenBlocked(context.TODO(), tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalid"})
			c.Abort()
			return
		}

		if isBlocked {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
			c.Abort()
			return
		}

		secretKey := config.GetJWTSecret()

		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalid"})
			c.Abort()
			return
		}

		if claims.Subject == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token without subject"})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

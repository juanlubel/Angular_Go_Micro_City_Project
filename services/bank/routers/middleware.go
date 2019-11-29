package router

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// APIError struct
type APIError struct {
	Code    int
	Message string
}

// RespondWithError function handler error
func RespondWithError(code int, message string, c *gin.Context) {
	c.JSON(code, &APIError{Code: code, Message: message})
	c.Abort()
}

// JWTAuthMiddleware middleware function implementation
func JWTAuthMiddleware(encoded bool, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		parts := strings.Fields(auth)

		// Token base validation
		if auth == "" {
			RespondWithError(401, "API token required", c)
			return
		}

		// Token base validation
		if strings.ToLower(parts[0]) != "bearer" {
			RespondWithError(401, "Authorization header must start with Bearer", c)
			return
		} else if len(parts) == 1 {
			RespondWithError(401, "Token not found", c)
			return
		} else if len(parts) > 2 {
			RespondWithError(401, "Authorization header must be Bearer and token", c)
			return
		}

		// parse the encrypted string to JSON Web Tokens
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err == nil && token.Valid {
			c.Next()
		} else {
			RespondWithError(401, err.Error(), c)
			return
		}

	}
}

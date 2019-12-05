package router

import (
	"os"
	"strings"

	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/danilopolani/gocialite"
	"github.com/gin-gonic/gin"
)

var gocial = gocialite.NewDispatcher()

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
func JWTAuthMiddleware(checkauth bool, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		parts := strings.Fields(auth)

		// Token base validation
		if auth == "" {
			if checkauth {
				RespondWithError(401, "API token required", c)
			}
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

// Redirect to correct oAuth URL
func redirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("clientID"),
			"clientSecret": os.Getenv("clientSecret"),
			"redirectURL":  "http://bank.localhost:3010/auth/github/callback",
		},
		/* 		"google": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/google/callback",
		}, */
	}

	providerScopes := map[string][]string{
		"github": []string{"public_repo"},
		/* "google": []string{}, */

	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func callbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	/* provider := c.Param("provider") */

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)

	// If no errors, show provider name
	c.Writer.Write([]byte("Hi, " + user.FullName))
}

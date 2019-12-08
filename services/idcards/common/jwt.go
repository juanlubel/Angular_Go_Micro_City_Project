package common

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var Secret = os.Getenv("SECRET")

func GenerateJWT(id string) (string, error) {
	duration, err := time.ParseDuration("30m")
	if err != nil {
		return "", err
	}

	// the token expires after 30 minutes
	expiration := time.Now().Add(duration).Unix()

	secret := []byte(Secret)

	claims := struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}{
		id,
		jwt.StandardClaims{
			ExpiresAt: expiration,
		},
	}

	// generate JSON Web Tokens
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// transform token to encrypted string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (string, error) {
	claims := struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}{}

	// parse the encrypted string to JSON Web Tokens
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err == nil && token.Valid {
		return claims.ID, nil
	} else {
		return "", nil
	}
}

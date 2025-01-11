package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(Email string) (string, error) {
	cfg := Config

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(cfg.AccessTokenExpiry).Unix(),
		"email": Email,
	})

	return token.SignedString([]byte(cfg.AccessTokenSecret))
}

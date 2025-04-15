package utils

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(privateKey string, sub string, exp time.Duration, payload interface{}) (string, error) {
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"iss": "auth-service",
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

}

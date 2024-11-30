package common

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetJwtSecretKey() []byte {
    jwtSecretKey, exists := os.LookupEnv("JWT_SECRET_KEY")
    if !exists {
        panic("\"JWT_SECRET_KEY\" is not set")
    }
    return []byte(jwtSecretKey)
}

func GenerateJwt(claims jwt.Claims) (*string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
    jwtToken, err := token.SignedString(GetJwtSecretKey())
    if err != nil {
        return nil, err
    }
    return &jwtToken, nil
}
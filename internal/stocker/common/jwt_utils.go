package common

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func GetUserIdByContext(ctx *fiber.Ctx) (*uuid.UUID, error) {
    // JWT Claimsを取得
    user := ctx.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)

    // "sub"をUUIDに変換
    userId, err := uuid.Parse(claims["sub"].(string))
    if err != nil {
        return nil, err
    }
    return &userId, nil
}
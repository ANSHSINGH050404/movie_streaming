package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var SECRET_KEY string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")

	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY not found in environment or .env file")
	}
}

func GenerateJWT(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func ValidateToken(signedToken string) (*jwt.MapClaims, string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, err.Error()
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, "Invalid claims"
	}

	return claims, ""
}

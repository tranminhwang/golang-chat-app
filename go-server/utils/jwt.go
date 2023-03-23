package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type TokenPayload struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func CreateToken(p TokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      p.ID,
		"email":    p.Email,
		"username": p.Username,
		"exp":      float64(time.Now().Add(time.Hour * 24).Unix()),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", fmt.Errorf("failed create token")
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	return token, err
}

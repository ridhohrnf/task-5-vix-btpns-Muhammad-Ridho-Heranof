package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("your-secret-key"))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	userId := claims["userId"].(float64)

	return uint(userId), nil
}

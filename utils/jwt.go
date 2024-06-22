package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "REST-API"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return SECRET_KEY, nil
	})

	if err != nil {
		return errors.New("Could not parse token.")
	}

	tokenIsValid := token.Valid
	if !tokenIsValid {
		return errors.New("Invalid token.")
	}

	// claims, ok := token.Claims.(jwt.MapClaims);
	// if  !ok {
	// 	return errors.New("Invalid token claims.")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}

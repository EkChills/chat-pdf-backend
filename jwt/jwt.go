package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateSignedToken(username, email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"email":    email,
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", errors.New("could not generate token")
	}

	return tokenString, err
}

func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid claims")
	}

	userIdFloat, ok := claims["userId"].(float64)

	if !ok {
		return 0, errors.New("invalid user id")
	}

	return int64(userIdFloat), nil

}

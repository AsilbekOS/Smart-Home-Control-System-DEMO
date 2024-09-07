package getid

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GetUserIDFromToken(tokenString string, secretKey string) (string, error) {
	key := []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return "", fmt.Errorf("tokenni dekodlashda xatolik: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userId"].(string)
		if !ok {
			return "", fmt.Errorf("user_id mavjud emas")
		}

		return userID, nil
	} else {
		return "", fmt.Errorf("token yaroqsiz")
	}
}

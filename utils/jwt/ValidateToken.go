package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func ValidationToken(accessToken string) (bool, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return false, errors.New("неверный access token")
	}

	// Можно добавить получением claims (логина в заголовке по надобности)
	return true, nil
}

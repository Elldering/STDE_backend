package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AccessTokenExpiry  = time.Minute * 15
	RefreshTokenExpiry = time.Hour * 24 * 5
)

func GenerateTokens(user *models.AuthUserRequest, JWTSecret string) (string, string, error) {
	if JWTSecret == "" {
		return "", "", errors.New("JWT_SECRET не настроен")
	}

	if err := validation.ValidateEmptyFields(user.Login); err != nil {
		return "", "", errors.New("логин не может быть пустым")
	}

	if err := repositories.UpdateLastLogin(user); err != nil {
		return "", "", fmt.Errorf("ошибка при попытке обновить время входа: %v", err)
	}

	// Генерация access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Login,
		"exp":      time.Now().Add(AccessTokenExpiry).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("ошибка при подписании access токена: %v", err)
	}

	// Генерация refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Login,
		"exp":      time.Now().Add(RefreshTokenExpiry).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("ошибка при подписании refresh токена: %v", err)
	}

	return accessTokenString, refreshTokenString, nil
}

package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/validation"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateTokens создает access и refresh токены для пользователя
func GenerateTokens(user *models.AuthUser, JWTSecret string) (string, string, error) {

	err := repositories.UpdateLastLogin(user)
	if err != nil {
		return "", "", fmt.Errorf(" ошибка попытке обновить время входа: %v", err)
	}

	// Генерация access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Access токен действует 15 минут
	})

	accessTokenString, err := accessToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("ошибка при подписании access токена: %v", err)
	}

	// Генерация refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 5).Unix(), // Refresh токен действует 7 дней
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("ошибка при подписании refresh токена: %v", err)
	}

	return accessTokenString, refreshTokenString, nil
}

// GenerateAccessToken создает только access токен для пользователя
func GenerateAccessToken(user *models.AuthUser, JWTSecret string) (string, error) {
	// Генерация access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Access токен действует 15 минут
	})

	accessTokenString, err := accessToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", fmt.Errorf("ошибка при подписании access токена: %v", err)
	}

	return accessTokenString, nil
}

// Authenticate аутентифицирует пользователя на основе логина и пароля
func Authenticate(data models.AuthUser, JWTSecret string) (string, string, error) {

	err := validation.CheckEmailOrPhoneNumber(&data)
	if err != nil {
		return "", "", err
	}

	user, err := repositories.FindByUsername(data.Login)
	if err != nil {
		return "", "", err
	}

	err = repositories.CheckVerifyEmail(data.Login)
	if err != nil {
		return "", "", err
	}

	if !validation.ValidatePassword(data.Password) {
		return "", "", errors.New("некорректный пароль")
	}

	if !hash.CheckPasswordHash(data.Password, user.Password) {
		return "", "", fmt.Errorf("неверные учетные данные")
	}

	accessToken, refreshToken, err := GenerateTokens(user, JWTSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

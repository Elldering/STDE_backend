package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// RefreshToken обновляет access токен на основе refresh токена
func RefreshToken(data models.AuthUser, JWTSecret string) (string, error) {
	// Проверяем, не отозван ли refresh токен
	isInvalidRefresh, err := repositories.IsRefreshTokenInvalidated(data.RefreshToken)
	if err != nil {
		log.Printf("Ошибка проверки токена: %v", err)
		return "", fmt.Errorf("ошибка проверки токена")
	}
	if isInvalidRefresh {
		return "", fmt.Errorf("refresh токен недействителен")
	}

	// Парсим refresh токен
	token, err := jwt.Parse(data.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		log.Printf("Ошибка при разборе токена: %v", err)
		return "", fmt.Errorf("неверный токен")
	}

	// Извлекаем claims из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("неверный токен")
	}

	// Находим пользователя по username из claims
	data.Login = claims["username"].(string)
	log.Print(data.Login)
	err = validation.CheckEmailOrPhoneNumber(&data)

	user, err := repositories.FindByUsername(data)
	if err != nil {
		return "", fmt.Errorf("пользователь не найден")
	}

	//// Генерируем новый access токен
	accessToken, _, err := GenerateTokens(user, JWTSecret)
	if err != nil {
		return "", fmt.Errorf("ошибка генерации токена: %v", err)
	}

	return accessToken, nil
}

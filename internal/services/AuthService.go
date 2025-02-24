package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/validation"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// GenerateTokens создает access и refresh токены для пользователя
func GenerateTokens(user *models.AuthUser, JWTSecret string) (string, string, error) {
	log.Printf("В GenerateTokens: user.Login = %s", user.Login)
	err := repositories.UpdateLastLogin(user)
	if err != nil {
		return "", "", fmt.Errorf(" ошибка попытке обновить время входа: %v", err)
	}
	// Генерация access токена

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Login,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Access токен действует 15 минут
	})

	accessTokenString, err := accessToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("ошибка при подписании access токена: %v", err)
	}

	// Генерация refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Login,
		"exp":      time.Now().Add(time.Hour * 24 * 5).Unix(), // Refresh токен действует 7 дней
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("ошибка при подписании refresh токена: %v", err)
	}

	return accessTokenString, refreshTokenString, nil
}

// Authenticate аутентифицирует пользователя на основе логина и пароля
func Authenticate(data models.AuthUser, JWTSecret string) (string, string, error) {

	err := validation.CheckEmailOrPhoneNumber(&data)
	if err != nil {
		return "", "", err
	}

	user, err := repositories.FindByUsername(data)
	if err != nil {
		return "", "", err
	}
	log.Println(data.Password)
	switch data.TypeLogin {
	case "email":
		err = repositories.CheckVerifyEmail(data.Login)
		if err != nil {
			return "", "", err
		}
	case "phone_number":
		log.Println(data.Login)
	}

	log.Println(data.Password)
	log.Println(user.Password)
	//if !validation.ValidatePassword(data.Password) {
	//	return "", "", errors.New("некорректный пароль")
	//}

	if !hash.CheckPasswordHash(data.Password, user.Password) {
		return "", "", fmt.Errorf("неверные учетные данные")
	}

	accessToken, refreshToken, err := GenerateTokens(user, JWTSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware проверяет JWT токен в заголовке Authorization
//
// Параметры:
// - JWTSecret: Секретный ключ для проверки подписи JWT токена
//
// Возвращает:
// - gin.HandlerFunc: Функция-обработчик для Gin Middleware
//
// Обработчик:
// - Получает токен из заголовка Authorization
// - Проверяет наличие токена в заголовке
// - Удаляет префикс "Bearer " из токена
// - Проверяет валидность токена и его подписи
// - Извлекает и устанавливает значение username из claims в контекст
// - Возвращает ошибку 401, если токен не валиден или отсутствует
func AuthMiddleware(JWTSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Нет токена в заголовке"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):] // Удаляем префикс "Bearer "

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWTSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		c.Next()
	}
}

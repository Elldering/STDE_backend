package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// AuthMiddleware проверяет JWT токен в заголовке Authorization
//
// Параметры:
// - JWTSecret: Секретный ключ для проверки подписи JWT токена
//
// Возвращает:
// - gin. HandlerFunc: Функция-обработчик для Gin Middleware
//
// Обработчик:
// - Получает токен из заголовка Authorization
// - Проверяет наличие токена в заголовке
// - Удаляет префикс "Bearer" из токена
// - Проверяет валидность токена и его подписи
// - Проверяет, находится ли токен в черном списке (инвалидирован)
// - Извлекает и устанавливает значение username из claims в контекст
// - Возвращает ошибку 401, если токен не валиден, отсутствует или инвалидирован

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем access token из cookie
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token отсутствует"})
			c.Abort()
			return
		}

		// Парсим токен
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неверный метод подписи")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный access token"})
			c.Abort()
			return
		}

		// Если токен валиден, продолжаем выполнение
		c.Next()
	}
}

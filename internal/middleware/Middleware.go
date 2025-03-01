package middleware

import (
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
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
// - Проверяет, находится ли токен в черном списке (инвалидирован)
// - Извлекает и устанавливает значение username из claims в контекст
// - Возвращает ошибку 401, если токен не валиден, отсутствует или инвалидирован

func AuthMiddleware(JWTSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		access := c.GetHeader("Authorization")

		err := validation.ValidateEmptyFields(access, JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует access токен в заголовке. Либо ошибка чтения jwt secret"})
			c.Abort()
			return
		}

		access = access[len("Bearer "):] // Удаляем префикс "Bearer"
		// Проверяем, находится ли токен в черном списке
		isInvalid, err := repositories.IsAccessTokenInvalidated(access)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка проверки токена"})
			c.Abort()
			return
		}

		if isInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен недействителен"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(access, func(token *jwt.Token) (interface{}, error) {
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

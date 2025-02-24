package Auth

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// RefreshToken обрабатывает запросы на обновление JWT access токена
//
// Параметры:
// - ctx: Контекст Gin, содержащий запрос и ответ
//
// Обработчик:
// - Парсит JSON тело запроса, содержащее refresh токен
// - Проверяет корректность данных и наличие префикса "Bearer"
// - Удаляет префикс "Bearer" из токена
// - Проверяет валидность токена и извлекает claims
// - Ищет пользователя по username из claims
// - Генерирует новый access токен и возвращает его в ответе
// - Возвращает статус ошибки и сообщение в случае неудачи
func RefreshToken(c *gin.Context) {
	var data models.AuthUser
	// Парсим JSON-тело запроса
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh токен не был передан"})
		return
	}

	// Убираем префикс "Bearer" из токена
	if strings.HasPrefix(data.RefreshToken, "Bearer ") {
		data.RefreshToken = strings.TrimPrefix(data.RefreshToken, "Bearer ")
	}

	// Вызываем сервис для обновления токена
	JWTSecret := os.Getenv("JWT_SECRET")
	accessToken, err := services.RefreshToken(data, JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем новый access токен
	c.JSON(http.StatusOK, gin.H{
		"access": accessToken,
	})
}

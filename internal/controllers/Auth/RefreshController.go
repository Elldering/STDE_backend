package Auth

import (
	"STDE_proj/internal/controllers/Auth/model"
	"STDE_proj/internal/repositories"
	"STDE_proj/internal/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
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

	var refresh model.LogoutRequest

	if err := c.ShouldBindJSON(&refresh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh токен не был передан"})
		return
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	//log.Printf("Полученный Refresh токен: %s", refresh) // Логирование полученного токена

	// Проверка на наличие префикса Bearer
	if strings.HasPrefix(refresh.Refresh, "Bearer ") {
		refresh.Refresh = strings.TrimPrefix(refresh.Refresh, "Bearer ")
	}

	isInvalid, err := repositories.IsRefreshTokenInvalidated(refresh.Refresh)
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

	token, err := jwt.Parse(refresh.Refresh, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		log.Printf("Ошибка при разборе токена: %v", err) // Логирование ошибки разбора токена
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := repositories.FindByUsername(claims["username"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
			return
		}

		// Генерируем только новый access токен
		accessToken, _, err := services.GenerateTokens(user, JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access": accessToken,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
	}
}

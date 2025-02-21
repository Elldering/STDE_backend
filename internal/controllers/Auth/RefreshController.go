package Auth

import (
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
func RefreshToken(ctx *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	log.Printf("Полученный Refresh токен: %s", request.RefreshToken) // Логирование полученного токена

	// Проверка на наличие префикса Bearer
	if !strings.HasPrefix(request.RefreshToken, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный формат токена"})
		return
	}

	// Удаление префикса Bearer
	tokenString := strings.TrimPrefix(request.RefreshToken, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		log.Printf("Ошибка при разборе токена: %v", err) // Логирование ошибки разбора токена
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := repositories.FindByUsername(claims["username"].(string))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
			return
		}

		// Генерируем только новый access токен
		accessToken, _, err := services.GenerateTokens(user, JWTSecret)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
	}
}

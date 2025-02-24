package Auth

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// LoginHandler обрабатывает запросы на авторизацию пользователя
//
// Параметры:
// - ctx: Контекст Gin, содержащий запрос и ответ
//
// Обработчик:
// - Парсит JSON тело запроса, содержащего логин и пароль пользователя
// - Проверяет корректность данных
// - Аутентифицирует пользователя с помощью `services.Authenticate`
// - Возвращает JWT access и refresh токены в случае успешной аутентификации
// - Возвращает статус ошибки и сообщение в случае неудачи
func LoginHandler(ctx *gin.Context) {
	
	var data models.AuthUser

	JWTSecret := os.Getenv("JWT_SECRET")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := services.Authenticate(data, JWTSecret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access": accessToken, "refresh": refreshToken})
}

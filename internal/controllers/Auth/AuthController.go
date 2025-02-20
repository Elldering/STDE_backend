package Auth

import (
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
	// Анонимная структура, в которую будут подставлены логин и пароль пользователя
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	JWTSecret := os.Getenv("JWT_SECRET")

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := services.Authenticate(loginRequest.Username, loginRequest.Password, JWTSecret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access token": accessToken, "refresh token": refreshToken})
}

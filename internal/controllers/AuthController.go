package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
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

func AuthenticationHandler(ctx *gin.Context) {

	var data models.AuthUserRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.Authentication(data)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Аутентификация успешна. Код доступа отправлен"})
}

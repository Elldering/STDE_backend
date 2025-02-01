package controllers

import (
	"STDE_proj/internal/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetUserGroupsHandler обрабатывает запросы на получение всех групп пользователей.
// Использует метод GetAllUserGroups из пакета repositories для получения данных.
// В случае ошибки возвращает HTTP статус 500 и сообщение об ошибке.
// В случае успеха возвращает HTTP статус 200 и JSON с данными групп пользователей.
func GetUserGroupsHandler(c *gin.Context) {
	// Получаем все группы пользователей из репозитория
	userGroups, err := repositories.GetAllUserGroups()
	if err != nil {
		// Логируем ошибку и возвращаем HTTP статус 500 и сообщение об ошибке
		log.Printf("Ошибка при получении групп пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	// Возвращаем HTTP статус 200 и JSON с данными групп пользователей
	c.JSON(http.StatusOK, userGroups)
}

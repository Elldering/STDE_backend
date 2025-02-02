package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// PostAuthGroupPermissionsHandler обрабатывает запросы на создание связи группы и прав доступа.
// Использует метод CreateAuthGroupPermission из пакета services для создания связи.
// В случае ошибки возвращает HTTP статус 500 и сообщение об ошибке.
// В случае успеха возвращает HTTP статус 200 и JSON с данными связи.
func PostAuthGroupPermissionsHandler(c *gin.Context) {
	var agp models.AuthGroupPermissions
	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateAuthGroupPermission(agp)
	if err != nil {
		log.Printf("Ошибка при создании связи группы и прав доступа: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании связи"})
		return
	}

	c.JSON(http.StatusOK, agp)
}

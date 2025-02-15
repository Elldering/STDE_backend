package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAuthGroupPermissionsHandler(c *gin.Context) {
	AuthPermissionsGroups, err := services.GetAuthGroupPermissions()
	if err != nil {
		log.Println("Ошибка при выводе данных: ", err)
		c.JSON(http.StatusOK, gin.H{"error": "Ошибка при выводе данных"})
		return
	}
	c.JSON(http.StatusOK, AuthPermissionsGroups)
}

func GetAuthGroupPermissionsIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	data, err := services.GetAuthGroupPermissionsId(id)
	if err != nil {
		log.Printf("Ошибка при получение данных о связи групп и прав доступа:%s ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PutAuthGroupPermissionsHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var data models.AuthGroupPermissions
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err = services.PutAuthGroupPermissions(id, data)
	if err != nil {
		log.Printf("Ошибка при получения связи групп и прав доступа:%s ", err)
		c.JSON(http.StatusOK, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func DeleteAuthGroupPermissionsHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeleteAuthGroupPermissions(id)
	if err != nil {
		log.Printf("Ошибка при удалении связи группы с правами доступа с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении связи группы с правами доступа"})
		return
	}
	log.Printf("связи группы с правами доступа с ID %d успешно удалено", id)
	c.JSON(http.StatusOK, gin.H{"message": "связи группы с правами доступа успешно удалено"})
}

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

	err := services.PostAuthGroupPermission(agp)
	if err != nil {
		log.Printf("Ошибка при создании связи группы и прав доступа: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании связи"})
		return
	}

	c.JSON(http.StatusOK, agp)
}

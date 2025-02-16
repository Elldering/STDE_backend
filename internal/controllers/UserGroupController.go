package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUserGroupsHandler обрабатывает запросы на получение всех групп пользователей.
// Использует метод GetAllUserGroups из пакета repositories для получения данных.
// В случае ошибки возвращает HTTP статус 500 и сообщение об ошибке.
// В случае успеха возвращает HTTP статус 200 и JSON с данными групп пользователей.
func GetUserGroupsHandler(c *gin.Context) {
	// Получаем все группы пользователей из репозитория
	userGroups, err := services.GetAllUserGroups()
	if err != nil {
		// Логируем ошибку и возвращаем HTTP статус 500 и сообщение об ошибке
		log.Printf("Ошибка при получении групп пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	// Возвращаем HTTP статус 200 и JSON с данными групп пользователей
	c.JSON(http.StatusOK, userGroups)
}
func GetUserGroupByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Ошибка при преобразовании id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	userGroup, err := services.GetUserGroupById(id)
	if err != nil {
		log.Printf("Ошибка при получении групы пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, userGroup)

}
func PostUserGroupHandler(c *gin.Context) {
	var agp models.UserGroup
	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.PostUserGroup(agp)
	if err != nil {
		log.Printf("Ошибка при создании связи группы и прав доступа: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании связи"})
		return
	}

	c.JSON(http.StatusOK, agp)
}

func PutUserGroupHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var agp models.UserGroup
	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.PutUserGroup(id, agp)
	if err != nil {
		log.Printf("Ошибка при обновлении группы пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении группы пользователей"})
		return
	}

	c.JSON(http.StatusOK, agp)
}
func DeleteUserGroupHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Ошибка при преобразовании id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeleteUserGroup(id)
	if err != nil {
		log.Printf("Ошибка при удалении группы пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении группы пользователей"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Группа пользователей успешно удалена"})
}

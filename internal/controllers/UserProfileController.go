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

func GetUserProfileHandler(c *gin.Context) {

	userProfile, err := services.GetAllUserProfile()
	if err != nil {

		log.Printf("Ошибка при получении профелей пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func GetUserProfileByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Ошибка при преобразовании id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	userProfile, err := services.GetUserProfileById(id)
	if err != nil {
		log.Printf("Ошибка при получении: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, userProfile)
}

func PostUserProfileHandler(c *gin.Context) {
	var agp models.UserProfile

	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.PostUserProfile(agp)
	if err != nil {
		log.Printf("Ошибка при добавлении нового профиля: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании профиля"})
		return
	}

	c.JSON(http.StatusOK, agp)
}

func PutUserProfileHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var agp models.UserProfile
	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.PutUserProfile(id, agp)
	if err != nil {
		log.Printf("Ошибка при обновлении профилей пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении профилей пользователей"})
		return
	}
	c.JSON(http.StatusOK, agp)
}

func DeleteUserProfileHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Ошибка при преобразовании id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeleteUserProfile(id)
	if err != nil {
		log.Printf("Ошибка при удалении профиля пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении профиля пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль пользователя успешно удален"})
}

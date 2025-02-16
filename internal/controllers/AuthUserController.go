package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//func GetAllAuthUserHandler(c *gin.Context) {
//	data, err := services.GetAllAuthUser()
//	if err != nil {
//		log.Printf("Не удалось получить список пользователь: %v", err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
//		return
//	}
//	c.JSON(http.StatusOK, data)
//}

func PostAuthUserHandler(c *gin.Context) {
	var data models.AuthUser
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err := services.PostAuthUser(data)
	if err != nil {
		log.Printf("Ошибка при создании пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func DeleteAuthUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeleteAuthUser(id)
	if err != nil {
		log.Printf("Ошибка при удалении пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
		return
	}

	log.Printf("Пользователь с ID %d успешно удалено", id)
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})
}

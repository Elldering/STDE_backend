package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAuthUserGroupsAllHandler(c *gin.Context) {
	data, err := services.GetAuthUserGroupsAll()
	if err != nil {
		log.Printf("Не удалось получить список связаных групп с пользователями: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetAuthUserGroupsByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetAuthUserGroupsById(id)
	if err != nil {
		log.Printf("Ошибка при получении связи группы и пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostAuthUserGroupsHandler(c *gin.Context) {
	var data models.AuthUserGroups
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostAuthUserGroups(data)
	if err != nil {
		log.Printf("Ошибка при создании связи групп и пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании групп и пользователей"})
		return
	}

	log.Println(" Право доступа успешно создано")
	c.JSON(http.StatusOK, data)
}

func PutAuthUserGroupsHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var data models.AuthUserGroups
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err = services.PutAuthUserGroups(id, data)
	if err != nil {
		log.Printf("Ошибка при получении связи группы и пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	log.Printf("Право доступа с ID %d успешно обновлено", id)

	c.JSON(http.StatusOK, data)
}

func DeleteAuthUserGroupsHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeleteAuthUserGroups(id)
	if err != nil {
		log.Printf("Ошибка при удалении связи группы и пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении связи группы и пользователя"})
		return
	}

	log.Printf("связи группы и пользователя с ID %d успешно удалено", id)
	c.JSON(http.StatusOK, gin.H{"message": "Связь группы и пользователя успешно удалено"})
}

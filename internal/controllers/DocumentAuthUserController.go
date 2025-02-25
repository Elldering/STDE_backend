package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetDocumentAuthUsersHandler(c *gin.Context) {
	data, err := services.GetDocumentAuthUsers()
	if err != nil {
		log.Printf("Ошибка при получении связей документов и пользователей: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetDocumentAuthUserByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetDocumentAuthUserById(id)
	if err != nil {
		log.Printf("Ошибка при получении связи документа и пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostDocumentAuthUserHandler(c *gin.Context) {
	var docAuthUser models.DocumentAuthUser
	if err := c.ShouldBindJSON(&docAuthUser); err != nil {
		log.Printf("Некорректный формат данных: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostDocumentAuthUser(docAuthUser)
	if err != nil {
		log.Printf("Ошибка при создании связи документа и пользователя: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при создании связи"})
		return
	}
	c.JSON(http.StatusOK, docAuthUser)
}

func PutDocumentAuthUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var docAuthUser models.DocumentAuthUser
	if err := c.ShouldBindJSON(&docAuthUser); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат JSON"})
		return
	}
	err = services.PutDocumentAuthUser(id, docAuthUser)
	if err != nil {
		log.Printf("Ошибка при обновлении связи документа и пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при обновлении связи"})
		return
	}
	log.Printf("Связь документа и пользователя с ID %d успешно обновлена", id)
	c.JSON(http.StatusOK, docAuthUser)
}

func DeleteDocumentAuthUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err = services.DeleteDocumentAuthUser(id)
	if err != nil {
		log.Printf("Ошибка при удалении связи документа и пользователя с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении связи"})
		return
	}
	log.Printf("Связь документа и пользователя с ID %d успешно удалена", id)
	c.JSON(http.StatusOK, gin.H{"message": "Связь успешно удалена"})
}

package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserDocumentsHandler(c *gin.Context) {
	data, err := services.GetUserDocuments()
	if err != nil {
		log.Printf("Ошибка при получении пользовательских документов: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetUserDocumentByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetUserDocumentById(id)
	if err != nil {
		log.Printf("Ошибка при получении документа с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostUserDocumentHandler(c *gin.Context) {
	var userDoc models.UserDocument
	if err := c.ShouldBindJSON(&userDoc); err != nil {
		log.Printf("Некорректный формат данных: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostUserDocument(userDoc)
	if err != nil {
		log.Printf("Ошибка при создании пользовательского документа: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при создании документа"})
		return
	}
	c.JSON(http.StatusOK, userDoc)
}

func PutUserDocumentHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var userDoc models.UserDocument
	if err := c.ShouldBindJSON(&userDoc); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат JSON"})
		return
	}
	err = services.PutUserDocument(id, userDoc)
	if err != nil {
		log.Printf("Ошибка при обновлении документа с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при обновлении документа"})
		return
	}
	log.Printf("Документ с ID %d успешно обновлен", id)
	c.JSON(http.StatusOK, userDoc)
}

func DeleteUserDocumentHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err = services.DeleteUserDocument(id)
	if err != nil {
		log.Printf("Ошибка при удалении документа с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении документа"})
		return
	}
	log.Printf("Документ с ID %d успешно удален", id)
	c.JSON(http.StatusOK, gin.H{"message": "Документ успешно удален"})
}

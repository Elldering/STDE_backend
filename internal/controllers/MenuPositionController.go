package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetMenuPositionsHandler(c *gin.Context) {
	data, err := services.GetMenuPositions()
	if err != nil {
		log.Printf("Ошибка при получении позиций меню: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetMenuPositionByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetMenuPositionById(id)
	if err != nil {
		log.Printf("Ошибка при получении позиции меню с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostMenuPositionHandler(c *gin.Context) {
	var menuPosition models.MenuPosition
	if err := c.ShouldBindJSON(&menuPosition); err != nil {
		log.Printf("Некорректный формат данных: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostMenuPosition(menuPosition)
	if err != nil {
		log.Printf("Ошибка при создании позиции меню: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при создании позиции меню"})
		return
	}
	c.JSON(http.StatusOK, menuPosition)
}

func PutMenuPositionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var menuPosition models.MenuPosition
	if err := c.ShouldBindJSON(&menuPosition); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат JSON"})
		return
	}
	err = services.PutMenuPosition(id, menuPosition)
	if err != nil {
		log.Printf("Ошибка при обновлении позиции меню с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при обновлении позиции меню"})
		return
	}
	log.Printf("Позиция меню с ID %d успешно обновлена", id)
	c.JSON(http.StatusOK, menuPosition)
}

func DeleteMenuPositionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err = services.DeleteMenuPosition(id)
	if err != nil {
		log.Printf("Ошибка при удалении позиции меню с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении позиции меню"})
		return
	}
	log.Printf("Позиция меню с ID %d успешно удалена", id)
	c.JSON(http.StatusOK, gin.H{"message": "Позиция меню успешно удалена"})
}

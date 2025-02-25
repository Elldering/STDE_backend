package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetOrderPositionsHandler(c *gin.Context) {
	data, err := services.GetOrderPositions()
	if err != nil {
		log.Printf("Ошибка при получении позиций заказа: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetOrderPositionByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetOrderPositionById(id)
	if err != nil {
		log.Printf("Ошибка при получении позиции заказа с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostOrderPositionHandler(c *gin.Context) {
	var orderPosition models.OrderPosition
	if err := c.ShouldBindJSON(&orderPosition); err != nil {
		log.Printf("Некорректный формат данных: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostOrderPosition(orderPosition)
	if err != nil {
		log.Printf("Ошибка при создании позиции заказа: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при создании позиции заказа"})
		return
	}
	c.JSON(http.StatusOK, orderPosition)
}

func PutOrderPositionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var orderPosition models.OrderPosition
	if err := c.ShouldBindJSON(&orderPosition); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат JSON"})
		return
	}
	err = services.PutOrderPosition(id, orderPosition)
	if err != nil {
		log.Printf("Ошибка при обновлении позиции заказа с ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при обновлении позиции заказа"})
		return
	}
	log.Printf("Позиция заказа с ID %d успешно обновлена", id)
	c.JSON(http.StatusOK, orderPosition)
}

func DeleteOrderPositionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err = services.DeleteOrderPosition(id)
	if err != nil {
		log.Printf("Ошибка при удалении позиции заказа с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении позиции заказа"})
		return
	}
	log.Printf("Позиция заказа с ID %d успешно удалена", id)
	c.JSON(http.StatusOK, gin.H{"message": "Позиция заказа успешно удалена"})
}

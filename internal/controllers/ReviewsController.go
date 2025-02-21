package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetReviewsAllHandler(c *gin.Context) {
	data, err := services.GetReviewsAll()
	if err != nil {
		log.Printf("Не удалось получить список отзывов: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetReviewsByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	data, err := services.GetReviewsById(id)
	if err != nil {
		log.Printf("Ошибка при получении отзыва с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostReviewsHandler(c *gin.Context) {
	var data models.Reviews
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err := services.PostReviews(data)
	if err != nil {
		log.Printf("Ошибка при создании отзыва: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании отзыва"})
		return
	}

	log.Printf(" Право доступа %s успешно создано")
	c.JSON(http.StatusOK, data)
}

package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetPositionsHandler(c *gin.Context) {
	data, err := services.GetAllPositions()
	if err != nil {
		log.Printf("Ошибка при получении позиций: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetPositionByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	data, err := services.GetPositionById(id)
	if err != nil {
		log.Printf("Ошибка при получении блюда с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostPositionHandler(c *gin.Context) {
	var agp models.Position
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostPosition(agp)
	if err != nil {
		log.Printf("Ошибка при создании блюда: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании блюда"})
		return
	}

	log.Printf("Блюдо %s успешно создано", agp.Name)
	c.JSON(http.StatusOK, agp)
}

func PutPositionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var agp models.Position
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err = services.PutPosition(id, agp)
	if err != nil {
		log.Printf("Ошибка при обновлении блюда с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении права доступа"})
		return
	}
	log.Printf("Блюдо с ID %d успешно обновлено", id)
	c.JSON(http.StatusOK, agp)
}

func DeletePositionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeletePosition(id)
	if err != nil {
		log.Printf("Ошибка при удалении блюда с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении блюда"})
		return
	}

	log.Printf("блюдо с ID %d успешно удалено", id)
	c.JSON(http.StatusOK, gin.H{"message": "Блюдо успешно удалено"})
}

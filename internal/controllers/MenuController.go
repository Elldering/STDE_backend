package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetMenuHandler(c *gin.Context) {
	data, err := services.GetMenu()
	if err != nil {
		log.Printf("Ошибка при получени Меню")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}
func GetMenuByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetMenuById(id)
	if err != nil {
		log.Printf("Ошибка при получении блюда с ID %d, %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, data)
}
func PostMenuHandler(c *gin.Context) {
	var agp models.Menu
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный формат данных: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostMenu(agp)
	if err != nil {
		log.Printf("Ошибка при создании меню: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при создании меню"})
		return
	}
	c.JSON(http.StatusOK, agp)
}
func PutMenuHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var agp models.Menu
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат JSON"})
		return
	}

	err = services.PutMenu(id, agp)
	if err != nil {
		log.Printf("Ошибка при обновлении меню с ID %d, %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при обновлении меню"})
		return
	}
	log.Printf("Меню с Id %d успешно обновлено", id)
	c.JSON(http.StatusOK, agp)
}
func DeleteMenuHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err = services.DeleteMenu(id)
	if err != nil {
		log.Printf("Ошибка при удалении Меню с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении Меню"})
		return
	}

	log.Printf("Меню с ID %d успешно удалено", id)
	c.JSON(http.StatusOK, gin.H{"message": "Меню успешно удалено"})
}

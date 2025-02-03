package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetPermissionHandler(c *gin.Context) {
	permissions, err := services.GetPermissions()
	if err != nil {
		log.Printf("Не удалось получить список прав доступа: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, permissions)
}

func GetPermissionByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	permission, err := services.GetPermissionById(id)
	if err != nil {
		log.Printf("Ошибка при получении права доступа с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, permission)
}

func PostPermissionHandler(c *gin.Context) {
	var agp models.Permission
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err := services.PostPermission(agp)
	if err != nil {
		log.Printf("Ошибка при создании права доступа: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании права доступа"})
		return
	}

	log.Printf(" Право доступа %s успешно создано", agp.Codename)
	c.JSON(http.StatusOK, agp)
}

func PutPermissionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var agp models.Permission
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err = services.PutPermission(id, agp)
	if err != nil {
		log.Printf("Ошибка при обновлении права доступа с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении права доступа"})
		return
	}

	log.Printf("Право доступа с ID %d успешно обновлено", id)
	c.JSON(http.StatusOK, agp)
}

func DeletePermissionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeletePermission(id)
	if err != nil {
		log.Printf("Ошибка при удалении права доступа с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении права доступа"})
		return
	}

	log.Printf("Право доступа с ID %d успешно удалено", id)
	c.JSON(http.StatusOK, gin.H{"message": "Право доступа успешно удалено"})
}

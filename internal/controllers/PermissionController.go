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
		// Логируем ошибку и возвращаем HTTP статус 500 и сообщение об ошибке
		log.Printf("Ошибка при получении групп пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	// Возвращаем HTTP статус 200 и JSON с данными групп пользователей
	c.JSON(http.StatusOK, permissions)
}
func GetPermissionByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	log.Printf("Ошибка при преобразовании id: %d", id)
	if err != nil {
		log.Printf("Ошибка при преобразовании id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	permission, err := services.GetPermissionById(id)
	if err != nil {
		log.Printf("Ошибка при получении групы пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, permission)

}
func PostPermissionHandler(c *gin.Context) {
	var agp models.Permisson
	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.PostPermission(agp)
	if err != nil {
		log.Printf("Ошибка при создании связи группы и прав доступа: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании связи"})
		return
	}

	c.JSON(http.StatusOK, agp)
}

func PutPermissionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var agp models.Permisson
	if err := c.ShouldBindJSON(&agp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.PutPermission(id, agp)
	if err != nil {
		log.Printf("Ошибка при обновлении группы пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении группы пользователей"})
		return
	}

	c.JSON(http.StatusOK, agp)
}
func DeletePermissionHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Ошибка при преобразовании id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = services.DeletePermission(id)
	if err != nil {
		log.Printf("Ошибка при удалении группы пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении группы пользователей"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Группа пользователей успешно удалена"})
}

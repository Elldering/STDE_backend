package controllers

import (
	"STDE_proj/internal/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserGroupsHandler(c *gin.Context) {
	userGroups, err := repositories.GetAllUserGroups()
	if err != nil {
		log.Printf("Ошибка при получении групп пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, userGroups)
}

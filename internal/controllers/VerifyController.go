package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func VerifyControllerHandler(c *gin.Context) {
	typeParam := c.Param("type")
	var data models.VerifyCode
	if err := c.BindJSON(&data); err != nil {
		log.Printf("Ошибка при обработке данных запроса: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные запроса"})
		return
	}
	data.Type = typeParam
	err := services.VerifyService(data)
	if err != nil {
		log.Printf("Ошибка при подтверждении кода: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Почта успешно подтверждена")
	switch data.Type {
	case "reg":
		c.JSON(http.StatusOK, gin.H{"message": "Почта успешно подтверждена"})
	case "auth":
		c.JSON(http.StatusOK, gin.H{"message": "Верный код доступа"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Неверный тип верификации"})
	}

}

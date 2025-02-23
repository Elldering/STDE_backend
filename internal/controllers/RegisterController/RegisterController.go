package RegisterController

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services/RegisterService"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterControllerHandler(c *gin.Context) {

	var data models.Register

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := RegisterService.Register(data)
	if err != nil {
		log.Printf("Ошибка при регистрации: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Регистрация прошла успешно")
	c.JSON(http.StatusOK, data)
}

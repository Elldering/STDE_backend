package RegisterController

import (
	"STDE_proj/internal/services/RegisterService"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterControllerHandler(c *gin.Context) {
	var registerRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := RegisterService.Register(registerRequest.Username, registerRequest.Password)
	if err != nil {
		log.Printf("Ошибка при регистрации: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при регистрации"})
		return
	}
	log.Println("Регистрация прошла успешно")
	c.JSON(http.StatusOK, registerRequest)
}

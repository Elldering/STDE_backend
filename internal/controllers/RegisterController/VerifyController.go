package RegisterController

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services/RegisterService"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func VerifyControllerHandler(c *gin.Context) {

	var data models.VerifyCode

	if err := c.BindJSON(&data); err != nil {
		log.Printf("Ошибка при обработке данных запроса: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные запроса"})
		return
	}

	err := RegisterService.Verify(data)
	if err != nil {
		log.Printf("Ошибка при подтверждении кода: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Почта успешно подтверждена")
	c.JSON(http.StatusOK, gin.H{"message": "Почта успешно подтверждена"})
}

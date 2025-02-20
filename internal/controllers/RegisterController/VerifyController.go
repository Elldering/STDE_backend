package RegisterController

import (
	"STDE_proj/internal/services/RegisterService"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func VerifyControllerHandler(c *gin.Context) {
	var req struct {
		Code int `json:"code"`
	}
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Ошибка при обработке данных запроса: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные запроса"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	err = RegisterService.Verify(id, req.Code)
	if err != nil {
		log.Printf("Ошибка при подтверждении кода: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при подтверждении кода"})
		return
	}
	log.Println("Почта успешно подтверждена")
	c.JSON(http.StatusOK, gin.H{"message": "Почта успешно подтверждена"})
}

package Auth

import (
	"STDE_proj/internal/controllers/Auth/model"
	"STDE_proj/internal/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LogoutHandler(c *gin.Context) {
	access := c.Request.Header.Get("Authorization")
	var refresh model.LogoutRequest

	if err := c.ShouldBindJSON(&refresh); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	if access == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access Токен не предоставлен"})
		return
	}

	if refresh.Refresh == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh Токен не предоставлен"})
		return
	}

	err := repositories.InvalidateToken(access, refresh.Refresh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка инвалидации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Успешный выход"})
}

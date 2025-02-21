package Auth

import (
	"STDE_proj/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Токен не предоставлен"})
		return
	}

	err := repositories.InvalidateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка инвалидации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Успешный выход"})
}

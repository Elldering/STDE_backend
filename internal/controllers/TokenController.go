package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GenerateAccessRefreshToken(c *gin.Context) {
	var data models.AuthUser

	JWTSecret := os.Getenv("JWT_SECRET")

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := services.GenerateTokens(&data, JWTSecret)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})

}

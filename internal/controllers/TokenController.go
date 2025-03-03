package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

const (
	AccessTokenExpiry  = time.Minute * 15
	RefreshTokenExpiry = time.Hour * 24 * 5
)

func GenerateAccessRefreshToken(c *gin.Context) {
	var data models.AuthUserRequest

	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		c.Header("X-Is-Authenticated", "false")
		c.Status(http.StatusUnauthorized)
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := services.GenerateTokens(&data, JWTSecret)
	if err != nil {
		c.Header("X-Is-Authenticated", "false")
		c.Status(http.StatusUnauthorized)
		return
	}

	// Установка access token в cookie
	c.SetCookie("access_token", accessToken, int(AccessTokenExpiry.Seconds()), "/", "", false, true)

	// Установка refresh token в cookie
	c.SetCookie("refresh_token", refreshToken, int(RefreshTokenExpiry.Seconds()), "/", "", false, true)

	// Возвращаем успешный ответ
	c.Header("X-Is-Authenticated", "true")
	c.Status(http.StatusOK)
}

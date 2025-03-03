package controllers

import (
	"STDE_proj/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenCheckController(c *gin.Context) {

	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.Header("X-Is-Authenticated", "false")
		c.Status(http.StatusUnauthorized)
		return
	}

	isValid, _ := jwt.ValidationToken(accessToken)
	if !isValid {
		c.Header("X-Is-Authenticated", "false")
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Header("X-Is-Authenticated", "true")
	c.Status(http.StatusOK)

}

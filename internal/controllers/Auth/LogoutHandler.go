package Auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandler(c *gin.Context) {
	
	c.Header("X-Is-Authenticated", "false")

	c.SetCookie("access_token", "", 0, "/", "", false, true)

	c.SetCookie("refresh_token", "", 0, "/", "", false, true)

	c.Status(http.StatusOK)

}

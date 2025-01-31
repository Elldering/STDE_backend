package routes

import (
	"STDE_proj/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/usergroups", controllers.GetUserGroupsHandler)
}

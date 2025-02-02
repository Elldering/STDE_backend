package routes

import (
	"STDE_proj/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/user_groups", controllers.GetUserGroupsHandler)
	router.POST("/auth_group_permissions", controllers.PostAuthGroupPermissionsHandler)
}

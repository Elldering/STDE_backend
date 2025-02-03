package routes

import (
	"STDE_proj/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	//User groups
	router.GET("/user_groups", controllers.GetUserGroupsHandler)
	router.POST("/user_groups", controllers.PostUserGroupHandler)
	router.GET("/user_groups/:id", controllers.GetUserGroupByIdHandler)
	router.DELETE("/user_groups/:id", controllers.DeleteUserGroupHandler)
	router.PUT("/user_groups/:id", controllers.PutUserGroupHandler)
	//Permission
	router.GET("/permission", controllers.GetPermissionHandler)
	router.POST("/permission", controllers.PostPermissionHandler)
	router.GET("/permission/:id", controllers.GetPermissionByIdHandler)
	router.DELETE("/permission/:id", controllers.DeletePermissionHandler)
	router.PUT("/permission/:id", controllers.PutPermissionHandler)

	router.POST("/auth_group_permissions", controllers.PostAuthGroupPermissionsHandler)
}

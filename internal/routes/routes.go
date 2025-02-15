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
	//Auth group permissions
	router.GET("/auth_group_permissions", controllers.GetAuthGroupPermissionsHandler)
	router.GET("/auth_group_permissions/:id", controllers.GetAuthGroupPermissionsIdHandler)
	router.POST("/auth_group_permissions", controllers.PostAuthGroupPermissionsHandler)
	router.PUT("/auth_group_permissions/:id", controllers.PutAuthGroupPermissionsHandler)
	router.DELETE("/auth_group_permissions/:id", controllers.DeleteAuthGroupPermissionsHandler)
	// User profile
	router.GET("/user_profile", controllers.GetUserProfileHandler)
	router.GET("/user_profile/:id", controllers.GetUserProfileByIdHandler)
	router.POST("/user_profile", controllers.PostUserProfileHandler)
	router.PUT("/user_profile/:id", controllers.PutUserProfileHandler)
	router.DELETE("/user_profile/:id", controllers.DeleteUserProfileHandler)
	// Auth User
	router.POST("/auth_user", controllers.PostAuthUserHandler)
}

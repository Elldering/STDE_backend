package routes

import (
	"STDE_proj/internal/controllers"
	"STDE_proj/internal/controllers/Auth"
	"STDE_proj/internal/controllers/RegisterController"
	"STDE_proj/internal/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func Routes(router *gin.Engine) {
	JWTSecret := os.Getenv("JWT_SECRET")

	public := router.Group("/api/public")
	{
		user := public.Group("/user")
		{
			user.GET("/logout", Auth.LogoutHandler)
			// Маршрут не реализован как регистрация. Реализован как обычный Create
			//user.POST("/add", controllers.PostAuthUserHandler)
			register := user.Group("/register")
			{
				register.POST("/", RegisterController.RegisterControllerHandler)
				register.PATCH("/verify/:id", RegisterController.VerifyControllerHandler)

			}

		}

		auth := public.Group("/auth")
		{
			auth.POST("/login", Auth.LoginHandler)
			auth.POST("/refresh", Auth.RefreshToken)
		}
	}


	protected := router.Group("/api/private")
	protected.Use(middleware.AuthMiddleware(JWTSecret))
	{
		user := protected.Group("/user")
		{
			user.POST("/delete/:id", controllers.DeleteAuthUserHandler)
		}

		reviews := protected.Group("/reviews")
		{
			reviews.GET("/", controllers.GetReviewsAllHandler)
			reviews.GET("/:id", controllers.GetReviewsByIdHandler)
			reviews.POST("/", controllers.PostReviewsHandler)
		}

		AuthUserGroups := protected.Group("/auth-user-groups")
		{
			AuthUserGroups.GET("/", controllers.GetAuthUserGroupsAllHandler)
			AuthUserGroups.GET("/:id", controllers.GetAuthUserGroupsByIdHandler)
			AuthUserGroups.POST("/", controllers.PostAuthUserGroupsHandler)
			AuthUserGroups.PUT("/:id", controllers.PutAuthUserGroupsHandler)
			AuthUserGroups.DELETE("/:id", controllers.DeleteAuthUserGroupsHandler)
		}

		UserGroups := protected.Group("/user-groups")
		{
			UserGroups.GET("/", controllers.GetUserGroupsHandler)
			UserGroups.POST("/", controllers.PostUserGroupHandler)
			UserGroups.GET("/:id", controllers.GetUserGroupByIdHandler)
			UserGroups.DELETE("/:id", controllers.DeleteUserGroupHandler)
			UserGroups.PUT("/:id", controllers.PutUserGroupHandler)
		}

		Permissions := protected.Group("/permissions")
		{
			Permissions.GET("/", controllers.GetPermissionHandler)
			Permissions.POST("/", controllers.PostPermissionHandler)
			Permissions.GET("/:id", controllers.GetPermissionByIdHandler)
			Permissions.DELETE("/:id", controllers.DeletePermissionHandler)
			Permissions.PUT("/:id", controllers.PutPermissionHandler)
		}

		Position := protected.Group("/position")
		{
			Position.GET("/", controllers.GetPositionsHandler)
			Position.POST("/", controllers.PostPositionHandler)
			Position.GET("/:id", controllers.GetPositionByIdHandler)
			Position.DELETE("/:id", controllers.DeletePositionHandler)
			Position.PUT("/:id", controllers.PutPositionHandler)
		}

		AuthGroupsPermissions := protected.Group("/auth-groups-permissions")
		{
			AuthGroupsPermissions.GET("/", controllers.GetAuthGroupPermissionsHandler)
			AuthGroupsPermissions.GET("/:id", controllers.GetAuthGroupPermissionsIdHandler)
			AuthGroupsPermissions.POST("/", controllers.PostAuthGroupPermissionsHandler)
			AuthGroupsPermissions.PUT("/:id", controllers.PutAuthGroupPermissionsHandler)
			AuthGroupsPermissions.DELETE("/:id", controllers.DeleteAuthGroupPermissionsHandler)
		}

		UserProfile := protected.Group("/user-profile")
		{
			UserProfile.GET("/", controllers.GetUserProfileHandler)
			UserProfile.GET("/:id", controllers.GetUserProfileByIdHandler)
			UserProfile.POST("/", controllers.PostUserProfileHandler)
			UserProfile.PUT("/:id", controllers.PutUserProfileHandler)
			UserProfile.DELETE("/:id", controllers.DeleteUserProfileHandler)
		}

	}

}

package routes

import (
	"STDE_proj/internal/controllers"
	"STDE_proj/internal/controllers/Auth"
	"github.com/gin-gonic/gin"
)

func TestRoutes(router *gin.Engine) {
	test := router.Group("api/test-routes")
	{
		user := router.Group("/user")
		{
			user.POST("/register", controllers.PostAuthUserHandler)
			user.POST("/delete/:id", controllers.DeleteAuthUserHandler)
		}

		metrics := router.Group("/metrics")
		{
			metrics.GET("/user-count", controllers.MetricsHandler())
		}
		auth := router.Group("/auth")
		{
			auth.POST("/login", Auth.LoginHandler)
			auth.POST("/refresh", Auth.RefreshToken)
		}

		s3Group := router.Group("/s3")
		{
			// POST — загрузка файла
			s3Group.POST("/upload", controllers.PostFileHandler)

			// DELETE — удаление файла
			s3Group.DELETE("/files/:filename", controllers.DeleteFileHandler)

			// Список бакетов
			s3Group.GET("/buckets", controllers.ListBucketsHandler)

			// Список файлов
			s3Group.GET("/files", controllers.ListFilesHandler)

			// GET — получение файла/подписанного URL
			s3Group.GET("/files/url/:bucket/:filename", controllers.GetFileURLHandler)

			// Скачивание файла
			s3Group.GET("/download/:filename", controllers.DownloadFileHandler)

			// Пока не сделал
			//s3Group.PUT("/files/:filename", s3Controller.PutFileHandler)
		}

		// Создание вложенной группы маршрутов для "reviews" внутри группы "test"
		reviews := test.Group("/reviews")
		{
			// Обработчики маршрутов для "reviews"
			reviews.GET("/", controllers.GetReviewsAllHandler)
			reviews.GET("/:id", controllers.GetReviewsByIdHandler)
			reviews.POST("/", controllers.PostReviewsHandler)
		}
		authUserGroups := test.Group("/auth-user-groups")
		{
			// Auth user groups
			authUserGroups.GET("/", controllers.GetAuthUserGroupsAllHandler)
			authUserGroups.GET("/:id", controllers.GetAuthUserGroupsByIdHandler)
			authUserGroups.POST("/", controllers.PostAuthUserGroupsHandler)
			authUserGroups.PUT("/:id", controllers.PutAuthUserGroupsHandler)
			authUserGroups.DELETE("/:id", controllers.DeleteAuthUserGroupsHandler)
		}
		UserGroups := test.Group("/user-groups")
		{
			UserGroups.GET("/", controllers.GetUserGroupsHandler)
			UserGroups.POST("/", controllers.PostUserGroupHandler)
			UserGroups.GET("/:id", controllers.GetUserGroupByIdHandler)
			UserGroups.DELETE("/:id", controllers.DeleteUserGroupHandler)
			UserGroups.PUT("/:id", controllers.PutUserGroupHandler)
		}
		Permissions := test.Group("/permissions")
		{
			Permissions.GET("/", controllers.GetPermissionHandler)
			Permissions.POST("/", controllers.PostPermissionHandler)
			Permissions.GET("/:id", controllers.GetPermissionByIdHandler)
			Permissions.DELETE("/:id", controllers.DeletePermissionHandler)
			Permissions.PUT("/:id", controllers.PutPermissionHandler)
		}
		Position := test.Group("/position")
		{
			Position.GET("/", controllers.GetPositionsHandler)
			Position.POST("/", controllers.PostPositionHandler)
			Position.GET("/:id", controllers.GetPositionByIdHandler)
			Position.DELETE("/:id", controllers.DeletePositionHandler)
			Position.PUT("/:id", controllers.PutPositionHandler)
		}
		AuthGroupsPermissions := test.Group("/auth-groups-permissions")
		{
			AuthGroupsPermissions.GET("/", controllers.GetAuthGroupPermissionsHandler)
			AuthGroupsPermissions.GET("/:id", controllers.GetAuthGroupPermissionsIdHandler)
			AuthGroupsPermissions.POST("/", controllers.PostAuthGroupPermissionsHandler)
			AuthGroupsPermissions.PUT("/:id", controllers.PutAuthGroupPermissionsHandler)
			AuthGroupsPermissions.DELETE("/:id", controllers.DeleteAuthGroupPermissionsHandler)
		}
		UserProfile := test.Group("/user-profile")
		{
			UserProfile.GET("/", controllers.GetUserProfileHandler)
			UserProfile.GET("/:id", controllers.GetUserProfileByIdHandler)
			UserProfile.POST("/", controllers.PostUserProfileHandler)
			UserProfile.PUT("/:id", controllers.PutUserProfileHandler)
			UserProfile.DELETE("/:id", controllers.DeleteUserProfileHandler)
		}

	}
}
